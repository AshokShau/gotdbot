package gotdbot

import (
	"encoding/json"
	"strconv"
)

// Int64Slice is a slice of int64 that is marshalled to/from a slice of strings in JSON.
// This is necessary because TDLib sends vector<int64> as a list of strings.
type Int64Slice []int64

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *Int64Slice) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*s = nil
		return nil
	}
	var numbers []json.Number
	if err := json.Unmarshal(data, &numbers); err != nil {
		return err
	}
	values := make(Int64Slice, len(numbers))
	for i, number := range numbers {
		val, err := number.Int64()
		if err != nil {
			return err
		}
		values[i] = val
	}
	*s = values
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (s Int64Slice) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte("null"), nil
	}
	strSlice := make([]string, len(s))
	for i, val := range s {
		strSlice[i] = strconv.FormatInt(val, 10)
	}
	return json.Marshal(strSlice)
}
