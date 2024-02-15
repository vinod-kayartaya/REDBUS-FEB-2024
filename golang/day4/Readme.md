## Concurrency

- [Introduction to concurrency in Go](#Introduction_to_concurrency_in_Go)
- [Goroutines and their creation](#Goroutines_and_their_creation)
- [Channels and communication between Goroutines](#Channels_and_communication_between_Goroutines)
- [Buffered channels](#Buffered_channels)
- [Select statement for managing multiple channels](#Select_statement_for_managing_multiple_channels)
- [Synchronization using WaitGroups](#Synchronization_using_WaitGroups)
- [Mutexes and RWMutexes for safe concurrent access](#Mutexes_and_RWMutexes_for_safe_concurrent_access)

<div id='Introduction_to_concurrency_in_Go'></div>

## Introduction to concurrency in Go

Concurrency in Go is a fundamental aspect of the language's design, empowering developers to create highly efficient and scalable programs. It is achieved through goroutines, lightweight threads managed by the Go runtime, and channels, which facilitate communication and synchronization between goroutines.

### Goroutines:

1. **Lightweight Threads**: Goroutines are lightweight, user-space threads managed by the Go runtime. They are cheaper to create and require less memory compared to traditional operating system threads.
2. **Concurrency**: Goroutines allow concurrent execution of functions. They enable developers to write code that can perform multiple tasks simultaneously, improving efficiency and responsiveness.
3. **goroutine Keyword**: Creating a goroutine is as simple as prefixing a function call with the `go` keyword. This instructs the Go runtime to execute the function concurrently.

   ```go
   func main() {
       go func() {
           // Concurrent task
       }()
       // Other tasks
   }
   ```

### Channels:

1. **Communication**: Channels facilitate communication and synchronization between goroutines. They provide a safe way for goroutines to share data by passing messages.
2. **Blocking Operations**: Sending and receiving data through channels are blocking operations. A send operation blocks until a receiver is ready, and a receive operation blocks until data is available.
3. **Unbuffered Channels**: By default, channels are unbuffered, meaning they have no capacity to store data. This ensures synchronization between senders and receivers.

   ```go
   ch := make(chan int) // Create an unbuffered channel
   ```

4. **Buffered Channels**: Channels can also be buffered, allowing a certain number of values to be stored without a corresponding receiver. This can decouple senders and receivers, but be careful of potential deadlocks if the buffer fills up.

   ```go
   ch := make(chan int, 10) // Create a buffered channel with capacity 10
   ```

### Select Statement:

1. **Multi-Channel Communication**: The `select` statement allows you to wait on multiple channel operations simultaneously. It's useful for coordinating communication between multiple goroutines.
2. **Non-Blocking Operations**: If multiple channels are ready, `select` will choose one at random. It can also handle cases where no channel operations are ready without blocking.

   ```go
   select {
   case <-ch1:
       // Handle data from ch1
   case <-ch2:
       // Handle data from ch2
   default:
       // If no channel operation is ready
   }
   ```

### Concurrency Patterns:

1. **Fan-In**: Multiple goroutines can send data to a single channel, known as fan-in. This is useful for aggregating data from multiple sources.
2. **Fan-Out**: A single goroutine can send data to multiple channels, known as fan-out. This is useful for distributing work across multiple workers.
3. **Worker Pools**: Worker pools utilize a fixed number of goroutines to process tasks from a work queue. This pattern is commonly used for parallelizing tasks with a limited resource pool.

Concurrency in Go is a powerful feature that enables developers to write highly efficient and scalable programs. Goroutines and channels provide a simple yet effective way to achieve concurrency, allowing for the creation of responsive and parallelized applications. Understanding these concepts and patterns is essential for building robust concurrent software in Go.

<div id='Goroutines_and_their_creation'></div>

## Goroutines and their creation

Goroutines are lightweight threads managed by the Go runtime. They allow concurrent execution of functions and are a fundamental building block of concurrent Go programs. Goroutines are more efficient than traditional operating system threads, as they are multiplexed onto a smaller number of OS threads. Creating a goroutine is simple, and it's done using the `go` keyword followed by a function call.

Here's a detailed explanation of goroutines with examples:

### Goroutines:

1. **Creation**:

   - Goroutines are created using the `go` keyword followed by a function invocation.
   - The function is executed concurrently in its own goroutine, independently of the main program flow.

   Example:

   ```go
   package main

   import (
       "fmt"
       "time"
   )

   func sayHello() {
       fmt.Println("Hello, Goroutine!")
   }

   func main() {
       go sayHello() // Create a goroutine
       time.Sleep(100 * time.Millisecond) // Wait for goroutine to finish
       fmt.Println("Main function")
   }
   ```

2. **Concurrency**:

   - Goroutines enable concurrent execution of functions, allowing multiple tasks to be performed simultaneously.
   - They are lightweight and have a small overhead, making it feasible to create thousands or even millions of goroutines in a single Go program.

   Example:

   ```go
   package main

   import (
       "fmt"
       "sync"
   )

   func printNumbers() {
       for i := 0; i < 5; i++ {
           fmt.Println(i)
       }
   }

   func main() {
       go printNumbers() // Concurrently print numbers
       fmt.Println("Main function")
   }
   ```

3. **Anonymous Functions**:

   - Goroutines can execute anonymous functions, which are functions without a name defined inline.
   - This is a common pattern for concurrent execution of short tasks.

   Example:

   ```go
   package main

   import (
       "fmt"
       "time"
   )

   func main() {
       go func() {
           for i := 0; i < 5; i++ {
               fmt.Println("Goroutine:", i)
           }
       }() // Anonymous function as a goroutine
       time.Sleep(100 * time.Millisecond) // Wait for goroutine to finish
       fmt.Println("Main function")
   }
   ```

4. **Passing Arguments**:

   - Goroutines can accept arguments just like regular functions.
   - However, it's important to pass arguments as values to ensure they are not shared between goroutines unless synchronized.

   Example:

   ```go
   package main

   import (
       "fmt"
       "time"
   )

   func greet(name string) {
       fmt.Println("Hello,", name)
   }

   func main() {
       go greet("Alice") // Pass argument to goroutine
       go greet("Bob")   // Pass argument to another goroutine
       time.Sleep(100 * time.Millisecond) // Wait for goroutines to finish
   }
   ```

5. **Synchronization**:

   - Goroutines can be synchronized using synchronization primitives like channels or `sync.WaitGroup` to coordinate their execution.
   - Without proper synchronization, concurrent goroutines may exhibit race conditions or other unexpected behavior.

   Example with `sync.WaitGroup`:

   ```go
   package main

   import (
       "fmt"
       "sync"
   )

   func printNumbers(wg *sync.WaitGroup) {
       defer wg.Done() // Mark this goroutine as done when finished
       for i := 0; i < 5; i++ {
           fmt.Println(i)
       }
   }

   func main() {
       var wg sync.WaitGroup
       wg.Add(1) // Increment WaitGroup counter
       go printNumbers(&wg) // Start goroutine
       wg.Wait() // Wait for all goroutines to finish
       fmt.Println("Main function")
   }
   ```

Goroutines are a powerful feature of Go, enabling efficient concurrent programming. They allow developers to write concurrent code in a straightforward manner, making it easier to utilize the full potential of modern multi-core processors. However, it's important to understand and manage goroutines properly to avoid common concurrency pitfalls such as race conditions and deadlocks.

<div id='lifecycle_states'></div>

## Lifecycle states

Goroutines in Go have three main lifecycle states:

1. **Runnable:** When a goroutine is created using the `go` keyword, it enters the runnable state. It's ready to be executed but may not be actively running at that moment due to scheduling by the Go runtime.

2. **Running:** Once a runnable goroutine gets scheduled by the Go runtime, it enters the running state. In this state, the goroutine is actively executing its code.

3. **Blocked:** A goroutine can enter the blocked state when it's waiting for some external event, such as I/O operations, channel operations, or synchronization with other goroutines. While in this state, the goroutine is not consuming CPU time and is effectively paused until the event it's waiting for occurs.

<div id='go_scheduler'></div>

## Go scheduler

Go has its own scheduler called the Go scheduler, which is responsible for managing goroutines. It's different from the scheduler used in Java, but it serves a similar purpose. The Go scheduler is designed to efficiently manage the execution of goroutines across multiple OS threads, utilizing techniques like preemptive scheduling, work-stealing, and blocking/unblocking mechanisms to maximize CPU utilization and concurrency. This scheduler is part of the Go runtime and works alongside other components to provide efficient and concurrent execution of Go programs.

Here's an overview of the process involved in creating and executing a goroutine in Go:

1. **Goroutine Creation:** Goroutines are lightweight threads managed by the Go runtime. They are created using the `go` keyword followed by a function call. For example:

   ```go
   go myFunction()
   ```

   When this line executes, a new goroutine is spawned to execute the `myFunction()` concurrently with other goroutines.

2. **Stack Allocation:** When a goroutine is created, the Go runtime allocates a small initial stack (usually a few kilobytes) for it. This stack space is used to store local variables, function parameters, and other execution-related data.

3. **Scheduling:** Once a goroutine is created, the Go scheduler decides when and on which OS thread to run it. The scheduler may run multiple goroutines simultaneously on different OS threads, or it may multiplex goroutines onto fewer OS threads, depending on the available CPU resources and workload.

4. **Execution:** When a goroutine gets scheduled to run, it starts executing its associated function from the beginning. It continues executing until it reaches a blocking operation, such as I/O operation, channel operation, or synchronization, or until it explicitly completes by returning from the function.

5. **Concurrency:** Goroutines run concurrently, meaning they can execute simultaneously with other goroutines. The Go scheduler handles the coordination and switching between goroutines to ensure fair execution and efficient CPU utilization.

6. **Completion:** When a goroutine completes its execution (either by reaching the end of its function or encountering a `return` statement), it exits, and its stack space is deallocated by the Go runtime.

<div id='Channels_and_communication_between_Goroutines'></div>

## Channels and communication between Goroutines

Channels in Go facilitate communication and synchronization between goroutines, providing a safe and efficient way to share data and coordinate concurrent execution. They are the primary means of communication in Go's concurrency model. Channels can be thought of as pipes through which data flows between goroutines.

Here's a detailed explanation of channels and communication between goroutines:

### Channels:

1. **Creation**:

   - Channels are created using the `make()` function, specifying the data type of the values that will be transmitted through the channel.
   - Channels can be unbuffered (default) or buffered.
   - Unbuffered channels block sender goroutines until a receiver is ready, ensuring synchronization.

   Example:

   ```go
   ch := make(chan int) // Create an unbuffered channel for transmitting integers
   ```

2. **Sending and Receiving**:

   - Data is sent to a channel using the `<-` operator with the channel variable on the left-hand side.
   - Data is received from a channel using the same `<-` operator, but with the channel variable on the right-hand side.
   - Both send and receive operations are blocking by default, ensuring synchronization between sender and receiver.

   Example:

   ```go
   ch <- value // Send value to the channel
   data := <-ch // Receive data from the channel
   ```

3. **Closing Channels**:

   - Channels can be closed using the `close()` function.
   - Closing a channel indicates that no more data will be sent through it.
   - It's important to always close channels after use to avoid deadlocks.

   Example:

   ```go
   close(ch) // Close the channel
   ```

4. **Buffered Channels**:

   - Buffered channels have a fixed capacity, allowing a certain number of values to be stored without a corresponding receiver.
   - Sending to a buffered channel blocks only when the buffer is full, and receiving blocks only when the buffer is empty.

   Example:

   ```go
   ch := make(chan int, 10) // Create a buffered channel with capacity 10
   ```

5. **Select Statement**:

   - The `select` statement allows goroutines to wait on multiple channel operations simultaneously.
   - It's useful for coordinating communication between multiple channels and goroutines.

   Example:

   ```go
   select {
   case data := <-ch1:
       // Handle data received from ch1
   case ch2 <- value:
       // Send value to ch2
   }
   ```

### Communication between Goroutines:

1. **Producer-Consumer Pattern**:

   - Channels are commonly used to implement the producer-consumer pattern, where one goroutine produces data and sends it to another goroutine for consumption.

   Example:

   ```go
   func producer(ch chan<- int) {
       for i := 0; i < 5; i++ {
           ch <- i // Send data to the channel
       }
       close(ch) // Close the channel when done
   }

   func consumer(ch <-chan int) {
       for num := range ch {
           fmt.Println("Received:", num) // Receive data from the channel
       }
   }

   func main() {
       ch := make(chan int)
       go producer(ch)
       consumer(ch)
   }
   ```

2. **Worker Pool Pattern**:

   - Channels can be used to implement a worker pool, where a fixed number of goroutines (workers) consume tasks from a channel.

   Example:

   ```go
   func worker(id int, tasks <-chan int, results chan<- int) {
       for task := range tasks {
           // Process task
           results <- task // Send result to the results channel
       }
   }

   func main() {
       numWorkers := 3
       tasks := make(chan int, 10)
       results := make(chan int, 10)

       // Create worker goroutines
       for i := 0; i < numWorkers; i++ {
           go worker(i, tasks, results)
       }

       // Produce tasks
       for i := 0; i < 10; i++ {
           tasks <- i // Send task to the tasks channel
       }
       close(tasks) // Close tasks channel to indicate no more tasks

       // Collect results
       for i := 0; i < 10; i++ {
           result := <-results // Receive result from the results channel
           fmt.Println("Result:", result)
       }
   }
   ```

Channels are a powerful synchronization mechanism in Go, enabling effective communication and coordination between goroutines. By using channels, developers can write concurrent programs that are safe, efficient, and easy to reason about. However, it's crucial to understand the principles of channel usage to avoid common pitfalls such as deadlocks and race conditions.

<div id='Buffered_channels'></div>

## Buffered channels

Buffered channels in Go provide a means of decoupling the sending and receiving operations, allowing a certain number of elements to be stored in the channel's internal buffer. This buffer can help improve concurrency by reducing the potential for goroutines to block while waiting for communication. Understanding buffered channels is essential for building efficient concurrent applications in Go.

Let me illustrate buffered channels with code examples:

### Creation of Buffered Channels:

To create a buffered channel, you specify the buffer size when using the `make()` function. For instance, `make(chan int, 10)` creates a buffered channel of integers with a capacity of 10.

```go
ch := make(chan int, 3) // Create a buffered channel with a buffer size of 3
```

### Sending and Receiving from Buffered Channels:

Sending data to a buffered channel only blocks when the buffer is full. Similarly, receiving from a buffered channel blocks only when the buffer is empty.

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2) // Buffered channel with a capacity of 2

	// Sending data to the buffered channel
	ch <- 1
	ch <- 2

	// Receiving data from the buffered channel
	fmt.Println(<-ch) // Output: 1
	fmt.Println(<-ch) // Output: 2
}
```

### Blocking Behavior:

When the buffer is full, sending to a buffered channel blocks until space is available. Conversely, when the buffer is empty, receiving from a buffered channel blocks until data is available.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2) // Buffered channel with a capacity of 2

	// Sending data to the buffered channel
	ch <- 1
	ch <- 2

	// Attempting to send more data, which blocks until space is available
	go func() {
		ch <- 3
		fmt.Println("Sent 3 to channel")
	}()

	// Receiving data from the buffered channel
	time.Sleep(time.Second) // Adding delay to illustrate blocking behavior
	fmt.Println(<-ch)       // Output: 1
	fmt.Println(<-ch)       // Output: 2
	fmt.Println(<-ch)       // Output: 3
}
```

### Closing Buffered Channels:

Just like unbuffered channels, buffered channels can also be closed using the `close()` function. However, closing a buffered channel should be done with caution, as closing it while there are remaining values in the buffer can lead to panic.

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2) // Buffered channel with a capacity of 2

	ch <- 1
	ch <- 2

	close(ch) // Closing the buffered channel

	// Attempting to send more data after closing the channel results in panic
	ch <- 3 // panic: send on closed channel
}
```

By employing buffered channels, you can optimize your concurrent Go programs by reducing potential blocking operations and improving overall efficiency. However, it's crucial to understand the buffering behavior and handle channel closing appropriately to prevent runtime errors.

<div id='Select_statement_for_managing_multiple_channels'></div>

## Select statement for managing multiple channels

The `select` statement in Go is a powerful tool for managing multiple channels concurrently. It allows a Go routine to wait on multiple channel operations simultaneously, enabling flexible and efficient communication between goroutines. Here's how `select` works with code examples:

### Basic Usage:

The `select` statement waits for one of its cases to become ready for communication. If multiple cases are ready, it chooses one at random.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Hello"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "World"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received message from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received message from ch2:", msg2)
	}
}
```

In this example, the main goroutine waits for messages from `ch1` and `ch2` using `select`. Whichever channel sends a message first will have its corresponding case executed.

### Default Case:

You can include a `default` case in a `select` statement to perform non-blocking operations when no other case is ready.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received")
	}
}
```

In this example, if no message is received from `ch` within the specified time, the default case will be executed.

### Multi-Channel Communication:

`select` can wait on multiple channels simultaneously. It will choose a case randomly among those that are ready.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Hello"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "World"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received message from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received message from ch2:", msg2)
	}
}
```

In this example, the main goroutine waits for messages from `ch1` and `ch2` simultaneously. Whichever channel sends a message first will have its corresponding case executed.

### Timeouts:

`select` can be used in conjunction with `time.After` to implement timeouts for channel operations.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received message:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout occurred")
	}
}
```

In this example, if no message is received from `ch` within 1 second, the `time.After` case will be executed, indicating a timeout.

The `select` statement in Go provides a concise and powerful mechanism for managing multiple channels concurrently, enabling efficient communication and synchronization between goroutines. It's a fundamental tool for building robust and scalable concurrent applications in Go.

<div id='Synchronization_using_WaitGroups'></div>

## Synchronization using WaitGroups

Synchronization is essential in concurrent programming to ensure that multiple goroutines coordinate their execution correctly. `sync.WaitGroup` in Go provides a simple and effective way to synchronize goroutines, allowing one goroutine to wait for a group of goroutines to finish their tasks before proceeding.

Here's an in-depth explanation of synchronization using `sync.WaitGroup` in Go:

### Purpose of WaitGroups:

`sync.WaitGroup` allows you to wait for a collection of goroutines to complete their tasks. It helps synchronize the main goroutine with other goroutines by waiting until all goroutines in the group have finished before proceeding.

### Key Methods:

1. **Add()**: Used to add the number of goroutines to the WaitGroup. This number represents the number of goroutines the WaitGroup should wait for.
2. **Done()**: Signals that a goroutine has finished its task. It decrements the counter added by `Add()`.
3. **Wait()**: Blocks until the counter becomes zero. It waits for all goroutines to call `Done()` and for the counter to reach zero.

### Example:

Let's illustrate the usage of `sync.WaitGroup` with a simple example where multiple goroutines print their IDs and signal their completion using `Done()`:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion when exiting the function
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate some work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each goroutine
		go worker(i, &wg)
	}

	wg.Wait() // Wait until all goroutines are done
	fmt.Println("All workers have completed")
}
```

