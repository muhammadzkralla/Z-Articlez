import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;
import javax.net.ssl.SSLSocketFactory;

interface Callback<T> {
    void onComplete(T success, String failure);
}

class HttpClient {
    private Socket socket;
    private final String url;
    private final int port;
    private final boolean isHttps;

    public HttpClient(String url, int port) {
        this.url = url.replace("https://", "").replace("http://", "").split("/")[0];
        this.port = port;
        this.isHttps = url.startsWith("https://");
    }

    private void connect() throws IOException {
        if (isHttps) {
            socket = SSLSocketFactory.getDefault().createSocket(url, port);
        } else {
            socket = new Socket(url, port);
        }
    }

    public <T> void get(String endpoint, Callback<T> callback, Class<T> type) {

        try (BufferedReader in = new BufferedReader(new InputStreamReader(socket.getInputStream()));
                PrintWriter out = new PrintWriter(socket.getOutputStream(), true)) {
            connect();

            StringBuilder requestBuilder = new StringBuilder();
            requestBuilder.append("GET ").append(endpoint).append(" HTTP/1.1").append("\r\n");

            requestBuilder.append("Host: ").append(url).append("\r\n");
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
            callback.onComplete(null, "Error: " + ex.getMessage());
        } finally {
            try {
                if (socket != null)
                    socket.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }

    public <T> void post(String endpoint, Object body, Callback<T> callback, Class<T> type) {

        try (BufferedReader in = new BufferedReader(new InputStreamReader(socket.getInputStream()));
                PrintWriter out = new PrintWriter(socket.getOutputStream(), true)) {
            connect();

            String requestBody = body instanceof String ? (String) body : body.toString();
            int contentLength = requestBody.getBytes().length;

            StringBuilder requestBuilder = new StringBuilder();
            requestBuilder.append("POST ").append(endpoint).append(" HTTP/1.1").append("\r\n");

            requestBuilder.append("Host: ").append(url).append("\r\n");
            requestBuilder.append("User-Agent: zclient\r\n");
            requestBuilder.append("Accept: */*\r\n");
            requestBuilder.append("Content-Type: application/json\r\n");
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
            callback.onComplete(null, "Error: " + ex.getMessage());
        } finally {
            try {
                if (socket != null)
                    socket.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }

    public static void main(String[] args) {
        HttpClient client = new HttpClient("https://httpbin.org", 443);

        client.get("/get", (success, failure) -> {
            if (success != null)
                System.out.println("Success " + success);
            else
                System.out.println("Failure " + failure);
        }, String.class);

        String jsonBody = "{\"name\": \"John\", \"age\": 30}";
        client.post("/post", jsonBody, (success, failure) -> {
            if (success != null) {
                System.out.println("Success: " + success);
            } else {
                System.out.println("Failure: " + failure);
            }
        }, String.class);
    }
}
