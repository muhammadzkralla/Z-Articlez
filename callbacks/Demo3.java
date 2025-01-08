
interface Callback<T> {
    void onComplete(T success, String failure);
}

class Demo3 {

    static String getString() {
        try {
            Thread.sleep(2000);
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }

        return "Hello World!";
    }

    static int getInteger() {
        try {
            Thread.sleep(2000);
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }

        return 123;
    }

    static <T> void doApiCall(int x, Callback<T> listener, Class<T> type) {
        if (x == 0) {
            String response = getString();

            try  {
                listener.onComplete(type.cast(response), null);
            } catch (Exception e) {
                listener.onComplete(null, "Failed to cast response.");
            }
        } else {
            int response = getInteger();

            try {
                listener.onComplete(type.cast(response), null);
            } catch (Exception e) {
                listener.onComplete(null, "Failed to cast response.");
            }
        }
    }

    public static void main(String[] args) {
        doApiCall(0, (success, failure) -> {
		    if (success != null) System.out.println("Success " + success);
		    else System.out.println("Failure " + failure);
		}, String.class);

        doApiCall(0, (success, failure) -> {
		    if (success != null) System.out.println("Success " + success);
		    else System.out.println("Failure " + failure);
		}, Integer.class);

        doApiCall(1, (success, failure) -> {
            if (success != null) System.out.println("Success " + success);
            else System.out.println("Failure " + failure);
        }, Integer.class);

        doApiCall(1, (success, failure) -> {
            if (success != null) System.out.println("Success " + success);
            else System.out.println("Failure " + failure);
        }, String.class);

    }
}
