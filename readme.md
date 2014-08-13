# gochan


## Overview

gochan benchmarks several channel patterns using golang, which helps programmers to figure out how to write an efficient code in second.


#### Stop Pattern
stop pattern is the most channel use cases in golang's concurrent programming. You actually try to do something while listen the stop signal. If it signals, you know it is about time to exit.

For example:

```
go func (){
    for {
        select {
        case <-stopChan:
            // watch the stop signal
            return
        default:
            // keep do something
        }
    }
}()
```
This pattern is asynchronous pattern, which is mainly used to listen multiple channels. It has both high readability and maintainability, but it comes not free. Whereas the following pattern performs 10 times faster than the above pattern:

```
stop := false
go func(){
    <-stopChan
    stop = true
}()

go func(){
    for{
        if stop{
            return
        }
        // keep do something
    }
}

```
In this example, we use additional goroutine to wait on the stop signal, while use another goroutine to do jobs.
It performs extraordinary faster than the select clause which is a known performance bottleneck in go. Search "selectgo" to see how you can find some.

Here is our benchmark results run on mac air:

Pattern | result
------------ | -------------
nonblocking (1st pattern) | 1061923 ns/op
blocking (2nd pattern) | 102518 ns/op

