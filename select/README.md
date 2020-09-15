In this example, we are using `select` to receive messages from multiple goroutines in a central goroutine, using a single channel. This kind of mechanism is often used with long-lived workers.

We are also using a context to gracefully terminate our goroutines. We are listening OS signals SIGINT and SIGTERM to properly exit the program.

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