// Package null contains SQL types that consider zero input and null input as separate values,
// with convenient support for JSON and text marshaling.
// Types in this package will always encode to their null value if null.
// Use the zero subpackage if you want zero values and null to be treated the same.
package undefined

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"

	"gopkg.in/guregu/null.v4"
)

// nullBytes is a JSON null literal
var nullBytes = []byte("null")

// String is a nullable string. It supports SQL and JSON serialization.
// It will marshal to null if null. Blank string input will be considered null.
type String struct {
	Defined bool
	null.String
}

// StringFrom creates a new String that will never be blank.
func StringFrom(s string) String {
	return NewString(s, true)
}

// StringFromPtr creates a new String that be null if s is nil.
func StringFromPtr(s *string) String {
	if s == nil {
		return NewString("", false)
	}
	return NewString(*s, true)
}

// ValueOrZero returns the inner value if valid, otherwise zero.
func (s String) ValueOrZero() string {
	if !s.Valid {
		return ""
	}
	return s.String.String
}

// NewString creates a new String
func NewString(s string, valid bool) String {
	return String{
		Defined: true,
		String: null.String{
			NullString: sql.NullString{
				String: s,
				Valid:  valid,
			},
		},
	}
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string and null input. Blank string input does not produce a null String.
func (s *String) UnmarshalJSON(data []byte) error {
	s.Defined = true
	if bytes.Equal(data, nullBytes) {
		s.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &s.String); err != nil {
		return fmt.Errorf("null: couldn't unmarshal JSON: %w", err)
	}

	s.Valid = true
	return nil
}

// IsZero returns true for null strings, for potential future omitempty support.
func (s String) IsZero() bool {
	return !s.Defined
}

// Equal returns true if both strings have the same value or are both null.
func (s String) Equal(other String) bool {
	return s.Defined == other.Defined && s.Valid == other.Valid && (!s.Valid || s.String == other.String)
}
