import java.util.Random;

interface Callback {
    void onComplete(String success, String failure);
}

class Demo2 {
    
    static int generateRandomNumber() {
        Random random = new Random();
        return random.nextInt(10);
    }

    static void doSomething(Callback listener) {
        
        // simulate a long-running task
        try {
            Thread.sleep(2000);
        } catch (Exception e) {
            listener.onComplete(null, "Could not sleep.");
        }

        // let x be the output of the long-running task
        int x = generateRandomNumber();

        String message = "X was " + x;
        if (x >= 5) {
            listener.onComplete(message, null);
        } else {
            listener.onComplete(null, message);
        }
    }

    public static void main(String[] args) {
        doSomething((success, failure) -> {
			if (success != null) System.out.println("Success, " + success);
		    else System.out.println("Failure, " + failure);
		});
    }
}
