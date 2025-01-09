fun main() {
    val list1: List<Int> = listOf(1, 2, 3, 4)
    printList(list1)

    val list2: List<Number> = listOf(1, 2, 3, 4)
    printList(list2)
}

fun printList(list: List<in Int>) {
    println(list)
}
