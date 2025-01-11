class Demo1 {
    public static void main(String[] args) {
        long startTime = System.currentTimeMillis();

        System.out.println("Starting blocking operation...");
        String result = performBlockingOperation();
        System.out.println("Blocking operation result: " + result);

        long endTime = System.currentTimeMillis();
        System.out.println("Time taken: " + (endTime - startTime) + " ms");
    }

    private static String performBlockingOperation() {
        try {
            // Simulate time-consuming I/O operation
            Thread.sleep(3000); // 3 seconds
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        return "Operation completed";
    }
}
