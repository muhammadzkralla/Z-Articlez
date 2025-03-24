class Blocking {
    public void start() {
        long startTime = System.currentTimeMillis();
        while (true) {
            System.out.println("Rendring UI...." + " On: " + Thread.currentThread().getName());
            long now = System.currentTimeMillis();

            if (now - startTime >= 5000 && now - startTime <= 6000) {
                doRequest();
            }

            sleep(1000);
        }

    }

    private void doRequest() {
        System.out.println("Initiated the request...." + " On: " + Thread.currentThread().getName());
        sleep(4000);
        System.out.println("Finished the request On: " + Thread.currentThread().getName());

    }

    private void sleep(long time) {
        try {
            Thread.sleep(time);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}
