import java.util.*;

class Coroutine1 {
    private static final Queue<Runnable> eventLoop = new LinkedList<>();
    private Runnable task;
    
    public Coroutine1(Runnable task) {
        this.task = () -> {
            task.run();
        };
    }

    public void resume() {
        eventLoop.add(task);
    }

    public static void yield() {
        // TODO:
        // 1- save the current coroutine state in a continuation object (state machine)
        // 2- mimic function termination and yield the thread of the current coroutine dispatcher to other coroutines on the same dispatcher to run if possible
        if (!eventLoop.isEmpty()) {
            Runnable nextTask = eventLoop.poll();
            if (nextTask != null) nextTask.run();
        }
    }

    public static void runEventLoop() {
        while (!eventLoop.isEmpty()) {
            Runnable nextTask = eventLoop.poll();
            if (nextTask != null) nextTask.run();
        }
    }
}
