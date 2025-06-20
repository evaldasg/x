
# x

A small general-purpose Go helpers library.

## Features

✅ Immutable map (`x.Map`): generic, read-only map with safe access.

More helpers coming soon: immutable lists, sets, strings, math utilities, etc.

## Install

```
go get github.com/evaldasg/x
```

## Example

```go
import "github.com/evaldasg/x"

m := x.NewMap(map[int]string{
    200: "OK",
    404: "Not Found",
})

fmt.Println(m.MustGet(200)) // OK
fmt.Println(m.Get(500))     // "", false
```

## License

MIT
