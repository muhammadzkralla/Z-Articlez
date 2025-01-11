# Blocking vs Non-blocking code

Did you ask yourself before how can the computer handle multiple tasks at the same time? You can play games, listen to music, and talk with your friends on a voice chat and more all at the same time.

Similarly, an application can render UI, fetch data from some API, and perform a background task and more all at the same time.

The first scenario is handled for you by the OS, but the second one needs some work done by you, the application programmer.

I will not be explaining the OS part in this article, it's more focused on the application part.

## Concept

### What is a Thread? The Backbone of Modern Computing

A thread is the smallest unit of execution within a process. It represents a sequence of instructions that the CPU can execute independently. Threads enable parallelism and concurrency, allowing programs to perform multiple tasks simultaneously or appear to do so.

1- Thread vs. Process:
A process is an independent program running in its memory space, while a thread is a lightweight sub-unit of a process. A thread has its own stack for function calls, local variables, and a program counter to track instruction execution.

Multiple threads in a process share the same memory space, global variables, and resources, making communication and data sharing easier.

2- Lifecycle of a Thread: Threads typically go through the following states:
• New: Created but not started.
• Runnable: Ready to run, waiting for CPU time.
• Running: Actively executing.
• Blocked/Waiting: Paused while waiting for resources or conditions.
• Terminated: Completed its task.

### Blocking vs. Non-Blocking: The Basics

• Blocking: In blocking code, the execution of a program halts at a certain point until an operation completes (e.g., waiting for a file to load, a network response, or a database query). This can lead to inefficient resource utilization, especially in high-concurrency scenarios.

• Non-Blocking: Non-blocking code, on the other hand, allows the program to continue executing other tasks while waiting for the operation to complete. This improves efficiency and responsiveness, particularly in I/O-bound or high-concurrency applications.
