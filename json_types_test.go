package gotdbot

import (
	"encoding/json"
	"math"
	"reflect"
	"testing"
)

func TestInt64SliceJSON(t *testing.T) {
	for _, tc := range []struct {
		name, input, wire string
		want              Int64Slice
	}{
		{"null", "null", "null", nil},
		{"empty", "[]", "[]", Int64Slice{}},
		{"mixed", `[0,"-1",9007199254740993,"9223372036854775807","-9223372036854775808"]`, `["0","-1","9007199254740993","9223372036854775807","-9223372036854775808"]`, Int64Slice{0, -1, 9007199254740993, math.MaxInt64, math.MinInt64}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got Int64Slice
			if err := json.Unmarshal([]byte(tc.input), &got); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("decoded = %v, want %v", got, tc.want)
			}
			data, err := json.Marshal(got)
			if err != nil || string(data) != tc.wire {
				t.Fatalf("encoded = %s, %v", data, err)
			}
		})
	}
	for _, input := range []string{`[1.5]`, `["oops"]`, `[9223372036854775808]`, `[-9223372036854775809]`, `[true]`, `{}`, `[1,`} {
		t.Run(input, func(t *testing.T) {
			got := Int64Slice{7, 8}
			if err := json.Unmarshal([]byte(input), &got); err == nil {
				t.Fatal("invalid int64 accepted")
			}
			if !reflect.DeepEqual(got, Int64Slice{7, 8}) {
				t.Fatal("failed decode mutated destination")
			}
		})
	}
}

func TestUnmarshalWithClientNestedInterfaces(t *testing.T) {
	data := []byte(`{"@type":"updateNewMessage","@client_id":7,"@extra":"call-42","message":{"@type":"message","id":9007199254740993,"chat_id":42,"sender_id":{"@type":"messageSenderUser","user_id":43},"content":{"@type":"messageText","text":{"@type":"formattedText","text":"/start 😀","entities":[{"@type":"textEntity","offset":0,"length":6,"type":{"@type":"textEntityTypeBotCommand"}}]}}}}`)
	obj, id, extra, err := UnmarshalWithClient(data)
	if err != nil {
		t.Fatal(err)
	}
	u, ok := obj.(*UpdateNewMessage)
	if !ok || id != 7 || extra != "call-42" {
		t.Fatalf("envelope = (%T,%d,%q)", obj, id, extra)
	}
	if u.Message.Id != 9007199254740993 || u.Message.SenderID() != 43 || u.Message.GetText() != "/start 😀" || !u.Message.IsCommand() {
		t.Fatalf("decoded message = %+v", u.Message)
	}
	wire, err := json.Marshal(obj)
	if err != nil {
		t.Fatal(err)
	}
	decoded, _, err := Unmarshal(wire)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(obj, decoded) {
		t.Fatal("nested interface round trip changed message")
	}
}

func TestUnmarshalMalformedAndUnknownObjects(t *testing.T) {
	for _, input := range []string{``, `null`, `[]`, `{}`, `{"@type":"futureType"}`, `{"@type":1}`, `{"@type":"ok","@extra":1}`, `{"@type":"updateNewMessage","message":{"content":{"@type":"futureContent"}}}`} {
		if obj, _, _, err := UnmarshalWithClient([]byte(input)); err == nil || obj != nil {
			t.Errorf("input %q = (%v,%v)", input, obj, err)
		}
	}
}

func TestNullableClassFieldsRoundTrip(t *testing.T) {
	for _, obj := range []TlObject{
		&Message{Id: 1, ChatId: 42},
		&UpdateNewMessage{Message: testMessage(42, "hello")},
		&UpdateNewCallbackQuery{Payload: nil},
		&UpdateMessageSendSucceeded{Message: &Message{Id: 1}},
		&Messages{Messages: []Message{{Id: 1}, {Id: 2}}},
	} {
		t.Run(obj.GetType(), func(t *testing.T) {
			data, err := json.Marshal(obj)
			if err != nil {
				t.Fatal(err)
			}
			decoded, _, err := Unmarshal(data)
			if err != nil {
				t.Fatalf("nullable fields: %v", err)
			}
			if !reflect.DeepEqual(obj, decoded) {
				t.Fatalf("round trip changed %T", obj)
			}
		})
	}
	for _, input := range []string{"null", " \n\t null \r "} {
		if obj, err := UnmarshalMessageContent([]byte(input)); obj != nil || err != nil {
			t.Fatalf("null content = (%v,%v)", obj, err)
		}
		if obj, err := UnmarshalMessageSendingState([]byte(input)); obj != nil || err != nil {
			t.Fatalf("null state = (%v,%v)", obj, err)
		}
	}
}

func FuzzInt64SliceRoundTrip(f *testing.F) {
	for _, value := range []int64{0, 1, -1, 9007199254740993, math.MinInt64, math.MaxInt64} {
		f.Add(value)
	}
	f.Fuzz(func(t *testing.T, value int64) {
		original := Int64Slice{value, 0, math.MaxInt64, math.MinInt64}
		data, err := json.Marshal(original)
		if err != nil {
			t.Fatal(err)
		}
		var got Int64Slice
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, original) {
			t.Fatalf("lost precision: %v -> %v", original, got)
		}
	})
}

func FuzzInt64SliceDecode(f *testing.F) {
	for _, input := range []string{`null`, `[]`, `["9223372036854775807",-1]`, `[1.5]`, `[true]`, `[null]`} {
		f.Add([]byte(input))
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		got := Int64Slice{42}
		if err := json.Unmarshal(data, &got); err != nil {
			if !reflect.DeepEqual(got, Int64Slice{42}) {
				t.Fatal("failed decode changed destination")
			}
			return
		}
		wire, err := json.Marshal(got)
		if err != nil {
			t.Fatal(err)
		}
		var roundTrip Int64Slice
		if err := json.Unmarshal(wire, &roundTrip); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, roundTrip) {
			t.Fatal("successful decode cannot round trip")
		}
	})
}

func FuzzUnmarshalWithClient(f *testing.F) {
	for _, data := range []string{`{"@type":"ok"}`, `{"@type":"message","id":1,"content":null}`, `{"@type":"updateNewMessage","message":{"id":1}}`, `{"@type":"error","code":400,"message":"bad"}`, `null`, `{`} {
		f.Add([]byte(data))
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		obj, _, _, err := UnmarshalWithClient(data)
		if err != nil {
			return
		}
		if obj == nil || obj.GetType() == "" {
			t.Fatal("successful decode returned no object type")
		}
		wire, err := json.Marshal(obj)
		if err != nil {
			t.Fatal(err)
		}
		decoded, _, err := Unmarshal(wire)
		if err != nil {
			t.Fatalf("decoded object cannot round trip: %v", err)
		}
		if decoded.GetType() != obj.GetType() {
			t.Fatal("round trip changed object type")
		}
	})
}
