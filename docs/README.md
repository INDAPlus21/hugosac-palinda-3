# Task 1 - Matching Behaviour

### What happens if you remove the `go-command` from the `Seek` call in the `main` function?
**Hypothesis**: The `Seek` calls will run in the main goroutine in order. Because the channel is buffered it won't have to wait for another goroutine to read its value. The next call to `Seek` will read the channel's value and it will work fine. The messages will be sent in the same order as the names are declared in the `people` slice.

**Result**: The program still works and Anna sends to Bob, Cody sends to Dave and no one receives Eva message.

### What happens if you switch the declaration `wg := new(sync.WaitGroup)` to `var wg sync.WaitGroup` and the parameter `wg *sync.WaitGroup` to `wg sync.WaitGroup`?
**Hypothesis**: The declaration switch will change the `wg` variable to a `WaitGroup` instead of a pointer reference to a `WaitGroup`. The second switch will pass a copy of the `WaitGroup` instead of a pointer reference. The `Done` function is therefore called on a copy of the `WaitGroup` which will result in the original `WaitGroup` waiting forever. The execution therefore won't go past the `wg.Wait()` statement. To print statements of communication should be printed and the program should then enter a deadlock.

**Result**: Two lines of communication are printed and the program thereafter enters a deadlock.


### What happens if you remove the buffer on the channel match?
**Hypothesis**: After the last send, the channel is waiting for someone to read the message and the program will reach a deadlock. When the channel is buffered it can keep going and the value can be read later on.

**Result**: The program reaches a deadlock after the last message is written to the channel.


### What happens if you remove the default-case from the case-statement in the `main` function?
**Hypothesis**: Nothing should happen. There should be no difference at all. If there is no match with the first case the program should skip the select statement. As there is no code to execute in the default-case, there is no need for it.

**Result**: The program works as normal.


# Task 2 - Fractal Images

One function was removed and a new one added.

Five tests were ran on both the single threaded and multi threaded versions and their execution times are listed below. The average times rounded to decimal points are presented. The parallel version uses 8 CPUs on my computer. Ran on a Dell XPS 13 9310 with an Intel i7-1185G7 with 16GB of RAM.

**Single threaded**: 16,22 s

**Multi threaded**: 3,01 s
