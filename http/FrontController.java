import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.ConcurrentHashMap;

interface HttpHandler {
    HttpResponse handle(String method, String body);
}

class HttpResponse {
    public final int statusCode;
    public final String body;

    public HttpResponse(int statusCode, String body) {
        this.statusCode = statusCode;
        this.body = body;
    }
}

class FrontController {
    ServerSocket serverSocket;
    private final ConcurrentHashMap<String, HttpHandler> routes = new ConcurrentHashMap<>();

    public void start(int port) {
        try (ServerSocket serverSocket = new ServerSocket(port)) {
            System.out.println("Server started on port " + port);

            while (true) {
                Socket socket = serverSocket.accept();
                Thread clientHandler = new Thread(() -> handleClient(socket));
                clientHandler.start();
            }
        } catch (Exception ex) {
            System.out.println("Could not start service " + ex.getMessage());
        }
    }

    public void addRoute(String path, HttpHandler handler) {
        routes.put(path, handler);
    }

    private void handleClient(Socket socket) {
        try (BufferedReader in = new BufferedReader(new InputStreamReader(socket.getInputStream()));
                PrintWriter out = new PrintWriter(socket.getOutputStream(), true)) {

            String requestLine = in.readLine();
            if (requestLine == null || requestLine.isEmpty())
                return;

            String[] requestParts = requestLine.split(" ");
            if (requestParts.length < 3) {
                sendResponse(out, 400, "Bad Request");
                return;
            }

            String method = requestParts[0];
            String path = requestParts[1];
            int contentLength = 0;

            System.out.println("Incoming request: " + method + " " + path);

            List<String> headers = new ArrayList<>();
            String line;
            while (!(line = in.readLine()).isEmpty()) {
                headers.add(line);
                if (line.startsWith("Content-Length:")) {
                    contentLength = Integer.parseInt(line.split(": ")[1]);
                }
            }

            StringBuilder bodyBuilder = new StringBuilder();
            if (contentLength > 0) {
                char[] bodyChars = new char[contentLength];
                in.read(bodyChars, 0, contentLength);
                bodyBuilder.append(bodyChars);
            }

            String body = bodyBuilder.toString();

            HttpHandler handler = routes.get(path);
            if (handler == null) {
                sendResponse(out, 404, "Not Found");
                return;
            }

            HttpResponse response = handler.handle(method, body);
            sendResponse(out, response.statusCode, response.body);

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

    private void sendResponse(PrintWriter out, int statusCode, String body) {
        String statusLine = "HTTP/1.1 " + statusCode + " " + getStatusCode(statusCode) + "\r\n";
        String headers = "Content-Length: " + body.length() + "\r\n" +
                "Content-Type: text/plain\r\n\r\n";
        String response = (statusLine + headers + body);

        out.println(response);
    }

    private String getStatusCode(int statusCode) {
        return switch (statusCode) {
            case 200 -> "OK";
            case 201 -> "Created";
            case 400 -> "Bad Request";
            case 404 -> "Not Found";
            case 405 -> "Method Not Allowed";
            default -> "Internal Server Error";
        };
    }

    public static void main(String[] args) {
        FrontController framework = new FrontController();

        framework.addRoute("/hello", (method, body) -> {
            if ("GET".equals(method)) {
                return new HttpResponse(200, "Hello, World!");
            } else if ("POST".equals(method)) {
                return new HttpResponse(201, "You posted: " + body);
            }
            return new HttpResponse(406, "Method Not Allowed");
        });

        framework.start(420);
    }
}