In this example:

- The `worker` function simulates some work by sleeping for a second.
- Inside the main function, a `sync.WaitGroup` is created.
- For each worker goroutine, the main function calls `wg.Add(1)` to increment the WaitGroup counter before launching the goroutine.
- Inside each worker goroutine, `wg.Done()` is called to signal completion when the worker is done with its task.
- Finally, `wg.Wait()` is called in the main function to block until all worker goroutines have finished their tasks.

### Benefits of WaitGroups:

- **Simple Synchronization**: WaitGroups provide a straightforward way to synchronize goroutines without the need for additional synchronization primitives.
- **Flexible Usage**: WaitGroups are flexible and can be used to synchronize any number of goroutines.
- **No Race Conditions**: When used correctly, WaitGroups prevent race conditions and ensure safe coordination between goroutines.

### Use Cases:

- **Parallelism**: WaitGroups are commonly used to coordinate parallel tasks, where multiple goroutines perform tasks concurrently.
- **Fan-Out/Fan-In**: They are used in fan-out/fan-in patterns where multiple goroutines produce results (fan-out) and then multiple goroutines consume those results (fan-in).

### Conclusion:

`sync.WaitGroup` is a powerful synchronization primitive in Go that simplifies the coordination of concurrent tasks. By utilizing `Add()`, `Done()`, and `Wait()` methods, you can ensure that the main goroutine waits for other goroutines to finish their tasks before proceeding. Understanding and mastering WaitGroups is essential for writing efficient and safe concurrent Go programs.

