# undefined

`null.String` lets us more easily distinguish between zero and null JSON values, but doesn't distinguish betweeen null and undefined fields. Leaving out a field can be significant though, especially with [JSON SET](https://www.ehfeng.com/json-set/).

This is only useful for unmarshalling json.  `omitempty` [won't work](https://github.com/golang/go/issues/11939) for marshalling structs but will hopefully be included in the v2 of [json marshalling](https://github.com/golang/go/discussions/63397). To marshal `undefined.String` fields, you need to implement `MarshalJSON` on your struct (see below).

> "omitempty" option was narrowly defined as only omitting a field if it is a Go false, 0, a nil pointer, a nil interface value, and any empty array, slice, map, or string.

```go
package main

import (
    "gopkg.in/guregu/null.v4"
    "github.com/ehfeng/undefined"
)

type A struct {
    X string           `json:"x"`
    Y null.String      `json:"y"`
    Z undefined.String `json:"z"`
}

func (a A) MarshalJSON() ([]byte, error) {
    // otherwise, Z will be marshalled as `null`, even when undefined
    var ptr *undefined.String
    if a.Z.Defined {
        ptr = &a.Z
    }
    return struct{
        X string            `json:"x"`
        Y null.String       `json:"y"`
        Z *undefined.String `json:"z,omitempty"`
    }{
        X: a.X,
        Y: a.Y,
        Z: p,
    }
}

func main() {
    var b []byte

    var zeroes A
    json.Unmarshal([]byte(`{"x": "", "y": "", "z": ""}`), &zeroes)
    b, _ := json.Marshal(zeroes)
    fmt.Println(string(b)) // {"y": "", "z": ""}

    var nulls A
    json.Unmarshal([]byte(`{"x": null, "y": null, "z": null}`), &nulls)
    b, _ = json.Marshal(nulls)
    fmt.Println(string(b)) // {"y": null, "z": null}

    var undefineds A
    json.Unmarshal([]byte(`{}`), &undefineds)
    b, _ = json.Marshal(undefineds.B())
    fmt.Println(string(b)) // {"y": null}
}
```
