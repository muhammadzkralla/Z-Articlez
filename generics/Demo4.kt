fun printList(list: List<*>) {
    for (element in list) {
        println(element) // Elements are of type 'Any?'
    }
}

fun main() {
    val intList: List<Int> = listOf(1, 2, 3)
    val stringList: List<String> = listOf("a", "b", "c")

    printList(intList)
    printList(stringList)
}
