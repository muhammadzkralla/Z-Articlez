import java.util.Arrays;
import java.util.List;

class Demo2 {
    private static double sum(List<? extends Number> list) {
        double sum = 0.0;
        for (Number i : list) {
            sum += i.doubleValue();
        }

        return sum;
    }

    public static void main(String[] args) {
        List<Integer> list1 = Arrays.asList(1, 2, 3, 4, 5);
        System.out.println("Total sum is: " + sum(list1));

        List<Float> list2 = Arrays.asList(1.0f, 2.0f, 3.0f, 4.0f, 5.0f);
        System.out.println("Total sum is: " + sum(list2));
    }
}
