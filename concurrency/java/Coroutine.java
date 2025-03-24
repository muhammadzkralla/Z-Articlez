class Coroutine {
    private final Runnable task;
    private final Dispatcher dispatcher;
    private boolean isDone = false;

    public Coroutine(Runnable task, Dispatcher dispatcher) {
        this.task = task;
        this.dispatcher = dispatcher;
    }

    public void run() {
        task.run();
        isDone = true;
    }

    public boolean isDone() {
        return isDone;
    }

    public static void yield() {
        // Yield control back to the dispatcher
        Dispatcher.yield();
    }
}