<div id='Mutexes_and_RWMutexes_for_safe_concurrent_access'></div>

## Mutexes and RWMutexes for safe concurrent access

Mutexes and RWMutexes are synchronization primitives in Go used to control access to shared resources in concurrent programs. They help ensure safe and consistent access to data by allowing only one goroutine to access the resource at a time. While Mutexes are used for exclusive access (read and write), RWMutexes provide a more flexible solution, allowing multiple readers or a single writer at a time.

### Mutex (Mutual Exclusion):

Mutex, short for mutual exclusion, is a synchronization primitive used to protect shared resources from concurrent access by multiple goroutines. It ensures that only one goroutine can access the shared resource at any given time, thereby preventing data races and maintaining data consistency.

#### Key Methods:

- **Lock()**: Acquires the lock, blocking until it's available. If the lock is already held by another goroutine, the calling goroutine will wait until it's released.
- **Unlock()**: Releases the lock. It should always be called in a deferred manner to ensure that the lock is released even in case of errors.
- **RLock() and RUnlock()**: Mutexes don't have built-in support for read locks. Use `RLock()` to acquire a read lock, allowing multiple goroutines to read concurrently. `RUnlock()` is used to release the read lock.

#### Example:

```go
package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mu sync.Mutex

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
```

In this example, `increment()` function is protected by a Mutex. Only one goroutine can execute `increment()` at a time, ensuring that `counter` is updated atomically.

