package tdjson

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
)

func TestGoStringCopiesNullTerminatedBuffer(t *testing.T) {
	if goString(nil) != "" {
		t.Fatal("nil pointer not empty")
	}
	data := []byte("hello😀\x00ignored")
	got := goString(&data[0])
	data[0] = 'X'
	if got != "hello😀" {
		t.Fatalf("native buffer was not copied: %q", got)
	}
	empty := []byte{0}
	if goString(&empty[0]) != "" {
		t.Fatal("empty native string not empty")
	}
}

func TestTransportMarshallingAndResponseLifetime(t *testing.T) {
	oldSend, oldReceive, oldExecute := tdSend, tdReceive, tdExecute
	t.Cleanup(func() { tdSend, tdReceive, tdExecute = oldSend, oldReceive, oldExecute })
	type record struct {
		id      int32
		request string
	}
	var sent []record
	tdSend = func(id int32, p *byte) { sent = append(sent, record{id, goString(p)}) }
	Send(42, `{"@type":"getMe"}`)
	data := []byte(`{"@type":"getChat"}`)
	SendBytes(43, data)
	SendBytes(44, nil)
	if len(sent) != 3 || sent[0].id != 42 || sent[0].request != `{"@type":"getMe"}` || sent[1].request != string(data) || sent[2].request != "" {
		t.Fatalf("send records = %v", sent)
	}
	buffer := []byte("first\x00")
	tdReceive = func(float64) *byte { return &buffer[0] }
	first := Receive(0.1)
	copy(buffer, []byte("other\x00"))
	if first != "first" {
		t.Fatal("Receive retained native response buffer")
	}
	tdReceive = func(float64) *byte { return nil }
	if Receive(0) != "" {
		t.Fatal("nil Receive not empty")
	}
	tdExecute = func(p *byte) *byte {
		if goString(p) != "request" {
			t.Errorf("execute input = %q", goString(p))
		}
		return &buffer[0]
	}
	if Execute("request") != "other" {
		t.Fatal("Execute response corrupted")
	}
	tdExecute = func(*byte) *byte { return nil }
	if Execute("") != "" {
		t.Fatal("nil Execute not empty")
	}
	runtime.KeepAlive(buffer)
}

func TestReceiveAndExecuteSerializeNativeBufferAccess(t *testing.T) {
	oldReceive, oldExecute := tdReceive, tdExecute
	t.Cleanup(func() { tdReceive, tdExecute = oldReceive, oldExecute })
	buffer := make([]byte, 32)
	write := func(value string) *byte { clear(buffer); copy(buffer, value); return &buffer[0] }
	tdReceive = func(float64) *byte { return write("receive") }
	tdExecute = func(*byte) *byte { return write("execute") }
	var wg sync.WaitGroup
	for range 32 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 32 {
				if Receive(0) != "receive" {
					t.Error("Receive corrupted")
				}
				if Execute("test") != "execute" {
					t.Error("Execute corrupted")
				}
			}
		}()
	}
	wg.Wait()
}

func TestInitMissingLibraryConcurrent(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing-tdjson-library.so")
	var wg sync.WaitGroup
	for range 8 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := Init(path); err == nil || !strings.Contains(err.Error(), path) {
				t.Errorf("missing library error = %v", err)
			}
		}()
	}
	wg.Wait()
	if libLoaded {
		t.Fatal("failed initialization marked library loaded")
	}
	if name := getDefaultLibName(); name == "" {
		t.Fatal(fmt.Errorf("missing library name for %s", runtime.GOOS))
	}
}
