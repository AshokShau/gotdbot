# Development and troubleshooting

Contents: architecture; ownership; generation; validation; test strategy; diagnostic table.

## Repository architecture

| Area | Main source | Role |
| --- | --- | --- |
| Client lifecycle and options | `client.go`, `client_opts.go`, `client_auth.go` | Native loading, account authentication, workers, shutdown. |
| Shared receiver | `manager.go`, `response_router.go` | One receiver routes responses/updates to account clients. |
| Request lifecycle | `client_requests.go`, `request_lifecycle.go` | Correlation, deadlines, bounded retries, pending-message completion. |
| Dispatch and buffering | `dispatcher.go`, `handlers.go`, `update_queue.go` | Worker processing, immutable registration snapshots, bounded memory/disk spill. |
| Commands and filters | `command_handler.go`, `message_handler.go`, `update_handler.go`, `filters/` | Matching, handlers, composable predicates. |
| Waiters | `conversation.go`, `waiters.go` | Ask, timeout, cancellation, chat-specific/global candidate selection. |
| Send/reply/transfer | `send_helpers.go`, `message_reply.go`, `message_transfer.go` | Text/media/rich/ephemeral content and copied reply settings. |
| Read/edit/action helpers | `message_content.go`, `message_media.go`, `message_actions.go`, `message_edit.go`, `edit_messages.go` | Accessors, downloads, reactions, content edits. |
| Formatting | `formatting.go`, `formatting_parse.go`, `message_formatting.go` | Parsing, UTF-16 entity rendering, escaped mentions. |
| Update helpers | `callback_helpers.go`, `update_helpers.go` | Context-specific answers/edits and message/ID extraction. |
| Chat and bot helpers | `chat_helpers.go`, `chat_admin.go`, `forum_helpers.go`, `bot_profile.go` | Chat senders, moderation, topics, current bot profile. |
| Native bindings | All root `gen_*.go` | Types, unions, requests, methods, opts, event registration, filters, receiver helpers, typed TDLib options. |
| Binding generator | `internal/gen/` | Schema loading, type mapping, naming/collision decisions, all generated output. |
| Runtime loader | `internal/tdjson/` | purego calls to native td_create_client_id/td_send/td_receive/td_execute. |
| Utilities | `logger/`, `internal/qrcode/`, `scripts/tools/` | Application logs, auth QR display, native library downloader. |
| Usage examples | `examples/` | Echo, auto-delete, conversations, multiple accounts. |

The `internal/` packages are implementation details and cannot be imported from unrelated application modules. Use public APIs in consumers; package-local test seams are available only inside this library's tests.

## Where to make a change

- Convenience behavior belongs in handwritten helpers. Reuse native request structs and existing shared builders to keep delivery, reply, and formatting options consistent.
- Binding shape or emission belongs in `internal/gen`. Root generated files are outputs, not durable edit targets.
- Native method collision mappings live in `internal/gen/api_names.go`. The generator detects handwritten methods; Client collisions require an explicit native name, and generated receiver helpers skip existing handwritten methods.
- Optional/default parameter rules live in `internal/gen/parameters.go`; boolean/nullable/default decisions affect both methods and Opts.
- Type and codec rules live in `internal/gen/utils.go` and `types.go`; shared helper generation lives in `helpers.go`.
- Keep caller-owned configuration/options/request structs unchanged. Existing tests capture requests and compare preserved inputs; use them as implementation evidence.

Before adding another helper, inspect the equivalent send/reply/edit paths and tests. Ordinary application development does not need generator edits or binding regeneration.

## Generation and schema discipline

From the repository root:

```sh
go generate .
```

The baseline generator fetches the master JSON schema from `FallenProjects/tdlib-build`. That source can change independently of the installed TDLib binary. For a deliberate, reproducible generation task, supply the intended local schema:

```sh
go run ./internal/gen -schema /path/to/tdlib.json
```

The argument also accepts an HTTP(S) URL. Run from the module root because output is written relative to the working directory. It rewrites generated root bindings and formats them. Inspect schema Version, generated changes, `TDLibVersion`, downloader version, and README requirements for agreement. Do not casually run generation as a fix for an application compile error or automatically accept unrelated upstream schema changes.

Review regenerated native names, params/options, union codecs, filters/events, and handwritten collisions. Never use a newer JSON schema as proof that an older binary accepts its requests. Use an isolated output directory for generator experiments when avoiding main-tree rewrites; tests in `internal/gen` illustrate controlled generation.

## Validation

For a library implementation/generator change, run relevant focused tests first and complete the repository checks:

```sh
go test ./...
go test -race ./...
go vet ./...
```

