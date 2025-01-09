class Container<T> {
    val items = mutableListOf<T>()

    fun addItem(item: T) {
        items.add(item)
    }

    fun getItem(index: Int): T {
        return items[index]
    }

    fun print() {
        items.forEach { item -> 
            print("$item ")
        }

        println()
    }
}

fun main() {
    val cont1 = Container<String>()

    cont1.addItem("one")
    cont1.addItem("two")
    cont1.addItem("three")

    println("Item at index 2 is ${cont1.getItem(2)}")

    cont1.print()

    val cont2 = Container<Int>()

    cont2.addItem(1)
    cont2.addItem(2)
    cont2.addItem(3)

    println("Item at index 2 is ${cont2.getItem(2)}")

    cont2.print()

}
