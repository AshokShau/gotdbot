package logger

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"strings"
	"sync"
	"testing"
)

func TestLevelFormattingAndAttributes(t *testing.T) {
	var out bytes.Buffer
	l := New(WithLevel(slog.LevelInfo), WithOutput(&out))
	l.Debug("hidden")
	l.Debugf("hidden %d", 1)
	l.Info("info", "id", 42)
	l.Warnf("warn %s", "message")
	l.Errorf("error %d", 7)
	l.With("scope", "child").Info("child")
	text := out.String()
	for _, want := range []string{"INFO", "WARN", "ERROR", "info id=42", "warn message", "error 7", "child scope=child", "logger_test.go:"} {
		if !strings.Contains(text, want) {
			t.Errorf("missing %q in %q", want, text)
		}
	}
	if strings.Contains(text, "hidden") {
		t.Fatal("filtered debug message emitted")
	}
	if len(strings.Split(strings.TrimSpace(text), "\n")) != 4 {
		t.Fatalf("unexpected log lines: %q", text)
	}
}

func TestConcurrentChildLoggers(t *testing.T) {
	var out bytes.Buffer
	l := New(WithOutput(&out))
	const workers, records = 16, 32
	var wg sync.WaitGroup
	for i := range workers {
		wg.Go(func() {
			child := l.With("worker", i)
			for j := range records {
				child.Info("record", "index", j)
			}
		})
	}
	wg.Wait()
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != workers*records {
		t.Fatalf("record count = %d", len(lines))
	}
	for _, line := range lines {
		if !strings.Contains(line, "record index=") || strings.Count(line, "worker=") != 1 {
			t.Fatalf("interleaved log record %q", line)
		}
	}
}

func TestCustomHandlerAndNilLoggers(t *testing.T) {
	var out bytes.Buffer
	h := slog.NewJSONHandler(&out, &slog.HandlerOptions{Level: slog.LevelDebug})
	l := New(WithOutput(io.Discard), WithLevel(slog.LevelError), WithHandler(h))
	l.Debug("custom", "enabled", true)
	if !strings.Contains(out.String(), `"msg":"custom"`) {
		t.Fatal("custom handler did not override level/output")
	}
	for _, l := range []*Logger{nil, {}, Disabled(), From(nil)} {
		l.Debug("ignored")
		l.Info("ignored")
		l.Warn("ignored")
		l.Error("ignored")
		l.Debugf("ignored %d", 1)
		l.Infof("ignored %d", 1)
		l.Warnf("ignored %d", 1)
		l.Errorf("ignored %d", 1)
	}
	From(slog.New(h)).Info("wrapped")
	if !strings.Contains(out.String(), `"msg":"wrapped"`) {
		t.Fatal("wrapped slog logger did not log")
	}
}

func TestHandlerChildIsolationAndGroups(t *testing.T) {
	var out bytes.Buffer
	h := &plainHandler{out: &out, mu: &sync.Mutex{}, level: slog.LevelInfo}
	a := h.WithAttrs([]slog.Attr{slog.String("a", "one")})
	b := a.WithAttrs([]slog.Attr{slog.String("b", "two")})
	g := b.WithGroup("outer").WithGroup("inner")
	r := slog.Record{}
	r.Message = "message"
	r.Level = slog.LevelInfo
	for _, handler := range []slog.Handler{h, a, b, g} {
		if err := handler.Handle(context.Background(), r); err != nil {
			t.Fatal(err)
		}
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if strings.Contains(lines[0], "a=") || strings.Contains(lines[1], "b=") || !strings.Contains(lines[2], "a=one b=two") || !strings.Contains(lines[3], "outer.inner") {
		t.Fatalf("child handler mutation: %q", out.String())
	}
}
