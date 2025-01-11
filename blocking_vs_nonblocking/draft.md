# Introduction to Blocking Vs Non-blocking Programming

Before reading this article, it is important to inform you that this article is directly inspired by
and prepared from this documentation :

[Node.js Asynchronous flow control](https://nodejs.org/en/learn/asynchronous-work/asynchronous-flow-control)
[Node.js Overview of Blocking vs Non-Blocking](https://nodejs.org/en/learn/asynchronous-work/overview-of-blocking-vs-non-blocking)
[Node.js JavaScript Asynchronous Programming and Callbacks](https://nodejs.org/en/learn/asynchronous-work/javascript-asynchronous-programming-and-callbacks)
[Spring Web on Servlet Stack](https://docs.spring.io/spring-framework/reference/web.html)
[Spring Web on Reactive Stack](https://docs.spring.io/spring-framework/reference/web-reactive.html)

So, it's highly recommended to read them or at least just take a look.

## Concept

### Computers are asynchronous by design.

In my opinion, there are two main things that distinguish computers from humans, the computer can perform complex calculations, and therefore complex operations, tremendously fast (often in just a few milliseconds). And that the computer can multitask unlike you, a human, who cannot do two or more different things at the same time. This introduction takes us to the next question.

Have you ever thought about how can you play video games, listen to music, and talk to your friends in a voice channel
all at the same time? How can your computer handle all these operations at the same time with only just one processor?

Well, let me tell you the truth, there's actually no multiple operations happening at the same time, the processor runs every program for a specific time and then it stops its execution to let another program continue their execution. This thing runs in a cycle so fast that it's impossible for you to notice the delay. We think our computers run many programs simultaneously, but this is an illusion (except on multiprocessor machines).

But most programming languages are synchronous, meaning that they execute code sequentially one line after the other, so, if you are performing a process that will take a lot of time to execute, the program will be blocked and wait for this operation to complete before processing any other operations.

Ok, so what?

Actually, this is a big problem for scalability, very high throughput, head-of-line blocking, and more. Put the fancy words aside and let's discuss a simple example.

You are a mobile app developer who is developing a simple application that contains only one button and one text field. Every time you click the button, a spinning loading dialog appears while you are fetching a random meme from some memes API on the internet, and once you fetch the meme and that it's ready, the spinning loading dialog disappears and the meme is shown in the text field.

To sum up, here's the flow :
1- User clicks the button.
2- Loading dialog shows up.
3- Fetch a meme from an API.
4- When the meme is ready, dismiss the loading dialog and show the meme in the text field.

This flow has a problem in it, you need to show the spinning dialog and fetch the meme at the same time, and our progamming language runs sequentially so, it can not keep showing the spinning dialog while performing the fetch operation. What will happen is that the spinning dialog would freeze at the time that the program is trying to fetch the meme as it can not keep showing the spinning dialog while performing another operation.

Intrusive thought : Just remove the spinning dialog.

If you had the same thought above, not just the UI/UX team would be upset from you, but it's actually more complicated than what you think, let's see how.

In default conditions, the programming language performs all the code on the `main` thread. This thread should be always aware of the user's actions, like swipes, clicks and so on, let's say that there's a cancel button that cancels the operation once clicked on. If the `main` thread is busy performing the fetching operation itself and the UI is frozen, how would it be aware of the user clicks? Remember that it performs operations sequentially one after the other. So, the `main` thread should only process one task, that is the user's actions, and any other operation, like fetching data from an API or any other IO task, should be handled on a different thread and notify the `main` thread once it's done.

This means :
- The `main` thread has only one task, that is handling UI and users actions when triggered.
- Another thread has only one task, that is fetching data from an external API.

This is the non-blocking logic, where we did not block the `main` thread to perform some task that would take a lot of time like an IO operation, but instead, we delegated this process to another thread that will do it for us while the `main` thread is still responsive and once the other thread finishes the task, it notifies the `main` thread via callbacks.

## Code

Now let's see some code examples using two different programming languages, one multi-threaded language like Java, and the other is a single-threaded language like JavaScript.

### Blocking Example:

In Java, it will be something like that :

```java
class Demo1 {
    public static void main(String[] args) {
        long startTime = System.currentTimeMillis();

        System.out.println("Starting blocking operation...");
        String result = performBlockingOperation();
        System.out.println("Blocking operation result: " + result);

        long endTime = System.currentTimeMillis();
        System.out.println("Time taken: " + (endTime - startTime) + " ms");
    }

    private static String performBlockingOperation() {
        try {
            // Simulate time-consuming I/O operation
            Thread.sleep(3000); // 3 seconds
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        return "Operation completed";
    }
}
```

And the output will look like this :

```bash
❯ javac Demo1.java

❯ java Demo1
Starting blocking operation...
Blocking operation result: Operation completed
Time taken: 3004 ms
```

Note how the code has been executed sequentially one line after the other. Now let's see the same example in JavaScript :

```js
const startTime = new Date();

console.log("Starting blocking operation...");
const reslut = performBlockingOperation();
console.log(`blocking operation result: ${reslut}`);

const endTime = new Date();
console.log(`Time taken: ${endTime - startTime} ms`);

function performBlockingOperation() {
    const start = Date.now();
    while (Date.now() - start < 3000) {
        // Simulate time-consuming I/O operation
    }

    return "Operation completed";
}
```

And the output of this code :
```bash
❯ node Demo1.js
Starting blocking operation...
blocking operation result: Operation completed
Time taken: 3004 ms
```

### Non-Blocking Example:

Let's do the magic of asynchronous programming. We will start with the Java one :

```java
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.CountDownLatch;

class Demo2 {
    public static void main(String[] args) {
        long startTime = System.currentTimeMillis();
        CountDownLatch latch = new CountDownLatch(1);

        // Start non-blocking operation
        System.out.println("We are now on " + Thread.currentThread().getName());
        System.out.println("Starting non-blocking operation...");
        CompletableFuture<String> futureResult = performNonBlockingOperation();

        // this is to handle when the non-blocking operation finishes
        futureResult.whenComplete((res, t) -> {
            System.out.println(res);
        });

        long endTime = System.currentTimeMillis();
        System.out.println("Time taken: " + (endTime - startTime) + " ms");

        // this tells the program not to terminate before the futureResult resolves
        try {
			latch.await();
		} catch (InterruptedException e) {
			e.printStackTrace();
		} // Wait until the countdown reaches 0
    }

    private static CompletableFuture<String> performNonBlockingOperation() {
        return CompletableFuture.supplyAsync(() -> {
            System.out.println("performing non-blocking operation on " + Thread.currentThread().getName());
            try {
                // Simulate time-consuming I/O operation
                Thread.sleep(3000); // 3 seconds
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

            return "non-blocking operation completed";
        });
    }
}
```

Note: Please ignore the `CountDownLatch` part as it's there only to inform java not to exit the program before the non-blocking part finishes.

Let's see the output of this code :

```bash
❯ javac Demo2.java

❯ java Demo2
We are now on main
Starting non-blocking operation...
performing non-blocking operation on ForkJoinPool.commonPool-worker-1
Time taken: 13 ms
non-blocking operation completed
```

Notice how the `Time taken: 13 ms` part was printed before the non-blocking operation finishes, the `main` thread did not get blocked waiting for the non-blocking operation to be done and instead proceeded with the rest of the program.

Now let's see the exact same example in JavaScript :

```js
const startTime = new Date();

// Start non-blocking operation
console.log(`We are now on ${process.pid}`);
console.log("Starting non-blocking operation...");
const futureResult = performNonblockingOperation();

// this is to handle when the non-blocking operation finishes
futureResult.then((res) => {
    console.log(res);
});

const endTime = new Date();
console.log(`Time taken: ${endTime - startTime} ms`);

async function performNonblockingOperation() {
    return new Promise(async (resolve, reject) => {
        console.log(`performing non-blocking operation on ${process.pid}`);
        try {
            // Simulate time-consuming I/O operation
            await sleep(3000);
        } catch (error) {
            reject(error);
        }

        resolve('non-blocking operation completed');

    });
}

async function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
```

Let's see the output of this code :

```bash
❯ node Demo2.js
We are now on 98950
Starting non-blocking operation...
performing non-blocking operation on 98950
Time taken: 4 ms
non-blocking operation completed
```

Notice how the `Time taken: 4 ms` part was also printed before the non-blocking operation finishes.

### Is JavaScript faster than Java?

Did you also notice the execution time difference in both languages? Java took 13 ms and JavaScript took only 4 ms. Does this mean JavaScript is faster than Java? Let's see.

This timing does not measure the actual execution time of the non-blocking operation. Instead, it measures how long the program takes to schedule the operation and move on to the next instruction (logging the time taken). The non-blocking operation itself (which takes 3 seconds in both cases) completes after this time is logged.

Why the Time Differences?
Several factors contribute to the difference in logged time between Java and Node.js:

1- Event Loop vs. Thread Pool:

• Node.js uses an event loop that handles asynchronous I/O operations very efficiently. It quickly schedules the task (in this case, a 3-second sleep) and moves on.
• Java uses a thread pool for asynchronous tasks with CompletableFuture.supplyAsync(). The ForkJoinPool.commonPool() creates or reuses threads, which might take a bit longer to schedule compared to Node.js's event loop.

2- JVM Initialization Overhead:

• Java programs can have higher initial overhead due to the JVM (Java Virtual Machine) startup. This can cause slightly longer times for the initial scheduling of asynchronous tasks.
• Node.js has a lighter runtime compared to the JVM, so it can appear to log the time taken more quickly.

3- CPU Scheduling and OS-level Factors:

• The actual time reported (e.g., 13 ms in Java vs. 4 ms in Node.js) is influenced by many OS-level factors, including how the CPU schedules threads and manages processes. This doesn’t mean Node.js is inherently faster in execution but reflects how quickly the runtime environment schedules the asynchronous task.

So, Is Node.js Faster than Java?
No, this difference in the logged time does not mean Node.js is faster than Java in terms of overall performance.
