# undefined

`null.String` lets us more easily distinguish between zero and null JSON values, but doesn't distinguish betweeen null and undefined fields. Leaving out a field can be significant though, especially as I've defined in [JSON SET](https://www.ehfeng.com/json-set/).

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
    json.Unmarshal([]byte(`{}`), *undefineds)
    b, _ = json.Marshal(undefineds)
    fmt.Println(string(b)) // {"y": null}
}
```
