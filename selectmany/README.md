In this example, we are using `select` to receive messages from multiple goroutines and multiple channels in a central goroutine. Note that the ordering is not guaranteed.

We are also using a timeout in case the worker routines are too slow.

```
$ go run -race select.go 
Received message from ch1: 0
Received message from ch2: 0
Received message from ch1: 1
Received message from ch1: 2
Received message from ch2: 1
Received message from ch2: 2
All done!
```