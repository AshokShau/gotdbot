package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadAndConvertLocalSchema(t *testing.T) {
	path := filepath.Join(t.TempDir(), "schema.json")
	schema := `{"version":"test","classes":{"MessageContent":{"description":"content","types":["messageText"]}},"types":{"messageText":{"description":"text","type":"MessageContent","args":{"z_text":{"type":"string"},"a_optional":{"type":"int64","is_optional":true}}}},"updates":{"updateNewMessage":{"type":"Update","args":{"message":{"type":"messageText"}}}},"functions":{"getMe":{"type":"User","args":{}}},"options":{"online":{"type":"Bool","writable":true}}}`
	if err := os.WriteFile(path, []byte(schema), 0600); err != nil {
		t.Fatal(err)
	}
	data, err := LoadTDLibJSON(path)
	if err != nil {
		t.Fatal(err)
	}
	types, functions, classes, options := ConvertJSONToGen(data)
	if data.Version != "test" || len(types) != 2 || len(functions) != 1 || len(classes) != 1 || len(options) != 1 {
		t.Fatal("schema conversion lost sections")
	}
	for _, typ := range types {
		if typ.Name == "messageText" {
			if len(typ.Params) != 2 || typ.Params[0].Name != "a_optional" || !typ.Params[0].IsOptional || typ.Params[1].Name != "z_text" || typ.ResultType != "MessageContent" {
				t.Fatalf("sorted/optional parameters = %+v", typ)
			}
		}
	}
	if !reflect.DeepEqual(classes["MessageContent"].Implementations, []string{"messageText"}) {
		t.Fatal("class implementations lost")
	}
	for _, input := range []string{`{`, `{}`, `{"version":"x","types":{"a":{}}}`, `{"version":"x","functions":{"a":{}}}`} {
		if err := os.WriteFile(path, []byte(input), 0600); err != nil {
			t.Fatal(err)
		}
		if _, err := LoadTDLibJSON(path); err == nil {
			t.Errorf("invalid schema accepted: %s", input)
		}
	}
	if _, err := LoadTDLibJSON(filepath.Join(t.TempDir(), "missing.json")); err == nil {
		t.Fatal("missing schema accepted")
	}
}

func TestGoTypeMappingsAndMethodCollisions(t *testing.T) {
	classes := map[string]*TLClass{"MessageContent": {Name: "MessageContent"}}
	for input, want := range map[string]string{
		"int32": "int32", "int53": "int64", "int64": "int64", "double": "float64", "string": "string", "bytes": "[]byte", "Bool": "bool",
		"MessageContent": "MessageContent", "message": "*Message", "vector<int64>": "Int64Slice", "vector<message>": "[]Message", "vector<MessageContent>": "[]MessageContent", "vector<vector<int64>>": "[]Int64Slice",
	} {
		if got := toGoType(input, classes); got != want {
			t.Errorf("toGoType(%q) = %q, want %q", input, got, want)
		}
	}
	for input, want := range map[string]string{"sendMessage": "SendMessageContent", "editMessageText": "EditMessageTextContent", "editMessageCaption": "EditMessageCaptionText", "getMe": "GetMe", "get_chat": "GetChat"} {
		if got := clientMethodName(input); got != want {
			t.Errorf("method %q = %q, want %q", input, got, want)
		}
	}
	for input, want := range map[string]string{"": "", "__get__chat": "GetChat", "évent_type": "ÉventType"} {
		if got := toCamelCase(input); got != want {
			t.Errorf("name %q = %q, want %q", input, got, want)
		}
	}
}

