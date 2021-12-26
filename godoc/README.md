# Golab

Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

## pros and cons
### pros
#### fast 
Go is compiled to machine code, 
it will naturally outperform languages that are interpreted or have virtual runtimes

#### easy
Go’s syntax is small compared to other languages, and it’s easy to learn
#### strongly
Go is a strongly, statically typed language. There are primitive types like int, byte, and string. There are also structs. Like any strongly typed language, the type system allows the compiler helps catch entire classes of bugs. Go also has built-in types for lists and maps, and they are easy to use.
Easier Concurrency Model

#### Easier Concurrency Model
While concurrent programming is never easy, Go makes it easier than in other languages. It is almost trivial to create a lightweight thread, called a “goroutine”, and communicate with it via a “channel.” More complex patterns are possible.



#### Garbage Collection
Memory management in Go was intentionally made easier than in C and C++. Dynamically allocated objects are garbage collected. Go makes using pointers much safer because it doesn’t allow pointer arithmetic. It also gives you the option of using value types.

### cons
#### No Generics
First, the elephant in the room. Go doesn’t have generics. This is a big hurdle to get over when coming from languages like Java. It means a decreased level of reuse in your code. While Go has first-class functions, if you write functions like “map”, “reduce”, or “filter” that operates on a collection of one type, you can’t reuse those same functions for a collection of a different type. There are ways to deal with this, but they all ultimately involve writing more code. This hurts productivity and maintainability.


#### Poor Library Support
Library support for Go is spotty. Our API integrates with Contentful, which does not have an officially supported Go SDK. This meant we had to write (and maintain!) a TON of code to request and parse data from Contentful. We also had to rely on a third-party Elasticsearch library. The Go SDKs that are provided by vendors don’t get as much love as their Java, Ruby, or JavaScript counterparts.

#### Difficult Community
The Go community can be non-receptive to suggestions. Consider this issue in the GitHub repository for golint: https://github.com/golang/lint/issues/65. Someone requested the ability for golint to fail a build when there are warnings found (which is something we’re doing on our project!). The maintainer immediately dismissed the idea. Enough people commented on the issue, and the maintainer eventually added the requested feature, over a year later.

The Go community also appears to have an aversion to web frameworks. While Go’s HTTP library covers a lot, it doesn’t provide support for path parameters, input sanitization and validation, or many cross-cutting concerns often found in a web application. Ruby developers have Rails, Java developers have Spring MVC, and Python developers have Django. But many Go developers choose to avoid frameworks. Yet there are still frameworks out there. A lot of them! But it’s nearly impossible to choose one that won’t be abandoned after you’ve started a project with it.

### Fractured Dependency Management
For a long time Go did not have a stable, official package manager. After years of begging from the community, the Go project has recently released godep. Before that, many tools filled the gap. We use the very capable govendor in our project. But this means the community is fractured, and it can be extremely confusing for developers new to Go. Also, virtually all package management is backed by Git repositories, the history of which could change at any time. Compare this to something like Maven Central, which will never delete or change a library that your project depends on.


## Syntax

#### Declaring a Variable

**1. Using var keyword**

`var variable_name type = expression `

```go
var a int // a is 0
var b int = 100
var c = 100
```





**2. Using short variable declaration**

The **local variables**  which are declared and initialize **in the functions** are declared by using short variable declaration.

`variable_name:= expression` 

```go
e := 100 // in the function
```



### 3-const & iota

once the value of constant is defined it cannot be modified further.

Constants are declared like variables but in using a ***const*** keyword as a prefix to declare constant with a specific type. It cannot be declared using **:=** syntax. 

```go
    const const_string = "const_string"
```

iota

```
const={
	Argentina = iota
  Brazil
  Chile
}
```



### 5-init
### 6-pointer
### 7-defer
### 8-slice

In Go language slice is more powerful, flexible, convenient than an array and is a lightweight data structure

### 9-map
###  10-OOP
### 11-reflect
### 12-goroutine
###  13-channel
### 14-golang-IM-System





## Golang is perfectly suited for

### Real-time development

A technology to create programs that users sense as immediate or current, providing solutions for e-commerce, chatting, online gaming and more.


### Networking development

A study of how computers can be linked to share data. Today it involves web technologies, wireless trends and social media development. Golang features *goroutines* and *channels* were designed to meet modern tech demand. Go *attributes*, such as speed of development and deployment, also work perfectly for this task.


### Cloud infrastructure

Golang provides on-demand access to shared computer processing, allowing you to process your data on a privately owned cloud or a third-party server. It also worth noting that the well-known Google Cloud and Docker platform were developed using Golang.


### Micro services

A service-oriented architecture that structures an application, decomposes it to smaller services and makes it easier to understand, develop and test. Go-kit provides specialized support, such as infrastructure integration and system observability perfectly suited to writing microservices.


### Goroutine 


#### Goroutine vs Thread
Goroutine: A Goroutine is a function or method which executes independently and simultaneously in connection with any other Goroutines present in your program. Or in other words, every concurrently executing activity in Go language is known as a Goroutines.

Thread: A process is a part of an operating system which is responsible for executing an application. Every program that executes on your system is a process and to run the code inside the application a process uses a term known as a thread. A thread is a lightweight process, or in other words, a thread is a unit which executes the code under the program. So every program has logic and a thread is responsible for executing this logic.



| Goroutine                                                    | Thread                                                       |
| :----------------------------------------------------------- | :----------------------------------------------------------- |
| Goroutines are managed by the go runtime.                    | Operating system threads are managed by kernal.              |
| Goroutine are not hardware dependent.                        | Threads are hardware dependent.                              |
| Goroutines have easy communication medium known as channel.  | Thread does not have easy communication medium.              |
| Due to the presence of channel one goroutine can communicate with other goroutine with low latency. | Due to lack of easy communication medium inter-threads communicate takes place with high latency. |
| Goroutine does not have ID because go does not have Thread Local Storage. | Threads have their own unique ID because they have Thread Local Storage. |
| Goroutines are cheaper than threads.                         | The cost of threads are higher than goroutine.               |
| They are cooperatively scheduled.                            | They are preemptively scheduled.                             |
| They have fasted startup time than threads.                  | They have slow startup time than goroutines.                 |
| Goroutine has growable segmented stacks.                     | Threads does not have growable segmented stacks.             |