### RWMutex (Reader-Writer Mutex):

RWMutex, short for reader-writer mutex, is a synchronization primitive that allows multiple readers or a single writer at any given time. It provides a more flexible locking mechanism compared to Mutex, allowing concurrent reads but exclusive writes.

#### Key Methods:

- **RLock()**: Acquires a read lock, allowing multiple goroutines to read concurrently.
- **RUnlock()**: Releases the read lock.
- **Lock()**: Acquires an exclusive write lock, blocking until it's available. It prevents both readers and writers until the lock is acquired.
- **Unlock()**: Releases the write lock.

#### Example:

```go
package main

import (
	"fmt"
	"sync"
)

var (
	counter = 0
	mu      sync.RWMutex
)

func read() {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("Counter:", counter)
}

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			read()
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		increment()
	}()
	wg.Wait()
}
```

In this example, `read()` function is protected by a RWMutex. Multiple goroutines can read `counter` concurrently using `RLock()` while `increment()` function acquires an exclusive write lock using `Lock()` before updating `counter`.

### When to Use Mutexes vs RWMutexes:

- Use Mutex when you need exclusive access to a resource. For example, when modifying a shared variable.
- Use RWMutex when you have multiple readers and occasional writers. For example, when reading data from a shared resource that is infrequently updated.

Mutexes and RWMutexes are synchronization primitives in Go that help manage concurrent access to shared resources. By understanding their differences and appropriate use cases, you can write safe and efficient concurrent Go programs, preventing data races and ensuring data consistency.
