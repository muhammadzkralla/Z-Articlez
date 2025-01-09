fun sum(list: List<out Number>): Double {
    var sum = 0.0

    list.forEach { item ->
        sum += item.toDouble()
    }

    return sum
}

fun main() {
    val list1 = listOf<Int>(1, 2, 3, 4, 5)
    println("Total sum is: ${sum(list1)}")

    val list2 = listOf<Float>(1.0f, 2.0f, 3.0f, 4.0f, 5.0f)
    println("Total sum is: ${sum(list2)}")
}
