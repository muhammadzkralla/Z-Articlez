import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;

interface Callback<T> {
    void onComplete(T success, String failure);
}

class HttpClient {
    private Socket socket;
    private final String url;
    private final int port;

    public HttpClient(String url, int port) {
        this.url = url;
        this.port = port;
    }

    private void connect() {
        try {
            socket = new Socket(url, port);
        } catch (Exception e) {
            System.out.println("Error: " + e.getMessage());
        }
    }

    public <T> void get(String url, Callback<T> callback, Class<T> type) {
        connect();

        try (BufferedReader in = new BufferedReader(new InputStreamReader(socket.getInputStream()));
                PrintWriter out = new PrintWriter(socket.getOutputStream(), true)) {

            StringBuilder requestBuilder = new StringBuilder();
            requestBuilder.append("GET ").append(url).append(" HTTP/1.1");

            requestBuilder.append("\r\n");

            requestBuilder.append("Host: localhost:420\r\n");
            requestBuilder.append("User-Agent: zclient\r\n");
            requestBuilder.append("Accept: */*\r\n");

            requestBuilder.append("\r\n");

            String request = requestBuilder.toString();

            out.println(request);

            String statusLine = in.readLine();
            if (statusLine == null || !statusLine.startsWith("HTTP/1.1")) {
                callback.onComplete(null, "Invalid response from server");
                return;
            }

            int statusCode = Integer.parseInt(statusLine.split(" ")[1]);

            String line;
            StringBuilder headersBuilder = new StringBuilder();
            while ((line = in.readLine()) != null && !line.isEmpty()) {
                headersBuilder.append(line).append("\n");
            }

            StringBuilder bodyBuilder = new StringBuilder();
            while ((line = in.readLine()) != null) {
                bodyBuilder.append(line).append("\n");
            }

            String responseBody = bodyBuilder.toString();

            if (statusCode >= 200 && statusCode < 300) {
                if (type == String.class) {
                    callback.onComplete(type.cast(responseBody), null);
                } else {
                    callback.onComplete(null, "Unsupported response type");
                }
            } else {
                callback.onComplete(null, "HTTP Error " + statusCode + ": " + responseBody);
            }
        } catch (Exception ex) {
            System.out.println("Could not hanlde client " + ex.getMessage());
            try {
                socket.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        } finally {
            try {
                socket.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }

    public <T> void post(String url, Object body, Callback<T> callback, Class<T> type) {
        connect();

        try (BufferedReader in = new BufferedReader(new InputStreamReader(socket.getInputStream()));
                PrintWriter out = new PrintWriter(socket.getOutputStream(), true)) {

            String requestBody = body instanceof String ? (String) body : body.toString();
            int contentLength = requestBody.getBytes().length;

            StringBuilder requestBuilder = new StringBuilder();
            requestBuilder.append("POST ").append(url).append(" HTTP/1.1");

            requestBuilder.append("\r\n");

            requestBuilder.append("Host: localhost:420\r\n");
            requestBuilder.append("User-Agent: zclient\r\n");
            requestBuilder.append("Accept: */*\r\n");
            requestBuilder.append("Content-Type: application/x-www-form-urlencoded\r\n");
            requestBuilder.append("Content-Length: ").append(contentLength).append("\r\n");

            requestBuilder.append("\r\n");

            requestBuilder.append(requestBody);

            String request = requestBuilder.toString();

            out.println(request);

            String statusLine = in.readLine();
            if (statusLine == null || !statusLine.startsWith("HTTP/1.1")) {
                callback.onComplete(null, "Invalid response from server");
                return;
            }

            int statusCode = Integer.parseInt(statusLine.split(" ")[1]);

            String line;
            StringBuilder headersBuilder = new StringBuilder();
            while ((line = in.readLine()) != null && !line.isEmpty()) {
                headersBuilder.append(line).append("\n");
            }

            StringBuilder bodyBuilder = new StringBuilder();
            while ((line = in.readLine()) != null) {
                bodyBuilder.append(line).append("\n");
            }

            String responseBody = bodyBuilder.toString();

            if (statusCode >= 200 && statusCode < 300) {
                if (type == String.class) {
                    callback.onComplete(type.cast(responseBody), null);
                } else {
                    callback.onComplete(null, "Unsupported response type");
                }
            } else {
                callback.onComplete(null, "HTTP Error " + statusCode + ": " + responseBody);
            }
        } catch (Exception ex) {
            System.out.println("Could not hanlde client " + ex.getMessage());
            try {
                socket.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        } finally {
            try {
                socket.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }

    public static void main(String[] args) {
        HttpClient client = new HttpClient("localhost", 420);

        client.get("/hello", (success, failure) -> {
            if (success != null)
                System.out.println("Success " + success);
            else
                System.out.println("Failure " + failure);
        }, String.class);

        client.post("/hello", "test", (success, failure) -> {
            if (success != null)
                System.out.println("Success " + success);
            else
                System.out.println("Failure " + failure);
        }, String.class);

    }
}
