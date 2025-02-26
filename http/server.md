# HTTP Internals and Building Your Own Backend Framework From Scratch Part II

## Introduction

In the previous article, we talked about the structure of an HTTP request, how to build it from scratch, and we created a simple program that acts like a startup for an HTTP client library in Java. In this article, we will discuss the server side. We will talk about the structure of an HTTP response, how to build it from scratch, explain the front controller design pattern and how backend frameworks use it, and write a simple program that listens to the incomming requests and route them to their suitable handler or controller that will serve as a startup code to build your own backend framework.

It is recommended to read the previous article before you proceed in this article. If you understand the previous article well, this one will be more easy to you, as I will skip some parts that were discussed in details in the previous article.

You can refer back to the previous article from this [link](https://medium.com/@muhammad.heshamyt/http-internals-and-building-your-own-http-client-from-scratch-part-i-d62b1028408d).

## Structure

Now let's see the HTTP response formatting. An HTTP response consists of three main sections.

1- Status Line <br>
2- Headers <br>
3- Body <br>

### Status Line

The response line consists of three parts, the HTTP version, the status code, and the status text. Let's explain each part of them:

#### HTTP Version

The first part of the status line is the HTTP version of the request. The first version of HTTP was `HTTP/0.9` which was extremely simple and only supported GET requests with no headers, status codes, or error handling. Later `HTTP/1.0` was introduced with support for headers, status codes, and other HTTP methods like POST and HEAD. The most common version used today is HTTP/1.1, which supports additional methods and features. `HTTP/2` and `HTTP/3` are the latest versions of HTTP with more and more features and improvements but still `HTTP/1.1` is the most used version which will be discussed in this article. Attached below is a comparison between these different versions.

<INSERT TABLE>

#### Status Code

The status code is a three-digit number that represents the outcome of the request. It tells the client whether the request was successful, redirected, or encountered an error.
* codes between 100-199 are informational (request received, processing continues).
* Codes between 200-299 indicates success (request was successfully received, understood, and accepted).
* Codes between 300-399 indicates redirection (further action is needed to complete the request).
* Codes between 400-499 indicates a client error (the request has an issue from the client's side).
* Codes between 500-599 indicates a server error (the server failed to process a valid request).

Here's a [link](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status) where you can find a list of the most common status codes and their meaning.

#### Status Text

The third part of the status line is the status text. It is a human-readable description of the status code. It provides additional clarity but is not strictly required for HTTP processing.

Examples of Status Codes with Reason Phrases:
* 200 OK – The request was successful.
* 201 Created – A new resource was successfully created.
* 301 Moved Permanently – The requested resource has been moved to a new URL.
* 403 Forbidden – The client does not have permission to access the resource.
* 500 Internal Server Error – The server encountered an unexpected condition.

### Response Headers

The second section of the HTTP response is the response headers section. This section contains the meta-data of the response. Headers follow the `key: value` format, each header on a separate line and they are case-insensitive but the values are case-sensitive in some cases. They provide additional information about the request, such as the type of client, the format of the data being sent, authentication details, and more. Headers are essential for enabling features like content negotiation, caching, and security.

In our example, we will use three headers for GET requests and five headers for POST requests:

* Content-Type: Specifies the media type of the request body.
* Content-Length: Indicates the size of the request body in bytes.

HTTP headers play a crucial role in various aspects of web communication, including authentication, caching, request retries, sessions, and more. Below is a detailed breakdown of these topics and how headers are used to manage them:

HTTP headers are commonly used to send credentials or tokens for authentication. For basic authentication, a header like this is added to the request: `Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=`. And for OAuth and JWT, a header like this is added to the request: `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...`.

Headers can be used to manage retries for failed requests. For example, a header like this might be added to indicate how long the client should wait before retrying a request: `Retry-After: 120`

Other important use cases for headers might include content negotiation, CORS, security headers, rate limiting, and more. As we said earlier headers are the additional data (meta-data) about the current HTTP request so any additional piece of information you want to include for the current HTTP request, you don't add it directly to the request's data (body). Otherwise, you include a header for it.

### Response Body

The third and the last section of the HTTP request structure is the request body. The request body is the actual data or payload of the request. In this section you include the message you want to deliver or send to the server. But as we said earlier, there are some types of requests where we only need to retrieve data from the server and not to send data to the server, like the GET requests, or delete data from the server, like the DELETE requests, that's why this section (the request body) is optional and not required unless you need to send data to the server using something like POST, PUT, PATCH, and sometimes DELETE.

The body is usually encoded in JSON formatting, but you can use plain texts, XML, HTML, or even upload files!

Sometimes, you may want to separate your request into several parts. For example, you might send mixed data types in a single request, upload a file, upload a file with JSON objects, or upload multiple files and JSON objects. You might also divide your request into multiple parts for other reasons. To achieve this you can construct a `MULTIPART form data` request. In this kind of requests, you can upload files (the most common use case for it) or divide your request into multiple parts for any reason. Unfortunately `MULTIPART` requests are not going to be discussed in this article to keep this article concise.
