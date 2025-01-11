import java.util.concurrent.CompletableFuture;
import java.util.concurrent.CountDownLatch;

class Demo2 {
    public static void main(String[] args) {
        long startTime = System.currentTimeMillis();
        CountDownLatch latch = new CountDownLatch(1);

        // Start non-blocking operation
        System.out.println("We are now on " + Thread.currentThread().getName());
        System.out.println("Starting non-blocking operation...");
        CompletableFuture<String> futureResult = performNonBlockingOperation();

        // this is to handle when the non-blocking operation finishes
        futureResult.whenComplete((res, t) -> {
            System.out.println(res);
        });

        long endTime = System.currentTimeMillis();
        System.out.println("Time taken: " + (endTime - startTime) + " ms");

        // this tells the program not to terminate before the futureResult resolves
        try {
			latch.await();
		} catch (InterruptedException e) {
			e.printStackTrace();
		} // Wait until the countdown reaches 0
    }

    private static CompletableFuture<String> performNonBlockingOperation() {
        return CompletableFuture.supplyAsync(() -> {
            System.out.println("performing non-blocking operation on " + Thread.currentThread().getName());
            try {
                // Simulate time-consuming I/O operation
                Thread.sleep(3000); // 3 seconds
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

            return "non-blocking operation completed";
        });
    }
}
