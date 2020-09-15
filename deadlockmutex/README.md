In this example, we are protecting a global variable with a mutex. However, we are creating a deadlock because our `validateUsername()` and `updateUsername()` functions both use the mutex. `updateUsername()` calls `validateUsername()` before releasing the lock, so `validateUsername()` waits for the lock to be released forever and we end up with a deadlock.

This could be fixed by using a copy of the new username for validation instead, removing the need for a lock:

```go
func updateUsername(newName string) {
    defer wg.Done()
    if validateUsername(newName) {
	    mu.Lock()
        username = newName
	    fmt.Println("Name updated to", username)
	    mu.Unlock()
    } else {
		panic("You cannot be called Bob")
    }
}

func validateUsername(newName string) bool {
	return newName != "Bob"
}
```