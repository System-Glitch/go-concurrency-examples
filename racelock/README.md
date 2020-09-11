Example with a lock (mutex):
```
$ go run -race racelock.go 
firstsecond
$ go run -race racelock.go 
secondfirst
```

The order is not guaranteed, but there is no race condition anymore.