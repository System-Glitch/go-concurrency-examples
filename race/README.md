Race condition example:
```
$ go run -race race.go 
==================
WARNING: DATA RACE
Read at 0x0000006172b0 by goroutine 8:
  main.concat()
      ./examples/race.go:15 +0x3e

Previous write at 0x0000006172b0 by goroutine 7:
  main.concat()
      ./examples/race.go:15 +0x9b

Goroutine 8 (running) created at:
  main.main()
      ./examples/race.go:23 +0xa4

Goroutine 7 (finished) created at:
  main.main()
      ./examples/race.go:22 +0x74
==================
firstsecond
Found 1 data race(s)
exit status 66
$ go run -race race.go 
==================
WARNING: DATA RACE
Read at 0x0000006172b0 by goroutine 7:
  main.concat()
      ./examples/race.go:15 +0x3e

Previous write at 0x0000006172b0 by goroutine 8:
  main.concat()
      ./examples/race.go:15 +0x9b

Goroutine 7 (running) created at:
  main.main()
      ./examples/race.go:22 +0x74

Goroutine 8 (finished) created at:
  main.main()
      ./examples/race.go:23 +0xa4
==================
secondfirst
Found 1 data race(s)
exit status 66
```

A race condition is detected, and the concatenation doesn't always happen in the same order. There is even a case where the concurrent access can completely nullify one of the two concatenations. We will then only get "second" or "first" on the output.  