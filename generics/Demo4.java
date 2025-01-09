import java.util.List;

public class Demo4 {
    public static void printList(List<?> list) {
        for (Object element : list) {
            System.out.println(element); // Treated as 'Object'
        }
    }

    public static void main(String[] args) {
        List<Integer> intList = List.of(1, 2, 3);
        List<String> stringList = List.of("a", "b", "c");

        printList(intList);
        printList(stringList);
    }
}
