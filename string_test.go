package undefined

import (
	"encoding/json"
	"testing"
)

var (
	nullX      = []byte(`{"x":null}`)
	emptyX     = []byte(`{"x":""}`)
	undefinedX = []byte(`{}`)
)

type stringInStruct struct {
	X String `json:"x,omitempty"`
	Y string `json:"y,omitempty"`
}

func (s stringInStruct) MarshalJSON() ([]byte, error) {
	var ptr *String
	if s.X.Defined {
		ptr = &s.X
	}
	t := struct {
		X *String `json:"x,omitempty"`
		Y string  `json:"y,omitempty"`
	}{
		X: ptr,
		Y: s.Y,
	}
	return json.Marshal(t)
}

func TestUnmarshalStringField(t *testing.T) {
	var err error
	var s1 stringInStruct
	err = json.Unmarshal(nullX, &s1)
	if err != nil {
		t.Error(err)
	}
	if !s1.X.Defined {
		t.Errorf("Expected X to be defined")
	}
	if s1.X.Valid {
		t.Errorf("Expected X to be not valid")
	}

	var s2 stringInStruct
	err = json.Unmarshal(emptyX, &s2)
	if err != nil {
		t.Error(err)
	}
	if !s2.X.Defined {
		t.Errorf("Expected X to be defined")
	}
	if !s2.X.Valid {
		t.Errorf("Expected X to be valid")
	}
	if s2.X.String.String != "" {
		t.Errorf("Expected X to be empty string")
	}

	var s3 stringInStruct
	err = json.Unmarshal(undefinedX, &s3)
	if err != nil {
		t.Error(err)
	}
	if s3.X.Defined {
		t.Errorf("Expected X to be undefined")
	}
	if s3.X.Valid {
		t.Errorf("Expected X to be not valid")
	}
}

func TestMarshal(t *testing.T) {
	s1 := stringInStruct{
		X: NewString("foo", true),
	}
	b1, err := json.Marshal(s1)
	if err != nil {
		t.Error(err)
	}
	if string(b1) != `{"x":"foo"}` {
		t.Errorf("Expected marshaled JSON to be `{\"x\":\"foo\"}` but got `%s`", string(b1))
	}

	s2 := stringInStruct{
		X: NewString("", true),
	}
	b2, err := json.Marshal(s2)
	if err != nil {
		t.Error(err)
	}
	if string(b2) != `{"x":""}` {
		t.Errorf("Expected marshaled JSON to be `{\"x\":\"\"}` but got `%s`", string(b2))
	}

	s3 := stringInStruct{
		X: NewString("", false),
		Y: "hi",
	}
	b3, err := json.Marshal(s3)
	if err != nil {
		t.Error(err)
	}
	if string(b3) != `{"x":null,"y":"hi"}` {
		t.Errorf("Expected marshaled JSON to be `{\"x\":null}` but got `%s`", string(b3))
	}

	s4 := stringInStruct{Y: "hello"}
	b4, err := json.Marshal(s4)
	if err != nil {
		t.Error(err)
	}
	if string(b4) != `{"y":"hello"}` {
		t.Errorf("Expected marshaled JSON to be `{\"y\": \"hello\"}` but got `%s`", string(b4))
	}
}
