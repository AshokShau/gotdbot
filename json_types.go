package gotdbot

import (
	"bytes"
	"encoding/json"
	"strconv"
)

func isNilOrNullJSON(data []byte) bool {
	if len(data) == 0 {
		return true
	}
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return true
	}
	return bytes.Equal(trimmed, []byte("null"))
}

// Int64Slice is a slice of int64 that is marshalled to/from a slice of strings in JSON.
// This is necessary because TDLib sends vector<int64> as a list of strings.
type Int64Slice []int64

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *Int64Slice) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*s = nil
		return nil
	}
	var strSlice []string
	if err := json.Unmarshal(data, &strSlice); err != nil {
		return err
	}
	*s = make(Int64Slice, len(strSlice))
	for i, str := range strSlice {
		val, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		(*s)[i] = val
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (s *Int64Slice) MarshalJSON() ([]byte, error) {
	if s == nil || *s == nil {
		return []byte("null"), nil
	}
	strSlice := make([]string, len(*s))
	for i, val := range *s {
		strSlice[i] = strconv.FormatInt(val, 10)
	}
	return json.Marshal(strSlice)
}
