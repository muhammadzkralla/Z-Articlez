class Main {
    public static void main(String[] args) {
        Dispatcher dispatcher = new Dispatcher();

        Coroutine coroutine1 = new Coroutine(() -> {
            System.out.println("Coroutine 1 - Start");
            Coroutine.yield();
            System.out.println("Coroutine 1 - Resumed");
        }, dispatcher);

        Coroutine coroutine2 = new Coroutine(() -> {
            System.out.println("Coroutine 2 - Start");
            Coroutine.yield();
            System.out.println("Coroutine 2 - Resumed");
        }, dispatcher);

        Dispatcher.addCoroutine(coroutine1);
        Dispatcher.addCoroutine(coroutine2);

        Dispatcher.run();
    }
}
