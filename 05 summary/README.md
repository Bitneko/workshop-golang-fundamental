# Golang Introduction

## What is it?
Golang is a statically typed compiled programming language 

Statically typed programming language has the following advantages over a dynamically typed language.
* Better code completion
* Type constraints offer more opportunities for compiler optimisations which give better performances
* Reduce the likelihood of some kinds of errors due to type mismatch at runtime

It is syntactically similar to C, and is designed with the following in mind
* Static typing and run-time efficiency (like C++)
* Readability and usability (like Python or JavaScript)
* High-performance networking and multiprocessing

It is created with memory safety,  garbage collection,  structural typing, and **CSP** style concurrency.

[Type Systems: Structural vs. Nominal typing explained](https://medium.com/@thejameskyle/type-systems-structural-vs-nominal-typing-explained-56511dd969f4)
[Introduction to Communicating Sequential Processes(CSP)](https://www.youtube.com/watch?v=G9ePu0Nh2BQ)

## Composition over inheritance in Go
> Golang favour composition over inheritance.

Go does not support inheritance, however it does support composition. The generic definition of composition is "put together".

Struct and interface can be embedded to compose new struct, this is known as Embedding composition.

A struct is a data structure which define a physically grouped list of variables.

An interface is a collection of method signatures that an object can implicitly implement.

The method signatures declared in a interface form the contract of the interface to it’s concrete types and must be implemented.


## Concurrency vs Parallelism

> **Concurrency** is dealing with multiple things at once, 
> **Parallelism** is doing multiple things at once
> -- Rob Pike

### Concurrency

Concurrency is doing 1 things at a time. The CPU time are divided among tasks that need to be processed. Hence we get a sensation of multiple things happening at the same time where in reality, only one thing is happening at a time. 

Go achieve concurrency by dividing the workload based on the priority of each task in a single core processor with **Goroutine**.

<small>We can modify the Go program to run goroutines on different processor cores to achieve parallelism, but that do not necessarily give you better performance. This is due to the latency introduced from communicating on channels.</small>


### Parallelism

Running tasks in parallel form the concept of parallelism. When the CPU has multiple core, we use the different core to do multiple things at once.


### Computer Process
A process, in a nutshell, is a program running in memory.

Binary instructions in programs are compiled into machine code and sent to the OS to handle as a process. The OS allocates things like memory address space(where process’s heap and stacks will be located), program counter, PID(*process id*) and other crucial things to the process.

A process has at least a primary thread which could spawn multiple other threads. When the primary thread finish execution, the process exits.

 #### Thread
A thread is a light-weight process running in a computer process. A thread is what execute a piece of code and it has access to memory provided by the computer process, OS resources, and other things.

#### Memory stack
Thread store data inside the memory stack. A stack is created at compile time and is normally of a fixed size of ~1-2MB. The stack of a thread is only used by that thread, and not shared with other thread.

#### Memory Heap
A memory heap is a property of a process and, unlike stack, is available to use by any thread in a process. Heap is a shared memory space where data from one thread are accessible by other threads

#### Thread scheduling
When multiple threads are running in series or in parallel, threads need to work in coordination so that only one thread can access a particular data at any one time. This problem is called **race condition**.

Execution of multiple threads in some order is called scheduling to prevent race condition is called **scheduling**. OS threads are scheduled by the kernel and some threads are managed by runtime environment of the programming language, like JRE.

### Go Concurrency

When we run a Go program, Go runtime will create a few threads on a core on which all the goroutines are multiplexed (spawned). At any point in time, one thread will be executing one goroutine, and if that goroutine is blocked, then it will be swapped out for another goroutine that will be executed on that thread instead. This is like thread scheduling but handled by Go runtime and this is much faster.

In a nutshell, goroutine is an abstraction over threads. And you can use the `go` keyword to create goroutines.

```
go func runThisInGoroutine() {
  ...
}
```

A Goroutine has the following advantage over a OS thread

- Stack size are allocated dynamically by the go runtime, compared to a 1-2mb fixed stack size for a OS kernel thread, and also starts with only 2KB of stack space. This mean that you can spawn a lot more Goroutine without any problem.

- Goroutine uses **channel** to communicate with other Goroutine with low latency. With threads, there is usually huge latency in inter-thread communication.

- Goroutine has a very cheap setup and teardown cost as it's managed by the Go runtime which already maintain a pool of threads for Goroutines. Threads itself have significant setup and teardown cost because each thread has to request resources from OS and return once it's done.

- Switching between Goroutines is more efficient than thread as they run on 1 thread per Goroutine, and are cooperatively scheduled, so another is not scheduled until current Goroutine is unblocked. Scheduler typically needs to save/restore more than 50 registers switching between threads, compared to just 3 registers(stack pointer, program counter and data register) with Goroutine.


[Goroutines vs Threads](http://tleyden.github.io/blog/2014/10/30/goroutines-vs-threads/)
[X86 Assembly/X86 Architecture](https://en.wikibooks.org/wiki/X86_Assembly/X86_Architecture)


## Working with Go

#### CMD Folder
A convention where you can have multiple binaries when putting them into a sub-folder which is not possible in the root folder.

This is useful in saving time writing boilerplate code by making cleaner abstraction of common code logic.

The CMD folder is not enforced by the Go compiler.  
  
#### Internal Folder  
This is where you put code that’s specific to your project that you don’t want others to import.

The internal folder convention is enforced by the Go compiler.

#### Vendor Folder  
Similar to the node_modules folder of NPM. This is where dependencies are installed to when using `go_mod` or other dependencies management package.
  
#### Go Flag
Command line flag parsing. Useful when designing Go application to take in parameters from the command line on program start .

#### Makefile  
Useful in streamlining repetitive long commands with arguments, or chaining multiple commands into one.

[Using Go modules](https://blog.golang.org/using-go-modules)
[Go Makefile](https://sohlich.github.io/post/go_makefile/)
