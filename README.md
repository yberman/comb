# comb
combinatorics in golang

## Development Notes

### Initialization

All objects defined should be fully functional
with curly brace initialization

```go
b := Blah{}
b.FuncFoo(argBuzz)
```

### Method defintion

All objects which are structs
must define a private idempotent init method
to be used for defining other methods.

```go
type Blah struct {
    // stuff
}

func (b *Blah) FuncFoo(argBuzz Buzz) {
    b.init()
    // other stuff
}
```

### Pointer Receivers

Be consistent with wether the receivers are
pointer receivers or not.
