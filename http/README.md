# HTTP Internals and Building Your Own HTTP Client From Scratch Part I

## Introduction

How do light bulbs emit light? According to a caveman, it sounds like black magic. However, when you understand basic physics, you know that there's no black magic happening here, it's just the flow of electric current through a filament that resists electricity causing it to heat up to very high temperatures and hence, produce light. Everything seems like black magic as long as you don't understand how it functions internally.

Magic is nothing but lack of sufficient knowledge.

In this article, I will explain the structure and the internals of an HTTP request, showing how simple it is, and that HTTP requests and responses are primarily plain text for headers and metadata, making them human-readable. However, the body can contain binary data, such as images or videos. In `HTTP/2` and `HTTP/3`, even headers are transmitted in binary format for efficiency. For simplicity, we will be discussing `HTTP/1.1` in this article.

We will also craft a small HTTP client program that performs GET and POST requests to a given API's url and prints the server's response.

The program in this article is simple and can be considered as an introdution to build your own HTTP client library from scratch. But if you are curious to see a real implementation of a Kotlin-based, type-safe, asynchronous/synchronous, and thread-safe HTTP Client Library built for the JVM supporting real-world features like threading, automatic serialization and deserialization, cancellation strategies, request-retry mechanisms, authentication, header intreception, and a lot more, you can visit my own HTTP client library on GitHub from this [link](https://github.com/muhammadzkralla/ZHttp).

## Definition

We, humans, can communicate with each other and exchange information. This process happens using some sort of language that both ends agree upon. This language acts as a standard that both ends understand, enabling them to extract information from each other.

Just like how humans talk and communicate with each other, HTTP messages are the mechanism used to exchange data between a server and a client in the HTTP protocol. There are many versions of HTTP like `HTTP/1.1`, `HTTP/2`, and `HTTP/3`. In this article, only `HTTP/1.1` is discussed, but I will attach a simple comparison between these versions later in the article.

HTTP is an application-layer protocol that operates on top of the TCP/IP suite. It follows a request-response model, where a client sends a request to a server, and the server returns a response. HTTP is stateless, meaning each request is independent of others, and the server does not retain any information about previous requests.

HTTP is the foundation of data communication for the web, enabling the retrieval of resources such as HTML documents, images, videos, and more. It is also the basis for more advanced protocols like HTTPS (HTTP Secure), which adds encryption for secure communication.

Luckily HTTP requests and responses are plain texts, so they are human-readable, which has its own pros and cons, this makes HTTP easier to understand, construct, and debug. However, it makes it vulnerable if it's not secured (HTTPS), as if a hacker intercepted the request, they can understand all the information this request/response carries which can be sensitive data, so it's highly recommended to use HTTPS not straight HTTP which will be discussed in the next article.

## Structure

Now let's see the HTTP request formatting. An HTTP request consists of three main sections.

1- Request Line <br>
2- Headers <br>
3- Body (optional) <br>

### Request Line

The request line consists of three parts, the request method, the endpoint, and the HTTP version. Let's explain each part of them:

#### Request Method

The first part of the request is called the request method, which can be one of these methods: POST, GET, DELETE, PUT, PATCH, HEAD, OPTIONS, TRACE, CONNECT, and more. This part of the request specifies the operation we want to perform, whether we need to send something (POST), retrieve something (GET), remove something (DELETE), update something (PUT, PATCH), and so on. Attached below is an illustration of each method.

<INSERT TABLE>

These methods are not obligatory or strictly enforced, meaning that you can technically use a POST request to update something or use a GET request to delete something. As these methods are some kind of standards that we agree upon to improve readability. This means that you can use any kind of method to do whatever operation you want as long as you and the backend side agreed to use this particular method for this particular operation.

While you technically can do this, it's always strongly recommended to use each method with its standard usage for better readability and interoperability.

#### Endpoint

The second part of the request line is the endpoint, which represents a specific resource or functionality provided by the server. You can think of it as the address where the server can find your required operation. You tell the server, 'I want you to call this specific function written by the backend team and send me the results back.' This function can be a database query, an image location, another HTTP request to another API, or any other operation.

For example, you can agree with your backend team that you want to get the list of all users when you send a GET HTTP request on endpoint `/users` and get a specific user when you send a GET HTTP request on endpoint `/users/{id}`, where `{id}` is replaced by the user's id. In this scenario the backend team will write some functions called for example `getUsers` and `getUser(id)` that perform some database queries and tell the server that these functions must be assigned to GET requests with endpoints `/users`, and `/users/{id}` respectively. Since this article is not discussing the server side and only the client side, I will not explain how is this done by the backend team in this article. The next article will be all about the server side where I will explain how this process is done by the backend frameworks using the front controller design pattern.

The url  can contain some queries, each query specifiy some parameter. For example, a url with queries like `/users?id=10&type=admin` specifies some id parameter with value 10 and a type parameter with value 'admin', this information can be extracted in the backend to modify the performed operation or apply some filter on the results based on these parameters.

#### HTTP Version

The third part of the request line is the HTTP version of the request. The first version of HTTP was `HTTP/0.9` which was extremely simple and only supported GET requests with no headers, status codes, or error handling. Later `HTTP/1.0` was introduced with support for headers, status codes, and other HTTP methods like POST and HEAD. The most common version used today is HTTP/1.1, which supports additional methods and features. `HTTP/2` and `HTTP/3` are the latest versions of HTTP with more and more features and improvements but still `HTTP/1.1` is the most used version which will be discussed in this article. Attached below is a comparison between these different versions.

<INSERT TABLE>

### Request Headers

The second section of the HTTP request is the request headers section. This section contains the meta-data of the request. Headers follow the `key: value` format, each header on a separate line and they are case-insensitive but the values are case-sensitive in some cases. They provide additional information about the request, such as the type of client, the format of the data being sent, authentication details, and more. Headers are essential for enabling features like content negotiation, caching, and security.

In our example, we will use three headers for GET requests and five headers for POST requests:

* Host: Specifies the domain name of the server (required in HTTP/1.1).
* User-Agent: Identifies the client (browser or application).
* Accept: Specifies the media types the client can process.
* Content-Type: Specifies the media type of the request body.
* Content-Length: Indicates the size of the request body in bytes.

HTTP headers play a crucial role in various aspects of web communication, including authentication, caching, request retries, sessions, and more. Below is a detailed breakdown of these topics and how headers are used to manage them:

HTTP headers are commonly used to send credentials or tokens for authentication. For basic authentication, a header like this is added to the request: `Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=`. And for OAuth and JWT, a header like this is added to the request: `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...`.

Headers can be used to manage retries for failed requests. For example, a header like this might be added to indicate how long the client should wait before retrying a request: `Retry-After: 120`

Other important use cases for headers might include content negotiation, CORS, security headers, rate limiting, and more. As we said earlier headers are the additional data (meta-data) about the current HTTP request so any additional piece of information you want to include for the current HTTP request, you don't add it directly to the request's data (body). Otherwise, you include a header for it.

### Request Body

The third and the last section of the HTTP request structure is the request body. The request body is the actual data or payload of the request. In this section you include the message you want to deliver or send to the server. But as we said earlier, there are some types of requests where we only need to retrieve data from the server and not to send data to the server, like the GET requests, or delete data from the server, like the DELETE requests, that's why this section (the request body) is optional and not required unless you need to send data to the server using something like POST, PUT, PATCH, and sometimes DELETE.

The body is usually encoded in JSON formatting, but you can use plain texts, XML, HTML, or even upload files!

Sometimes, you may want to separate your request into several parts. For example, you might send mixed data types in a single request, upload a file, upload a file with JSON objects, or upload multiple files and JSON objects. You might also divide your request into multiple parts for other reasons. To achieve this you can construct a `MULTIPART form data` request. In this kind of requests, you can upload files (the most common use case for it) or divide your request into multiple parts for any reason. Unfortunately `MULTIPART` requests are not going to be discussed in this article to keep this article concise.

## Building Your Own HTTP Client

Now that we have learned about the construction of an HTTP request, let's build a program that constructs the HTTP request from scratch. Please note that this article assumes that you have a solid understanding of callbacks, generic data types, and socket programming. If you are not, please refer to these articles first in order to achieve solid understanding of this program:

1- [Callbacks](https://medium.com/@muhammad.heshamyt/callbacks-and-lambdas-in-java-vs-kotlin-675ed3495a65) <br>
2- [Generics](https://medium.com/@muhammad.heshamyt/generics-wildcards-variance-and-star-projection-in-java-vs-kotlin-9553093485eb) <br>
3- [Socket Programming](https://medium.com/@muhammad.heshamyt/introduction-to-raw-socket-programming-and-implementing-the-simplest-chat-service-java-17a014703f5f) <br>

Now that you are ready, let's code!

Let's start by defining a callback that will be called once the reponse is received from the server, indicating success or failure:

```java
interface Callback<T> {
    void onComplete(T success, String failure);
}
```

Note that the success argument is generic, this means that the success body can be of any type, refer back to generics article if confused.

Now let's define our `HttpClient` class that will contain our program logic:

```java
class HttpClient {
    private Socket socket;
    private final String host;
    private final int port;
    private final boolean isHttps;

    public HttpClient(String url, int port) {
        // extract the host from the given url
        this.host = url.replace("https://", "").replace("http://", "").split("/")[0];
        this.port = port;
        this.isHttps = url.startsWith("https://");
    }
}
```

In the previous code, we created three properties of the `HttpClient` class, the host, port, and if it's intended for performing HTTP or HTTPS requests. These properties will be specified for each instance of the `HttpClient` class.

We also declared a `socket` variable that we will be using to connect to the server and write the created HTTP request in the server using the socket's output stream, refer back to the sockets article if confused.

Now let's write the `connect` function that will open a socket connection with the server using the provided host and port number to the current instance of the `HttpClient` class:

```java
    private void connect() {
        try {
            if (isHttps) {
                socket = SSLSocketFactory.getDefault().createSocket(host, port);
            } else {
                socket = new Socket(host, port);
            }
        } catch (Exception e) {
            System.out.println("Error: " + e.getMessage());
        }
    }
```

Fine, until now, we created the callback and the connection logic. In the next part, we will write the function responsible for constructing a `GET` request and write it to the server using the socket's output stream and then read the server's response using the socket's input stream and call the callback with the response:

```java
    // the function takes three arguments, the targeted endpoint, the callback,
    // and the type of the expected response
    public <T> void get(String endpoint, Callback<T> callback, Class<T> type) {

        connect();

        // open the input and output streams of the socket after connection
        try (BufferedReader in = new BufferedReader(new InputStreamReader(socket.getInputStream()));
                PrintWriter out = new PrintWriter(socket.getOutputStream(), true)) {


            // the request builder will contain our request details
            StringBuilder requestBuilder = new StringBuilder();

            // the request line
            requestBuilder.append("GET ").append(endpoint).append(" HTTP/1.1").append("\r\n");

            // the headers
            requestBuilder.append("Host: ").append(host).append("\r\n");
            requestBuilder.append("User-Agent: zclient\r\n");
            requestBuilder.append("Accept: */*\r\n");

            // the empty line
            requestBuilder.append("\r\n");

            // write the request to the socket's output stream
            String request = requestBuilder.toString();
            out.println(request);

            // start to receive the server's response from the socket's input stream
            String statusLine = in.readLine();
            if (statusLine == null || !statusLine.startsWith("HTTP/1.1")) {
                callback.onComplete(null, "Invalid response from server");
                return;
            }

            // the status code is the second part of the status line
            // note that the response format will be discussed in the next article
            // about the server side of HTTP requests
            int statusCode = Integer.parseInt(statusLine.split(" ")[1]);

            // start reading the response headers
            String line;
            StringBuilder headersBuilder = new StringBuilder();
            while ((line = in.readLine()) != null && !line.isEmpty()) {
                headersBuilder.append(line).append("\n");
            }

            // start reading the response body
            StringBuilder bodyBuilder = new StringBuilder();
            while ((line = in.readLine()) != null) {
                bodyBuilder.append(line).append("\n");
            }

            String responseBody = bodyBuilder.toString();

            // call the callback with the suitable data either success or failure
            if (statusCode >= 200 && statusCode < 300) {
                if (type == String.class) {
                    callback.onComplete(type.cast(responseBody), null);
                } else {
                    callback.onComplete(null, "Unsupported response type");
                }
            } else {
                callback.onComplete(null, "HTTP Error " + statusCode + ": " + responseBody);
            }
        } catch (Exception ex) {
            callback.onComplete(null, "Error: " + ex.getMessage());
        } finally {
            // close the connection with server at the end
            try {
                if (socket != null)
                    socket.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
```

I explained each code block with some comments, but please note that the response structure will be discussed in details in the next article about the server side of HTTP.

Similarly, the function that performs the POST request will be identical to the `GET` function in addition to the request body and two more headers of the content type and content length:

```java
    public <T> void post(String endpoint, Object body, Callback<T> callback, Class<T> type) {

        connect();

        try (BufferedReader in = new BufferedReader(new InputStreamReader(socket.getInputStream()));
                PrintWriter out = new PrintWriter(socket.getOutputStream(), true)) {

            String requestBody = body instanceof String ? (String) body : body.toString();
            int contentLength = requestBody.getBytes().length;

            StringBuilder requestBuilder = new StringBuilder();

            requestBuilder.append("POST ").append(endpoint).append(" HTTP/1.1").append("\r\n");

            requestBuilder.append("Host: ").append(host).append("\r\n");
            requestBuilder.append("User-Agent: zclient\r\n");
            requestBuilder.append("Accept: */*\r\n");
            // the additional headers
            requestBuilder.append("Content-Type: application/json\r\n");
            requestBuilder.append("Content-Length: ").append(contentLength).append("\r\n");

            requestBuilder.append("\r\n");

            requestBuilder.append(requestBody);

            String request = requestBuilder.toString();
            out.println(request);

            String statusLine = in.readLine();
            if (statusLine == null || !statusLine.startsWith("HTTP/1.1")) {
                callback.onComplete(null, "Invalid response from server");
                return;
            }

            int statusCode = Integer.parseInt(statusLine.split(" ")[1]);

            String line;
            StringBuilder headersBuilder = new StringBuilder();
            while ((line = in.readLine()) != null && !line.isEmpty()) {
                headersBuilder.append(line).append("\n");
            }

            StringBuilder bodyBuilder = new StringBuilder();
            while ((line = in.readLine()) != null) {
                bodyBuilder.append(line).append("\n");
            }

            String responseBody = bodyBuilder.toString();

            if (statusCode >= 200 && statusCode < 300) {
                if (type == String.class) {
                    callback.onComplete(type.cast(responseBody), null);
                } else {
                    callback.onComplete(null, "Unsupported response type");
                }
            } else {
                callback.onComplete(null, "HTTP Error " + statusCode + ": " + responseBody);
            }
        } catch (Exception ex) {
            callback.onComplete(null, "Error: " + ex.getMessage());
        } finally {
            try {
                if (socket != null)
                    socket.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
```

Now the only part that's left is to test our program in the `main` method with some API:

```java
    public static void main(String[] args) {
        // create an instance of the HttpClient class with the API's url
        // and port 433 is the standard port for HTTPS
        HttpClient client = new HttpClient("https://httpbin.org", 443);

        // perform a GET request and print the result
        client.get("/get", (success, failure) -> {
            if (success != null)
                System.out.println("Success " + success);
            else
                System.out.println("Failure " + failure);
        }, String.class);

        // perform a POST request and print the result
        String jsonBody = "{\"name\": \"John\", \"age\": 30}";
        client.post("/post", jsonBody, (success, failure) -> {
            if (success != null) {
                System.out.println("Success: " + success);
            } else {
                System.out.println("Failure: " + failure);
            }
        }, String.class);
    }
```

Please refer back to the callbacks article if confused. Now let's run our program and see our results:

```bash
❯ javac HttpClient.java

❯ java HttpClient
Success {
  "args": {},
  "headers": {
    "Accept": "*/*",
    "Host": "httpbin.org",
    "User-Agent": "zclient",
    "X-Amzn-Trace-Id": "Root=1-678ffad3-634405b816c191854b5fa1cd"
  },
  "origin": "<MY-IP-ADDRESS>",
  "url": "https://httpbin.org/get"
}

Success: {
  "args": {},
  "data": "{\"name\": \"John\", \"age\": 30}",
  "files": {},
  "form": {},
  "headers": {
    "Accept": "*/*",
    "Content-Length": "27",
    "Content-Type": "application/json",
    "Host": "httpbin.org",
    "User-Agent": "zclient",
    "X-Amzn-Trace-Id": "Root=1-678ffb10-3863c9f4771a5bb611a89973"
  },
  "json": {
    "age": 30,
    "name": "John"
  },
  "origin": "<MY-IP-ADDRESS>",
  "url": "https://httpbin.org/post"
}
```

As we can see, the program has performed successfull GET and POST requests to the given `httpbin.org` API and printed the received reponse from the server.

Please note that I wrote this code for illustration purposes only. It is considered boilerplate and lacks refactoring to avoid duplicate code, testing, and more. It can be considered as a simple introdution to build your own HTTP client library, but an actual HTTP client library must contain more than just sending and receiving HTTP requests. A real HTTP client library should handle other HTTP methods, threading, serialization & deserialization, headers interception, request-retry mechanisms, cancellation strategies, logging, error handling, and really a lot more. If you are curious to explore more, you can visit my full HTTP client library implementation from scratch on GitHub from this [link](https://github.com/muhammadzkralla/ZHttp).
