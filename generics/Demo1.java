import java.util.ArrayList;
import java.util.List;

class Container<T> {
    private List<T> items = new ArrayList<>();

    public void addItem(T item) {
        items.add(item);
    }

    public T getItem(int index) {
        return items.get(index);
    }

    public void print() {
        for (T item : items) {
            System.out.print(item + " ");
        }

        System.out.println();
    }
}

class Demo1 {
    public static void main(String[] args) {
        Container<String> cont1 = new Container<String>();

        cont1.addItem("one");
        cont1.addItem("two");
        cont1.addItem("three");

        System.out.println("Item at index 2 is " + cont1.getItem(2));

        cont1.print();

        Container<Integer> cont2 = new Container<Integer>();

        cont2.addItem(1);
        cont2.addItem(2);
        cont2.addItem(3);

        System.out.println("Item at index 2 is " + cont2.getItem(2));

        cont2.print();
    }
}
