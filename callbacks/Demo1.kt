import java.util.Random

fun generateRandomNumber(): Int {
    val random = Random()
    return random.nextInt(10)
}

fun doSomething(
    onSuccess: (message: String) -> Unit,
    onFailure: (message: String) -> Unit
) {
    
    // simulate a long-running task
    try {
        Thread.sleep(2000)
    } catch (e: Exception) {
        onFailure("Could not sleep.")
    }

    // let x be the output of the long-running task
    val x = generateRandomNumber()

    val message = "X was $x"
    if (x >= 5) {
        onSuccess("Success, $message")
    } else {
        onFailure("Failure, $message")
    }
}

fun main() {
    val x = generateRandomNumber()

    doSomething(
        { success ->
            println(success)
        },
        { failure ->
            println(failure)
        })
}
