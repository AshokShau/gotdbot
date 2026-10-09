package tdjson

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

var (
	tdCreateClientId func() int32
	tdSend           func(int32, *byte)
	tdReceive        func(float64) *byte
	tdExecute        func(*byte) *byte

	libLoaded bool
	initMu    sync.Mutex
	resultMu  sync.Mutex
)

// Init initializes the TDLib JSON interface by loading the library.
func Init(libPath string) (err error) {
	initMu.Lock()
	defer initMu.Unlock()
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

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("invalid TDLib library %s: %v", libPath, r)
		}
		if !libLoaded {
			_ = purego.Dlclose(lib)
		}
	}()

	purego.RegisterLibFunc(&tdCreateClientId, lib, "td_create_client_id")
	purego.RegisterLibFunc(&tdSend, lib, "td_send")
	purego.RegisterLibFunc(&tdReceive, lib, "td_receive")
	purego.RegisterLibFunc(&tdExecute, lib, "td_execute")

	libLoaded = true

	// disables internal TDLib logging
	req := `{"@type": "setLogStream", "log_stream": {"@type": "logStreamEmpty"}}`
	Execute(req)
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
func CreateClientID() int {
	return int(tdCreateClientId())
}

// Send sends request to the TDLib client. May be called from any thread.
func Send(clientID int, request string) {
	reqBytes := append([]byte(request), 0)
	tdSend(int32(clientID), &reqBytes[0])
}

// SendBytes sends raw JSON request bytes to the TDLib client.
func SendBytes(clientID int, request []byte) {
	reqBytes := append(request, 0)
	tdSend(int32(clientID), &reqBytes[0])
}

// Receive receives incoming updates and request responses.
// Returns a JSON-serialized update or request response, or an empty string if the timeout expires.
func Receive(timeout float64) string {
	resultMu.Lock()
	defer resultMu.Unlock()
	ptr := tdReceive(timeout)
	if ptr == nil {
		return ""
	}
	return goString(ptr)
}

// Execute synchronously executes a TDLib request.
// Returns a JSON-serialized request response.
func Execute(request string) string {
	resultMu.Lock()
	defer resultMu.Unlock()
	reqBytes := append([]byte(request), 0)
	ptr := tdExecute(&reqBytes[0])
	if ptr == nil {
		return ""
	}
	return goString(ptr)
}

func goString(ptr *byte) string {
	if ptr == nil {
		return ""
	}

	var length int
	for {
		b := *(*byte)(unsafe.Add(unsafe.Pointer(ptr), length))
		if b == 0 {
			break
		}
		length++
	}
	if length == 0 {
		return ""
	}

	bytes := unsafe.Slice(ptr, length)
	return string(bytes)
}
