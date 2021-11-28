This is the full sourcecode for the book, "Concurrency in Go" published by O'Reilly.

For errata and more information, please refer to the book's hub at [[http://katherine.cox-buday.com/concurrency-in-go]].

### Difference Between Concurrency and Parallelism
```
1. ***Concurrency is a property of the code; parallelism is a property of the running program.***

```
   > Concurrency is a property of the code; parallelism is a property of the running program.
   > Main point is we don't write parallel code, only concurrent code that we hope will be run in parallel.
   > Parallelism mainly depend on Cores.
   > Go’s philosophy on concurrency can be summed up like this: aim for simplicity, use channels when possible, and treat goroutines like a free resource.
   > One of Go’s mottos is “Share memory by communicating, don’t communicate by sharing memory.”
   > CSP - Communicating Sequential Processes >> Go built on this principle.
