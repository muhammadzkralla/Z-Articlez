fun getString(): String {
    try {
        Thread.sleep(2000)
    } catch (e: Exception) {
        println(e.message)
    }

    return "Hello World!"
}

fun getInteger(): Int {
    try {
        Thread.sleep(2000)
    } catch (e: Exception) {
        println(e.message)
    }

    return 123
}

inline fun <reified T> doApiCall(
    x: Int,
    onComplete: (success: T?, failure: String?) -> Unit
) {
    if (x == 0) {
        val response = getString()

        try {
            onComplete(response as T, null)
        } catch (e: Exception) {
            onComplete(null, "Failed to cast response.")
        }
    } else {
        val response = getInteger()

        try {
            onComplete(response as T, null)
        } catch (e: Exception) {
            onComplete(null, "Failed to cast response.")
        }
    }
}

fun main() {
    doApiCall<String>(0, { success, failure ->
        success?.let { println(it) }
        failure?.let { println(it) }
    })

    doApiCall<Int>(0, { success, failure ->
        success?.let { println(it) }
        failure?.let { println(it) }
    })


    doApiCall<Int>(1, { success, failure ->
        success?.let { println(it) }
        failure?.let { println(it) }
    })

    doApiCall<String>(1, { success, failure ->
        success?.let { println(it) }
        failure?.let { println(it) }
    })

}
