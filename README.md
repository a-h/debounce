# Debounce Button

Debounce library for Go.

```go
onClick := func() {
 fmt.Print("Clicked...")
}
normallyClosed := true
sw := Button(onClick, normallyClosed)
```
