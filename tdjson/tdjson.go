package tdjson

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
)

var (
	tdCreateClientId func() int32
	tdSend           func(int32, *byte)
	tdReceive        func(float64) uintptr
	tdExecute        func(*byte) uintptr

	libLoaded bool
)

// Init initializes the TDLib JSON interface by loading the library.
// If libPath is empty, it attempts to load the library from the default system paths
// or the current directory using the default name (libtdjson.so, libtdjson.dylib, tdjson.dll).
func Init(libPath string) error {
	if libLoaded {
		return nil
	}

	if libPath == "" {
		libPath = getDefaultLibName()
	}

	lib, err := purego.Dlopen(libPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return fmt.Errorf("failed to load tdjson library from %s: %w", libPath, err)
	}

	purego.RegisterLibFunc(&tdCreateClientId, lib, "td_create_client_id")
	purego.RegisterLibFunc(&tdSend, lib, "td_send")
	purego.RegisterLibFunc(&tdReceive, lib, "td_receive")
	purego.RegisterLibFunc(&tdExecute, lib, "td_execute")

	libLoaded = true
	return nil
}

func getDefaultLibName() string {
	switch runtime.GOOS {
	case "windows":
		return "tdjson.dll"
	case "darwin":
		return "libtdjson.dylib"
	default:
		return "libtdjson.so"
	}
}

// CreateClientID returns an opaque identifier of a new TDLib instance.
// The TDLib instance will not send updates until the first request is sent to it.
func CreateClientID() int {
	ensureInitialized()
	return int(tdCreateClientId())
}

// Send sends request to the TDLib client. May be called from any thread.
func Send(clientID int, request string) {
	ensureInitialized()
	reqBytes := append([]byte(request), 0)
	tdSend(int32(clientID), &reqBytes[0])
}

// Receive receives incoming updates and request responses.
// Must not be called simultaneously from two different threads.
// Returns a JSON-serialized update or request response, or an empty string if the timeout expires.
func Receive(timeout float64) string {
	ensureInitialized()
	ptr := tdReceive(timeout)
	if ptr == 0 {
		return ""
	}
	return goString(ptr)
}

// Execute synchronously executes a TDLib request.
// Only a few requests can be executed synchronously.
// Returns a JSON-serialized request response.
func Execute(request string) string {
	ensureInitialized()
	reqBytes := append([]byte(request), 0)
	ptr := tdExecute(&reqBytes[0])
	if ptr == 0 {
		return ""
	}
	return goString(ptr)
}

func ensureInitialized() {
	if !libLoaded {
		if err := Init(""); err != nil {
			panic(fmt.Sprintf("tdjson not initialized: %v", err))
		}
	}
}

func goString(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}

	var length int
	for {
		b := *(*byte)(unsafe.Pointer(ptr + uintptr(length)))
		if b == 0 {
			break
		}
		length++
	}
	if length == 0 {
		return ""
	}

	bytes := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), length)
	return string(bytes)
}