These tests use transport-boundary fakes and do not need Telegram credentials or an installed shared library. Application compilation also does not load TDLib; `NewClient` at runtime does. A clean build is not evidence of successful authorization, native binary compatibility, delivery, or platform deployment.

For skill/documentation edits, validate frontmatter/links and compile Go examples against the chosen module. Do not regenerate bindings or run live Telegram mutations simply to validate prose. When updating this skill after an API change, revise the affected reference/examples and the SKILL.md baseline together.

## Useful test patterns and invariants

`newTestClient`, `captureHelperRequests`, and `startTestProcessors` in the root tests exercise the real lifecycle without TDLib. They are private test helpers, not an exported mocking API.

| Area | Tests to inspect | Invariant |
| --- | --- | --- |
| Request completion/cancel/retry | `client_test.go`, `request_lifecycle_test.go`, `response_router_test.go` | Deadline-bounded calls; pending map cleanup; sends complete independent of occupied handlers. |
| Handlers/groups/filter defaults | `handlers_test.go`, `update_handler_test.go`, `filters/*/*_test.go` | Correct first-match/group behavior, outgoing/default options, nil/regex/composition behavior. |
| Conversations | `conversation_test.go` | Sender/chat filtering, timeout/cancellation/closure, safe waiter callbacks. |
| Helper options/delivery | `helper_options_test.go`, `send_options_test.go`, `message_reply_test.go`, `helper_revision_test.go` | Explicit reply preservation; no caller mutation; metadata/delivery forwarded. |
| Formatting/media/rich edits | `formatting*_test.go`, `media_helpers_test.go`, `edit_helpers_test.go`, `rich_message_helpers_test.go`, `message_draft_test.go` | UTF-16 correctness, nil media handling, correct edit dispatch and draft fields. |
| JSON and generation | `json_types_test.go`, `internal/gen/*_test.go` | Integer precision, correct union codecs, generated signatures/names/options and collision rules. |

For behavior changes, test the observable request/result or lifecycle invariant. Avoid wording-only tests for docs or tests that simply copy the implementation. Use the race detector for changed shared state and worker interactions.

## Diagnose from the symptom

| Symptom | Check and next action |
| --- | --- |
| Cannot load tdjson | Explicit LibraryPath, process working directory, native OS/architecture, transitive shared dependencies, exported td_* symbols. |
| Unknown type or decode failure | Installed TDLib binary vs generated schema; inspect receiver errors rather than discarding updates. |
| Undefined method or unexpected argument count | Selected GotdBot version and `go doc`; distinguish convenience/native names and option structs. |
| Start waits for phone/code unexpectedly | Missing/malformed bot token is interpreted as phone login; validate config. |
| Authorization timeout | Account/token/API credentials, first-login stdin/QR flow, persisted session/key, timeout. Close an abandoned client. |
| Client identity unavailable | Start must succeed before identity-dependent/profile methods. |
| Wrong recipient or message not found | Native chat/message IDs, account association, cache/message-database behavior; optional lookup retries do not convert Bot API IDs. |
| Outgoing handler never runs | Enable `AllowOutgoing` in MatchingOptions; predicates cannot bypass it. |
| Later handler in same group never runs | Earlier handler matched; reorder or return ContinueHandlers. |
| Ordinary edits do not reach OnMessage | Register OnUpdateMessageEdited/OnUpdateMessageContent and fetch their message. |
| A custom command reaches the echo handler | `message.Command` relies on bot-command entities; custom triggers need explicit exclusions. |
| Ask times out despite incoming messages | All workers blocked, wrong sender/chat filter, business vs ordinary update, waiter registered after arrival. |
| Conversations interfere | Multiple waiters can match one update; scope and serialize application conversations by account/chat/sender. |
| Formatting fails or literal markup leaks | Entities/ParseMode/default precedence, unescaped user input, ordinary vs rich syntax. |
| Emoji entity spans are wrong | Compute offsets in UTF-16 units; do not use byte/rune counts. |
| Download returns no complete file | Supported attachment, original client/file ID, synchronous setting or UpdateFile completion. |
| Inline media edit fails | Files must already be uploaded; use inline IDs/options and supported editable content. |
| Moderation/profile/payment method rejected | Native account kind, ownership, member/admin rights, feature-specific constraints and target data. |
| Messages duplicate after retry | Timeout may follow server acceptance; inspect send-completion/state before retrying. |
| 429/flood wait | Parsed retry-after, enabled MaxFloodWait, bounded current RequestTimeout; no infinite immediate retry loops. |
| Update backlog grows | Handler duration/worker count, spool errors and temporary-disk capacity. Queue spill is not restart persistence. |
| Shutdown returns while jobs run | Close deliberately does not wait for application handlers; manage job draining separately. |

Report which behavior was actually tested. Separate compile/unit/race checks from real-account tests when describing readiness.
