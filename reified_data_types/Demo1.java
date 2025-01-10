class Demo1 {
    public static <T> boolean isTypeOf(Object obj, Class<T> clazz) {
        return clazz.isInstance(obj);
    }

    public static void main(String[] args) {
        System.out.println(isTypeOf("test", String.class));
        System.out.println(isTypeOf(2, String.class));
        System.out.println(isTypeOf(2, Integer.class));
    }
}
