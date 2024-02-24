package shelly

import (
	"encoding/json"
	"errors"
	"strings"
)

// NullString is similar to a *string but will encode to the JSON null value if empty, but
// *NullString(nil) will be omitted w/ omitempty flag.
type NullString string

func NewNullString(s string) *NullString {
	n := NullString(s)
	return &n
}

func (n *NullString) String() string {
	return string(*n)
}

func (n *NullString) UnmarshalJSON(b []byte) error {
	s := strings.TrimSpace(string(b))
	if s == "null" {
		return nil
	}
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return errors.New("JSON string type expected")
	}
	*n = NullString(s[1 : len(s)-1])
	return nil
}

func (n *NullString) MarshalJSON() ([]byte, error) {
	if n == nil || *n == "" {
		return []byte("null"), nil
	}
	return json.Marshal(*n)
}

func StrPtr(s string) *string {
	return &s
}

func IntPtr(i int) *int {
	return &i
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func BoolPtr(b bool) *bool {
	return &b
}
