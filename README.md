# Goroutines

### Basics of Goroutines in golang

`wg` is a `sync.WaitGroup` variable used for synchronization. Let's break down the purpose of each of the methods associated with `sync.WaitGroup`:

- `wg.Add(n int)`: This method is used to add `n` to the internal counter of the `WaitGroup`. The counter represents the number of goroutines that are yet to complete. In the code, `1` is added to the `WaitGroup` for each URL before launching a goroutine to fetch the status of that URL.

- `wg.Done()`: This method decrements the `WaitGroup` counter by `1`. It is typically called at the end of the goroutine to signal that the goroutine has completed its execution.

- `wg.Wait()`: This method blocks until the internal counter becomes zero. It waits for all the goroutines to call `Done()` and decrement the counter to zero. In the main function, after launching all the goroutines, `wg.Wait()` is used to wait for all of them to complete before proceeding with the rest of the program.

