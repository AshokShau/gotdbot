package gotdbot

import (
	"encoding/json"
	"testing"
)

func BenchmarkSendConcurrent(b *testing.B) {
	c := newTestClient(b)
	c.sendRequest = func(_ int, data []byte) {
		var wire testWireRequest
		if err := json.Unmarshal(data, &wire); err != nil {
			b.Error(err)
			return
		}
		c.deliverResponse(wire.Extra, &User{Id: 42})
	}
	req := &GetMe{}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := c.Send(req); err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkUpdateQueue(b *testing.B) {
	for _, capacity := range []int{0, 64} {
		name := "memory"
		if capacity == 0 {
			name = "disk"
		}
		b.Run(name, func(b *testing.B) {
			q := newUpdateQueue(capacity)
			defer q.close()
			obj := &UpdateNewMessage{Message: testMessage(42, "hello")}
			b.ReportAllocs()
			for b.Loop() {
				if err := q.push(obj); err != nil {
					b.Fatal(err)
				}
				if _, err := q.take(); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkUnmarshalMessage(b *testing.B) {
	data, err := json.Marshal(&UpdateNewMessage{Message: testMessage(42, "hello 😀")})
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.SetBytes(int64(len(data)))
	for b.Loop() {
		if _, _, err := Unmarshal(data); err != nil {
			b.Fatal(err)
		}
	}
}
