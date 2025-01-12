# Introduction to Raw Socket Programming and Implementing the Simplest Chat Service.

## The Problem
You want to build a chat service. A user should be able to send and receive messages from and to other users. You will face a problem implementing this service in a traditional RESTful API. To know what this problem is, we need to know how RESTful APIs work in the first place.

REST is not a protocol but rather a design philosophy that builds upon the principles of HTTP. In REST, clients should send a request to the server in order to receive a response. The server CANNOT send a response to the client without a request sent from them in the first place.

## The Trivial Solution
Knowing the information above, how would you handle the scenario where a user sends a message to another user? Let's take it step by step. The sender will send a POST request to the server with the message details. Assume that the server received the request successfully and responded back with a 200 status code. Now, the server needs to send this message to the recipient without them making a request in the first place. As explained earlier, this is not feasible with pure REST.

## The Solution and Definition of Sockets
For every problem, there's a solution. Sockets come to the rescue in these scenarios.

Socket programming is a very important concept in computer science. It provides a full-duplex connection between the client and the server and often uses event-based communication with third-party libraries like Socket.io to handle data transmission in production code.

Understanding how sockets operate under the hood is essential, as it's a crucial step in initiating an HTTP request aside direct use in a lot of other important use cases like chat services, games, and any service that needs simultaneous communication.

Unlike REST, where the connection is one-time (each request may receive one response, and the connection is terminated after), in sockets, the connection is maintained over long periods of time, and both ends can receive updates from the other side immediately.

It's important to note that I'm not comparing REST and sockets. They are two whole different things that cannot be compared. REST is an architectural style based on HTTP, while socket is a low-level communication mechanism. Comparing them both is like saying which is faster, the sea or the tree?

A server socket needs an IP address and a port to start listening to (an address such as 127.0.0.1:420), now clients can open a socket connection with this server socket and leverage the full-duplex connection between them. In this article, we will implement the simplest chat service with just raw sockets (no third-party libraries like socket.io) and see how they communicate and transmit data between each other.

## Code

I will implement this service using two programming languages, Java and JavaScript, but I will start with JavaScript first as it's easier and event-based (more close to socket.io).

Let's start with the server code, I used the net node module and the readline/promises node module for asynchronous reads from the console in this example :

```JavaScript
const net = require("net");
const readline = require("readline");

// readline reads messages from the console without blocking the program
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
});

// Initiate the server and accept socket connections
const server = net.createServer((socket) => {
    console.log("Client connected.");

    // This callback is called whenever new data are sent to the server
    socket.on("data", (data) => {
        const message = data.toString();
        console.log(`Client: ${message}`);
    });

    // This callback is called whenever you write anything in the console
    rl.on("line", (line) => {
        if (line === "quit") {
            console.log("Shutting down server...");
            socket.destroy();
        } else {
            socket.write(line);
        }
    });
});

// Start listening to port 420 on the localhost
server.listen(420, "localhost", () => {
    console.log("Server is listening on port 420");
});
```

And the client code :

```JavaScript
const net = require("net");
const readline = require("readline");

// readline reads messages from the console without blocking the program
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
});

// Initiate the socket connection with the server on localhost port 420
const socket = new net.Socket();
socket.connect(420, "localhost", () => {
    console.log("Client started.");
});

// This callback is called whenever new data are sent to the client
socket.on("data", (data) => {
    const message = data.toString();
    console.log(`Server: ${message}`);
});

// This callback is called whenever you write anything in the console
rl.on("line", (line) => {
    if (line === "quit") {
        console.log("Closing client...");
        socket.destroy();
    } else {
        socket.write(line);
    }
});
```

Let's run both programs and see the results :

<INSERT IMAGE>

Now, let's see the Java code, we will also start with the server :

```java
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.ServerSocket;
import java.net.Socket;

class Server1 {

    // The constructor will initiate the server
    public Server1() {
        try {

            // Create a ServerSocket object on port 420 and accept the first client
            ServerSocket serverSocket = new ServerSocket(420);
            Socket socket = serverSocket.accept();
            System.out.println("Client connected");

            // The ServerSocket object has an input stream to receive messages and an output stream to send messages
            PrintWriter socketOutputStream = new PrintWriter(socket.getOutputStream(), true);
            BufferedReader socketInputStream = new BufferedReader(new InputStreamReader(socket.getInputStream()));

            while (socket.isConnected()) {
                // Take input from the console to send over the output stream
                BufferedReader input = new BufferedReader(new InputStreamReader(System.in));

                // Read received messages from the input stream
                System.out.println(socketInputStream.readLine());

                // Read input from the console
                String message = input.readLine();

                if (message.equals("quit")) {
                    serverSocket.close();
                    socket.close();
                    break;
                }

                // Write the message from the console to the output stream
                socketOutputStream.println(message);
            }
        } catch (Exception e) {
            System.out.println("Could not init server socket!" + e.getMessage());
        }
    }

    public static void main(String[] args) {
        // Create a new server instance
        new Server1();
    }
}
```

