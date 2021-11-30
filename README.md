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

### Go's Concurrency Building Blocks

1. Goroutines:-

***Goroutines are unique to GO (though some other languages have a concurrency primitive that is similar).***
***They are not OS threads, and they are not exactly green threads--threads that are managed by a language's runtime--they are higher level of abstraction known as coroutines. Corotines are simply concurrent subroutines(functions, closures, or methods in Go) that are nonpreemptive--that is they cannot be interrupted. Instead, coroutines have multiple points throughout which allow for suspension or reentry***
```
1. What makes goroutines unique to Go are their deep integration with Go's runtime.
2. Go's runtime observes the runtime behavior of goroutines and automatically suspends them when they block and then resumes them when they become unblocked.
3. Goroutines can be considered a special class of coroutine.
```
***Go follows a model of concurrency called the fork-join model. The word fork refers to the fact that at any point in the program, it can split off a child branch of execution to be run concurrently with its parent. The word join refers to the fact that at some point in the future, these concurrent branches of execution will join back together.***
***Where the child rejoins the parent is called a join point***

Here is a graphical representation to help you picture it
![image](https://user-images.githubusercontent.com/21126970/144104737-86b72002-6b2a-4504-9043-70d20a0848de.png)


