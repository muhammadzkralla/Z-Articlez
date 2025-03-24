import java.util.ArrayList;
import java.util.List;

class Dispatcher {
    private static final List<Coroutine> coroutines = new ArrayList<>();
    private static int currentIndex = 0;

    public static void addCoroutine(Coroutine coroutine) {
        coroutines.add(coroutine);
    }

    public static void yield() {
        // Save the current coroutine's state (if needed)
        // Switch to the next coroutine
        if (!coroutines.isEmpty()) {
            currentIndex = (currentIndex + 1) % coroutines.size();
            Coroutine nextCoroutine = coroutines.get(currentIndex);
            nextCoroutine.run();
        }
    }

    public static void run() {
        while (!coroutines.isEmpty()) {
            Coroutine currentCoroutine = coroutines.get(currentIndex);
            if (!currentCoroutine.isDone()) {
                currentCoroutine.run();
            } else {
                coroutines.remove(currentIndex);
                if (currentIndex >= coroutines.size()) {
                    currentIndex = 0;
                }
            }
        }
    }
}