And now let's create the client code :

```java
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;

class Client1 {

    // The constructor will initiate the server
    public Client1() {
        try {

            // Initiate the socket connection with the server on localhost port 420
            Socket socket = new Socket("localhost", 420);
            System.out.println("Client started.");

            // The socket object has an input stream to receive messages and an output stream to send messages
            PrintWriter socketOutputStream = new PrintWriter(socket.getOutputStream(), true);
            BufferedReader socketInputStream = new BufferedReader(new InputStreamReader(socket.getInputStream()));

            while (socket.isConnected()) {
                // Take input from the console to send over the output stream
                BufferedReader input = new BufferedReader(new InputStreamReader(System.in));

                // Read received messages from the input stream
                String message = input.readLine();

                if (message.equals("quit")) {
                    socket.close();
                    break;
                }

                // Write the message from the console to the output stream
                socketOutputStream.println(message);

                // Read received messages from the input stream
                System.out.println(socketInputStream.readLine());
            }

        } catch (Exception e) {
            System.out.println("Could not init client socket!" + e.getMessage());
        }
    }

    public static void main(String[] args) {
        // Create a new client instance
        new Client1();
    }
}
```

This code will work fine, but it has a problem, the `readLine()` function is blocking so, the program will not process any incoming messages before it's done with the `readLine()` function. To solve this issue, we will run each process on a different thread.

NOTE: If you don't understand the difference between blocking and non-blocking code, please refer to this article.

This is the improved Java code that will run like the JavaScript example :

```Java
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.ServerSocket;
import java.net.Socket;

class Server2 {
    private ServerSocket serverSocket;
    private Socket socket;

    public Server2() {
        try {
            serverSocket = new ServerSocket(420);
            socket = serverSocket.accept();

            System.out.println("Client connected");

            // Handle client messages on a thread
            Thread clientHandler = new Thread(() -> handleClient());
            clientHandler.start();

            // Handle the user input from the console on another thread
            Thread userInputHandler = new Thread(() -> handleUserInput());
            userInputHandler.start();

        } catch (Exception e) {
            System.out.println("Error: " + e.getMessage());
        }
    }

    private void handleClient() {
        try {
            BufferedReader socketInputStream = new BufferedReader(new InputStreamReader(socket.getInputStream()));

            while (socket.isConnected()) {
                String clientMessage = socketInputStream.readLine();
                if (clientMessage == null) {
                    System.out.println("Client disconnected");
                    break;
                }

                System.out.println("Client: " + clientMessage);
            }
        } catch (Exception e) {
            System.out.println("Error handling client: " + e.getMessage());
        }
    }

    private void handleUserInput() {
        try {
            BufferedReader input = new BufferedReader(new InputStreamReader(System.in));
            PrintWriter socketOutputStream = new PrintWriter(socket.getOutputStream(), true);

            while (socket.isConnected()) {
                String message = input.readLine();
                if (message.equals("quit")) {
                    System.out.println("Shutting down server...");
                    serverSocket.close();
                    socket.close();
                    break;
                }

                socketOutputStream.println(message);
            }
        } catch (Exception e) {
            System.out.println("Error handling user input: " + e.getMessage());
        }
    }

    public static void main(String[] args) {
        new Server2();
    }
}
```

And the client code :

```Java
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;

class Client2 {
    private Socket socket;

    public Client2() {
        try {
            System.out.println("Client started.");
            socket = new Socket("localhost", 420);

            // Handle server messages on a thread
            Thread serverListener = new Thread(() -> listenToServer());
            serverListener.start();

            // Handle the user input from the console on another thread
            Thread userInputHandler = new Thread(() -> handleUserInput());
            userInputHandler.start();

        } catch (Exception e) {
            System.out.println("Error: " + e.getMessage());
        }
    }

    private void listenToServer() {
        try {
            BufferedReader socketInputStream = new BufferedReader(new InputStreamReader(socket.getInputStream()));

            while (socket.isConnected()) {
                String serverMessage = socketInputStream.readLine();
                if (serverMessage == null) {
                    System.out.println("Server disconnected");
                    break;
                }

                System.out.println("Server: " + serverMessage);
            }
        } catch (Exception e) {
            System.out.println("Error listening to server: " + e.getMessage());
        }
    }

    private void handleUserInput() {
        try {
            BufferedReader input = new BufferedReader(new InputStreamReader(System.in));
            PrintWriter socketOutputStream = new PrintWriter(socket.getOutputStream(), true);

            while (socket.isConnected()) {
                String message = input.readLine();
                if (message.equals("quit")) {
                    System.out.println("Closing client...");
                    socket.close();
                    break;
                }

                socketOutputStream.println(message);
            }
        } catch (Exception e) {
            System.out.println("Error handling user input: " + e.getMessage());
        }
    }

    public static void main(String[] args) {
        new Client2();
    }
}
```

Let's run both programs and see the results :

<INSERT IMAGE>
