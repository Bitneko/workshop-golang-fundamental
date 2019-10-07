### Goroutine

	####  Creating a GoRoutine

In the snippet below, `count("chair")` is a blocking call and will keep printing `chair`, and the program will never get to `count("table")` .
```
func main() {
	count("chair")
	count("table")
}

func count(object string) {
	for i :=1; true; i++ {
		fmt.Println(i, object)
		time.Sleep(time.Millisecond * 300)
	}
}
```

Adding the keyword `go` in front of `count("chair")` and cause it to be run as a GoRoutine, a GoRoutine is non-blocking and the `main` function will continue to run the rest of the program.

Note that the  `main` function in golang is a GoRoutine by itself. When you create  an application with Golang, you are starting with a single GoRoutine that handle the execution of the entire application.
```
func main() {
	go count("chair")
	count("table")
}
```

In the following snippet, there is no output when we run it. This is because a Golang program exit itself once it's done executing the code inside. Since a GoRoutine call is non-blocking, the `main` function which is also a GoRoutine will exit once the 2 `go` call are made, giving no time for the `count` method to execute.

One way to get around this is by adding a blocking call at the end of the `main` function, so that it does not exit immediately.
```
func main() {
	go count("chair")
	go count("table")
}
```

In the following example , the 2 `count` calls will keep running until a manual input is detected by `fmt.Scanln()`
** Doesn't work in playground
```
func main() {
	go count("chair")
	go count("table")
	fmt.Scanln()
}
```

#### WaitGroup
Using `fmt.Scanln()` to block a program is not a very optimise way to block the `main` function as it requires manual input.

`WaitGroup` from the `sync` package that provide basic synchronisation primitives on the other hand give us a more elegant way of managing `GoRoutine`. It is essentially a counter that you increment before spinning a GoRoutine and decrement after the GoRoutine is done running.
```
func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(1) // Increment 1 as we have 1 GoRoutine that we are going to call
	
	go func() {
		count("chair")
		// Decrement the waitgroup counter by 1
		// This must be called within a GoRoutine
		waitgroup.Done() 
	}()
	
	waitgroup.Wait() // Block the function as long as the value is above 0
}

func count(object string) {
	for i :=1; i < 5; i++ {
		fmt.Println(i, object)
		time.Sleep(time.Millisecond * 300)
	}
}
```

#### Channel
Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`. It is the de facto mean of communication between GoRoutines in Go.

```
func main() {
	c:= make(chan string)
	go count("pen", c)
	msg := <-c // receiver
	fmt.Println(msg)
}

func count(object string, c chan string) {
	for i :=0; i < 5; i++ {
		c <- object //sender
		time.Sleep(time.Millisecond * 300)
	}
}
```
The last example will output `pen` once and exit. To print all 5 instances of `pen`, you can use a `for` loop to have the receiver keep receiving message from the GoRoutine.
```
func main() {
	c:= make(chan string)
	go count("pen", c)

	for {
		msg := <-c
		fmt.Println(msg)
	}
}
```

##### Deadlock
In the last example, after printing 5 `pen`, a `deadlock` error would happen. This is because although the `count` function has finished executing, the receiver is still waiting to receive from the channel.

To solve this, the channel need to be closed by the sender.

Channel should always be closed by the sender because the receiver would not know when to close the channel. If a channel is closed prematurely, it would cause an `panic` error when the sender try to send a value to the closed channel.
```
func count(object string, c chan string) {
	for i :=1; i < 5; i++ {
		c <- object //sender
		time.Sleep(time.Millisecond * 300)
	}
	close(c)
}
```

Channel also return a second value that show if the channel is closed or open
```
func main() {
	c:= make(chan string)
	go count("pen", c)

	for {
		msg, open := <-c
		if !open {
			break // Break out of the for loop
		}
		fmt.Println(msg)
	}
}
```

You can use `range` to iterate through the values of a channel instead of looping through it with a `for` loop.
```
func main() {
	c:= make(chan string)
	go count('pen', c)
	msg := <-c // receiver
	for msg := range c {
		fmt:Println(msg)
	}
}
```
##### Buffer Channel
You cannot send a value to a channel before a receiver is declared to receive the value. Doing so will cause an `deadlock` error as the sender will block the program execution preventing the receiver from being declared.

To get around this, you can declare a capacity to the channel as a buffer. Note that if the value sent to the channel exceed the buffer, the `deadlock` will be back.

```
func main() {
	c:= make(chan string)
	c <- "hello" // Sender blocks the program
	msg := <-c
	fmt.Println(msg)
}
```
```
func main() {
	c:= make(chan string, 2)
	c <- "hello"
	c <- "world"
	msg := <-c
	fmt.Println(msg)
}
```
##### Select Channel
The example below prints `chair` and `table` in sequence even though `table` are set to send values to it's channel at a much faster rate. This is because a receiver is a blocking call, so `fmt.Println(<-c1)` will block after receiving a single value, and will only resume receiving new value after `fmt.Println(<-c2)` is done receiving it's value.
```
func main() {
	c1:= make(chan string)
	c2:= make(chan string)
	
	go count("chair", 500, c1)
	go count("table", 200, c2)
	
	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}

func count(object string, delay time.Duration, c chan string) {
	for i :=1; true; i++ {
		c <- object //sender
		time.Sleep(time.Millisecond * delay)
	}
}
```
`select` channel let us receive values from channels whenever they are ready.
```
func main() {
	c1:= make(chan string)
	c2:= make(chan string)
	
	go count("chair", 500, c1)
	go count("table", 200, c2)
	
	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
	
	for {
		select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
		}
	}
}
```

####  Worker Pool
We'll look at how to implement a work pool using **GoRoutines** and **channels**. A worker pool is simply a pool of GoRoutines that work together to compute a process a list of tasks.

Note that worker pool **does not guarantee** the order of which the task completed.
```
func main() {
	tasks := make(chan int, 50)
	results := make(chan int, 50)
	
	go worker(tasks, results)
	
	for i := 0; i < 50; i++ {
		tasks <- i
	}
	close(tasks)
	for j := 0; j < 50; j++ {
		fmt.Println(<-results)
	}
}

func worker(tasks <- chan int, results chan <- int) {
	for t := range tasks {
		results <- generateFibonacci(t)
	}
}

func generateFibonacci(num int) int {
	if num <= 1 {
		return num
	}
	
	return  generateFibonacci(num - 1) + generateFibonacci(num - 2)
}
```
