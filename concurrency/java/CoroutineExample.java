import java.util.concurrent.*;

public class CoroutineExample {
    public static void main(String[] args) {
        try (ExecutorService executor = Executors.newVirtualThreadPerTaskExecutor()) {
            CompletableFuture<Void> c1 = CompletableFuture.runAsync(() -> {
                System.out.println("c1 started on " + Thread.currentThread());
                Thread.yield();
                System.out.println("c1 continued on " + Thread.currentThread());
            }, executor);

            CompletableFuture<Void> c2 = CompletableFuture.runAsync(() -> {
                System.out.println("c2 started on " + Thread.currentThread());
                Thread.yield();
                System.out.println("c2 continued on " + Thread.currentThread());
            }, executor);

            CompletableFuture.allOf(c1, c2).join(); // Wait for all tasks to complete
        }
    }
}
