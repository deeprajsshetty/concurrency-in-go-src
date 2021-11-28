This is the full sourcecode for the book, "Concurrency in Go" published by O'Reilly.

For errata and more information, please refer to the book's hub at [[http://katherine.cox-buday.com/concurrency-in-go]].

### Difference Between Concurrency and Parallelism

***Concurrency is a property of the code; parallelism is a property of the running program.***
```
1. Main point is we don't write parallel code, only concurrent code that we hope will be run in parallel.
2. Parallelism mainly depend on Cores.
3. Go’s philosophy on concurrency can be summed up like this: 
   aim for simplicity,
   use channels when possible, and 
   treat goroutines like a free resource.
4. One of Go’s mottos is “Share memory by communicating, don’t communicate by sharing memory.”
5. CSP - Communicating Sequential Processes >> Go built on this principle.
```
