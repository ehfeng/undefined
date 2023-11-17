# undefined

This kind of sucks.

`null.String` lets us more easily distinguish between zero and null JSON values, but doesn't distinguish betweeen null and undefined fields. Leaving out a field can be significant though, especially as I've defined in [JSON SET](https://www.ehfeng.com/json-set/).

`omitempty` [won't work](https://github.com/golang/go/issues/11939) for marshalling custom types but will hopefully be included in the v2 of [json marshalling](https://github.com/golang/go/discussions/63397).

> "omitempty" option was narrowly defined as only omitting a field if it is a Go false, 0, a nil pointer, a nil interface value, and any empty array, slice, map, or string.

```go
package main

import (
    "gopkg.in/guregu/null.v4"
    "github.com/ehfeng/undefined"
)

type A struct {
    X string           `json:"x,omitempty"`
    Y null.String      `json:"y,omitempty"`
    Z undefined.String `json:"z,omitempty"`
}

func (a *A) B() B {
    var ptr *undefined.String
    if a.Z.Defined {
        ptr = &a.Z
    }
    return B{
        X: a.X,
        Y: a.Y,
        Z: p,
    }
}

// this secondary struct must be used for _marshalling_ because 
// golang json only tests against native zero values for omitempty
type B struct {
    X string
    Y null.String
    Z *undefined.String `json:"z,omitempty"`
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
