#  Comparing the performance of different methods/packages and HTTP servers in both Go and Rust.
 
We have created the following project in fullfillment of requirement of the course CS-F301 : Principles of Programming Language for IC-  Prof. Kunal Korgaonkar
 
 
## Group Members
Prasoon Bagla -2020B3A71159G
Vidhi Kabra -2020B3A70569G
Varun Gopal -2020B3A71785G
Yash Shardendu Agrawal -2020B1A71589G
 
 
 
 
 
## 1). Problem Statement
## Objective
This project aims to compare the performance features of Go and Rust in two specific areas: computational properties using the Fannkuch-redux algorithm and basic HTTP server functionality.
 
## POPL Angle
The Principles of Programming Languages (POPL) perspective is central to this analysis. By evaluating how Go and Rust perform in identical tasks, we gain insights into how language design and features impact performance. This comparative study focuses on syntactical and structural aspects, memory management, concurrency models, and type systems inherent to each language.
 
## Previous Solutions and Differentiation
Comparative studies between programming languages are common, but this project takes a practical, application-based approach. By using specific algorithms and server implementations, it provides direct insights into real-world performance, rather than relying on theoretical benchmarks.
 
 
## 2). Software Architecture 
 
## Computational Components: 
Standalone units for Go and Rust implementing the Fannkuch-redux algorithm.
 
## Web Server Components: 
Basic HTTP servers in both languages.
Data Logging/Reporting: Performance metrics are logged using Excel file generation and manipulation.
 
## Client-Server Architecture: 
The HTTP servers indicate a client-server setup.
## Testing Components: 
The placement of testing components (local or remote) is inferred but not explicitly stated in the code.
No Database Involvement: Data storage is file-based, focusing on Excel files rather than a database.
 
## Reused and Developed Components
Parts of the architecture, like HTTP server setup and Excel file manipulation, rely on existing libraries and frameworks. The core computational and performance measurement components have been developed to suit the specific needs of this project.
 
 
## 3). POPL Aspects in Implementation

## Go Implementation
Concurrency Model: The use of goroutines in fannkuch-redux.go demonstrates Go's lightweight thread management, a key aspect of its concurrency model.
Memory Management: Go's garbage collection is evident in how memory allocation is handled in fib.go .
## Rust Implementation
Ownership and Borrowing: Rust's unique approach to memory safety through ownership is showcased in fannkuch-redux.rs.
Concurrency with Safety: The use of threads in Rust, ensuring memory safety, is seen in fib.rs .
These aspects are crucial for understanding the performance and safety guarantees each language offers.

## Lines of Code:

## Go Files:
fannkuch-redux.go - Concurrency Model:

Line 21-23: go fannkuch(idxMax, ch) inside a loop. This is a clear example of Go's concurrency model where goroutines are used for parallel computation.
Line 29-31: Creation of slices p := make([]int, n). In Go, slices are a key data structure, showcasing its approach to memory management and data handling.
fib.go - Memory Management:

Line 18-20: http.HandleFunc("/", handler) and go http.ListenAndServe(":8080", nil). Here, Go handles HTTP requests concurrently, showcasing its efficient use of resources and memory management.
## Rust Files:
fannkuch-redux.rs - Ownership and Borrowing:

Line 10-12: fn rotate(x: &mut [i32]). Rust’s borrowing rules are applied here, where x is a mutable reference, ensuring that the function can modify the data it points to safely.
Line 16-18: fn next_permutation(perm: &mut [i32], count: &mut [i32]). This function signature again emphasizes Rust's ownership model, where mutable references are explicitly stated.
fib.rs - Concurrency with Safety:

Commented Line 1-3: Use of #[tokio::main] for asynchronous programming. This attribute macro sets up an asynchronous runtime, which is central to Rust's safe concurrency model.
Commented Line 5-7: warp::path!("hello" / String).map(|name| { format!("Hello, {}!", name) });. This line, though commented, hints at the safe handling of concurrent web requests using Warp, a Rust web server framework.
 
 
## 4). Results
 
