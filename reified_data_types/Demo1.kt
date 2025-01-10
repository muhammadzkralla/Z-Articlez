inline fun <reified T> isTypeOf(value: Any): Boolean {
    return value is T
}

fun main() {
    println(isTypeOf<String>("test"))
    println(isTypeOf<String>(2))
    println(isTypeOf<Int>(2))
}
