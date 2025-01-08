import java.util.Random;

interface Callback {
    void onSuccess(String message);
    void onFailure(String message);
}

class Demo1 {
    
    static int generateRandomNumber() {
        Random random = new Random();
        return random.nextInt(10);
    }

    static void doSomething(Callback listener) {
        
        // simulate a long-running task
        try {
            Thread.sleep(2000);
        } catch (Exception e) {
            listener.onFailure("Could not sleep.");
        }

        // let x be the output of the long-running task
        int x = generateRandomNumber();

        String message = "X was " + x;
        if (x >= 5) {
            listener.onSuccess("Success, " + message);
        } else {
            listener.onFailure("Failure, " + message);
        }
    }

    public static void main(String[] args) {
        doSomething(new Callback() {
            @Override
            public void onSuccess(String message) {
                System.out.println(message);
            }

            @Override
            public void onFailure(String message) {
                System.out.println(message);
            }
        });
    }
}
