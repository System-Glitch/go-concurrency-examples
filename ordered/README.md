Here, there is no race condition and the order is guaranteed. This is because we only modify the variable from the main goroutine et because we are using channels.

```
$ go run -race ordered.go
57
```

The int to string conversion is done in a parallel fashion, but the variable modification is not.