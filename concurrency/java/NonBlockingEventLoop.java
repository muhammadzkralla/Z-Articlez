import java.util.PriorityQueue;
import java.util.Queue;

interface Callback {
    void onComplete(String success);
}

class NonBlockingEventLoop {
    private final Queue<ScheduledTask> taskQueue = new PriorityQueue<>();

    public void start() {
        long startTime = System.currentTimeMillis();

        while (true) {
            System.out.println("Rendering UI.... On: " + Thread.currentThread().getName());

            long now = System.currentTimeMillis();

            if (now - startTime >= 5000 && now - startTime <= 6000) {
                doRequest();
            }

            processTasks();

            sleep(1000);
        }
    }

    private void doRequest(Callback callback) {
        System.out.println("Initiated the request.... On: " + Thread.currentThread().getName());
        sleep(4000);
        System.out.println("Finished the request On: " + Thread.currentThread().getName());
    }

    private void doRequest() {
        System.out.println("Initiated the request.... On: " + Thread.currentThread().getName());
        setTimeout(() -> System.out.println("Finished the request On: " + Thread.currentThread().getName()),
                System.currentTimeMillis() + 4000);
    }

    private void setTimeout(Runnable task, long executionTime) {
        taskQueue.add(new ScheduledTask(task, executionTime));
    }

    private void processTasks() {
        long now = System.currentTimeMillis();
        while (!taskQueue.isEmpty() && taskQueue.peek().executionTime <= now) {
            taskQueue.poll().task.run();
        }
    }

    private void sleep(long time) {
        try {
            Thread.sleep(time);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    private static class ScheduledTask implements Comparable<ScheduledTask> {
        Runnable task;
        long executionTime;

        ScheduledTask(Runnable task, long executionTime) {
            this.task = task;
            this.executionTime = executionTime;
        }

        @Override
        public int compareTo(ScheduledTask other) {
            return Long.compare(this.executionTime, other.executionTime);
        }
    }
}
