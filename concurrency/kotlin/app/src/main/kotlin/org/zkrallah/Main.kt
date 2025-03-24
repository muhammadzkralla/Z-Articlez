package org.zkrallah

import com.zkrallah.zhttp.client.ZHttpClient
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.yield
import kotlin.coroutines.resume
import kotlin.coroutines.suspendCoroutine

val client = ZHttpClient.Builder()
    .baseUrl("https://jsonplaceholder.typicode.com")
    .allowLogging(true)
    .build()

fun main() {
    runBlocking {
        launch {
            counter()
        }
        launch {
            counter()
        }
        // launch {
        //     println("c1 started on ${Thread.currentThread().name}")
        //     val response = client.get<String>("posts")
        //     println("c1 continued with ${response?.code} on ${Thread.currentThread().name}")
        // }
        // launch {
        //     println("c2 started on ${Thread.currentThread().name}")
        //     val response = client.get<String>("posts")
        //     println("c2 continued with ${response?.code} on ${Thread.currentThread().name}")
        // }
        // launch {
        //     println("c3 started on ${Thread.currentThread().name}")
        //     val response = client.get<String>("posts")
        //     println("c3 continued with ${response?.code} on ${Thread.currentThread().name}")
        // }
//        launch {
//            renderUI()
//        }
//        launch {
//            delay(5000)
//            doRequest()
//        }
    }
}

suspend fun counter() {
    for (i in 0..9) {
        println(i)
        yield()
    }
}

suspend fun renderUI() {
    while (true) {
        println("Rendering UI On... ${Thread.currentThread().name}")
        delay(1000)
        yield()
    }
}

suspend fun doRequest() {
    println("Initiated the Request On... ${Thread.currentThread().name}")
    delay(3000)
    println("Finished the Request On... ${Thread.currentThread().name}")
}

suspend fun customSuspendFunction(): String = suspendCoroutine { continuation ->
    println("Suspended... Now resuming manually.")
    continuation.resume("Hello from continuation!") // Manually resuming
}
