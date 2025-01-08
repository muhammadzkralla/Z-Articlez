import java.util.Random

fun generateRandomNumber(): Int {
    val random = Random()
    return random.nextInt(10)
}

fun doSomething(
    onComplete: (success: String?, failure: String?) -> Unit
) {

    // simulate a long-running task
    try {
        Thread.sleep(2000)
    } catch (e: Exception) {
        onComplete(null, "Could not sleep.")
    }

    // let x be the output of the long-running task
    val x = generateRandomNumber()

    val message = "X was $x"
    if (x >= 5) {
        onComplete("Success, $message", null)
    } else {
        onComplete(null, "Failure, $message")
    }
}

fun main() {
    doSomething() { success, failure ->
        success?.let { println(it) }
        failure?.let { println(it) }
    }
}