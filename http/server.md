# HTTP Internals and Building Your Own Backend Framework From Scratch Part II

## Introduction

In the previous article, we talked about the structure of an HTTP request, how to build it from scratch, and we created a simple program that acts like a startup for an HTTP client library in Java. In this article, we will discuss the server side. We will talk about the structure of an HTTP response, how to build it from scratch, explain the front controller design pattern and how backend frameworks use it, and write a simple program that listens to the incomming requests and route them to their suitable handler or controller that will serve as a startup code to build your own backend framework.

It is recommended to read the previous article before you proceed in this article. If you understand the previous article well, this one will be more easy to you. I will skip the definition section that was discussed in details in the previous article.

You can refer back to the previous article from this [link](https://medium.com/@muhammad.heshamyt/http-internals-and-building-your-own-http-client-from-scratch-part-i-d62b1028408d).
