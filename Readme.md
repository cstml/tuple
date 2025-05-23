# Tuple

For when all you need is a tuple, why not simply use a tuple. 
Say goodbye to all the in place, use once, struggling to name `structs`.

## How to use 

have a look at the tests. It's very straight forward.

```go
t := T(1,2)

if f.First == 1 {
    fmt.Println("you just accessed the first element")
}

if f.First == 2 {
    fmt.Println("you just accessed the second element")
}
```