## HTTP Server analysis
We have analysed the compared and analysed the response times of 1000 HTTP requests on both go and rust. The median response time for the Go server is 1.232 units, while for the Rust server it’s 2.101 units. This suggests that typically, the Go server has a faster response time than the Rust server. The standard deviation for the Go server is 0.543 units, while for the Rust server it’s 1.856 units. This indicates that the response times for the Rust server are more spread out and vary more widely, while the response times for the Go server are more consistent. The Go server’s response time peaks around 8 units, while the Rust server’s response time goes up to 20 units. This could suggest that the Go server is able to handle requests more quickly. In conclusion, based on these statistics, the Go server not only typically responds faster (lower median), but its response times are also more consistent (lower standard deviation). Many factors can influence these metrics and these are the factors that favour Go over Rust when it comes to HTTP servers:
1)Go (Goroutines): Go is designed to be concurrent with lightweight goroutines and provides built-in support for concurrency. The http.ListenAndServe function uses goroutines to handle incoming requests concurrently. This can lead to efficient utilization of resources and faster response times.  In the Rust code, tokio::spawn is used for asynchronous execution, but the server itself may not be handling requests concurrently. Depending on the specific warp configuration and how it handles incoming requests, it might not be as concurrent as the Go server.
2) Underlying Runtime and Execution Model:Go has a runtime that includes garbage collection and is optimized for concurrent execution. The Go scheduler is designed to efficiently manage goroutines.
Rust has a different execution model. Tokio provides asynchronous runtime.

 ## HTTP Server analysis Go :
 ![Alt text](/results/go%20graph.png)
 ## HTTP Server analysis Rust :
![Alt text](/results/rust%20graph.png)
## Pfannkuchen Algorithm analysis:
 
## Memory usage:
it appears that the Pfannkuchen function uses more memory on average in Rust than in Go. The mean peak memory used for Rust is 911 KB, while for Go it’s 106 KB. This suggests that, on average, Rust uses more memory than Go when executing the Pfannkuchen function.
 
However, the standard deviation, which measures the amount of variation or dispersion from the average, is higher for Go (151.296) than for Rust (37.176). This indicates that the memory usage for Go varies more widely than those for Rust.
 
The memory usage difference can be attributed to the design and optimization of the two languages:
 
Rust uses a strict and safe memory system that eliminates the need for a garbage collector, resulting in faster speed and lower memory usage. However, Rust’s static type system and ownership model, which allow the compiler to generate more efficient code, can lead to higher memory usage in some cases.
 
Go, on the other hand, uses garbage collection, which automatically frees the memory that is no longer needed by the program. Additionally, Go’s focus on simplicity means that it generally requires less code than Rust and, thus, less memory. However, Go’s garbage collector periodically works in the background to free up data once you hit some pre-specified value, which can add to system overhead and cause variations in memory usage.

## Peak Memory Usage Go:
![Alt text](/results/peak%20memory%20go.png)

## Peak Memory Usage Rust:
![Alt text](/results/peak%20memory%20rust.png)

## Time Elapsed
 
It appears that the Pfannkuchen function runs faster on average in Rust than in Go. The mean time elapsed for Rust is 652 microseconds, while for Go it’s 3295.834 microseconds. This suggests that, on average, Rust executes the Pfannkuchen function more quickly than Go.
 
However, the standard deviation, which measures the amount of variation or dispersion from the average, is higher for Rust (146467663.2) than for Go (19223660.31). This indicates that the execution times for Rust vary more widely than those for Go.
 
The performance difference can be attributed to the design and optimization of the two languages:
 
Rust is designed for systems programming and high-performance computing, with a focus on zero-cost abstractions, minimal runtime, and improved memory safety. It provides fine-grained control over how threads behave and how resources are shared between threads. This can lead to more efficient execution of certain tasks, such as the Pfannkuchen function.
 
Go, on the other hand, was designed with simplicity and readability in mind. It uses garbage collection and has features like Goroutines that make it easy for developers to build applications that take full advantage of concurrency. However, these features might introduce some overhead, leading to slightly slower execution times compared to Rust.

## Time Elapsed Go: 
![Alt text](/results/time%20elapsed%20go%20.png)
## Time Elapsed Rust: 
![Alt text](/results/time%20elapsed%20rust.png)
## 5). Potential for Future Work
Given more time, the project could explore:
 
Advanced Concurrency Patterns: Investigating complex concurrency models in Go and Rust.
Memory Optimization Techniques: Deeper analysis of memory usage and optimization.
Expanded Benchmarking: Including more diverse computational tasks and server functionalities.
Language Interoperability: Exploring how Go and Rust can interoperate in a larger system.
Additional POPL aspects like error handling paradigms, type inference, and metaprogramming capabilities could also be examined.
 
## Running Tests
 
To run tests, run the following commands
 
```bash
  cargo build 
  cargo run --build <file name> <input>
```
 
```bash
  go build
  go run <file name> <input>
```
 
 
