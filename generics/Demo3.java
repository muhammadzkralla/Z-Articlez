import java.util.Arrays;
import java.util.List;

class Demo3 {
    public static void main(String[] args) {
        List<Integer> list1 = Arrays.asList(1, 2, 3, 4);
        printList(list1);

        List<Number> list2 = Arrays.asList(1, 2, 3, 4);
        printList(list2);
    }

    public static void printList(List<? super Integer> list) {
        System.out.println(list);
    }
}