func TestGeneratedClassDecoderHandlesNullableFields(t *testing.T) {
	// Compile and execute generated code in isolation. This protects both the
	// generator and checked-in output from reintroducing the spool failure.
	dir := t.TempDir()
	t.Chdir(dir)
	classes := map[string]*TLClass{"Value": {Name: "Value", Implementations: []string{"stringValue"}}}
	generateClasses(classes)
	generateObjects([]TLType{
		{Name: "error", Params: []TLParam{{Name: "code", Type: "int32"}, {Name: "message", Type: "string"}}},
		{Name: "stringValue", ResultType: "Value", Params: []TLParam{{Name: "value", Type: "string"}}},
		{Name: "valueHolder", Params: []TLParam{{Name: "item", Type: "Value"}, {Name: "items", Type: "vector<Value>"}}},
	}, classes)
	files := map[string]string{
		"go.mod": "module generatorfixture\n\ngo 1.25.0\n",
		"fixture_test.go": `package gotdbot
import ("encoding/json"; "reflect"; "testing")
func TestDecoder(t *testing.T) {
 for _, data := range []string{"null", " \n null \t "} {
  obj, err := UnmarshalValue([]byte(data))
  if obj != nil || err != nil { t.Fatalf("nullable field: (%v, %v)", obj, err) }
 }
 obj, err := UnmarshalValue([]byte("{\"@type\":\"stringValue\",\"value\":\"hello\"}"))
 if err != nil || obj.(*StringValue).Value != "hello" { t.Fatalf("valid value: (%v, %v)", obj, err) }
 for _, data := range []string{"{}", "{", "{\"@type\":\"unknown\"}"} {
  if _, err := UnmarshalValue([]byte(data)); err == nil { t.Fatalf("invalid value accepted: %s", data) }
 }
}
func TestObjectNull(t *testing.T) {
 for _, data := range []string{"null", " \n null \t "} {
  got := ValueHolder{Item: &StringValue{Value: "keep"}, Items: []Value{&StringValue{Value: "keep"}}}
  want := got
  if err := json.Unmarshal([]byte(data), &got); err != nil { t.Fatal(err) }
  if !reflect.DeepEqual(got, want) { t.Fatal("null changed existing object") }
 }
 var holders []ValueHolder
 if err := json.Unmarshal([]byte("[null,{\"item\":{\"@type\":\"stringValue\",\"value\":\"hello\"},\"items\":[null,{\"@type\":\"stringValue\",\"value\":\"world\"}]}]"), &holders); err != nil { t.Fatal(err) }
 if len(holders) != 2 || !reflect.DeepEqual(holders[0], ValueHolder{}) { t.Fatalf("null element: %+v", holders) }
 got := holders[1]
 if got.Item.(*StringValue).Value != "hello" || len(got.Items) != 2 || got.Items[0] != nil || got.Items[1].(*StringValue).Value != "world" { t.Fatalf("class fields: %+v", got) }
}

`,
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0600); err != nil {
			t.Fatal(err)
		}
	}
	cmd := exec.Command("go", "test", ".")
	cmd.Env = append(os.Environ(), "GOWORK=off")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated decoder: %v\n%s", err, output)
	}
}

func TestTDLibOptionsGenerationLeavesHandwrittenConfigUntouched(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	config := []byte("package gotdbot\n\n// This file belongs to the client configuration.\ntype ClientOpts struct { UpdateWorkers int }\n")
	if err := os.WriteFile("client_opts.go", config, 0600); err != nil {
		t.Fatal(err)
	}
	generateTDLibOptions(map[string]*OptionDef{
		"online":   {Type: "Bool", Writable: true},
		"limit":    {Type: "int64", Writable: true},
		"language": {Type: "string", Writable: true},
		"version":  {Type: "string", Writable: false},
	})
	after, err := os.ReadFile("client_opts.go")
	if err != nil || !reflect.DeepEqual(after, config) {
		t.Fatal("generator changed handwritten client options")
	}
	files := map[string]string{
		"go.mod": "module optionsfixture\n\ngo 1.25.0\n",
		"fixture_test.go": `package gotdbot
import ("reflect"; "testing")
func TestOptions(t *testing.T) {
 opts := &TDLibOptions{Online: true, Limit: -42, Language: "en"}
 got := map[string]any{}
 opts.forEachSet(func(name string, value any) { got[name] = value })
 want := map[string]any{"online": true, "limit": int64(-42), "language": "en"}
 if !reflect.DeepEqual(got, want) { t.Fatalf("values: %v", got) }
 if reflect.TypeOf(*opts).NumField() != 3 { t.Fatal("read-only option generated") }
 var absent *TDLibOptions
 absent.forEachSet(func(string, any) { t.Fatal("nil options emitted values") })
 opts.forEachSet(nil)
}
`,
	}
	for name, content := range files {
		if err := os.WriteFile(name, []byte(content), 0600); err != nil {
			t.Fatal(err)
		}
	}
	cmd := exec.Command("go", "test", ".")
	cmd.Env = append(os.Environ(), "GOWORK=off")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated options: %v\n%s", err, output)
	}
}
