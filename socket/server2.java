import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.ServerSocket;
import java.net.Socket;

class Server2 {
    private ServerSocket serverSocket;
    private Socket socket;

    public Server2() {
        try {
            serverSocket = new ServerSocket(420);
            socket = serverSocket.accept();

            System.out.println("Client connected");

            Thread clientHandler = new Thread(() -> handleClient());
            clientHandler.start();

            Thread userInputHandler = new Thread(() -> handleUserInput());
            userInputHandler.start();

        } catch (Exception e) {
            System.out.println("Error: " + e.getMessage());
        }
    }

    private void handleClient() {
        try {
            BufferedReader socketInputStream = new BufferedReader(new InputStreamReader(socket.getInputStream()));

            while (socket.isConnected()) {
                String clientMessage = socketInputStream.readLine();
                if (clientMessage == null) {
                    System.out.println("Client disconnected");
                    break;
                }

                System.out.println("Client: " + clientMessage);
            }
        } catch (Exception e) {
            System.out.println("Error handling client: " + e.getMessage());
        }
    }

    private void handleUserInput() {
        try {
            BufferedReader input = new BufferedReader(new InputStreamReader(System.in));
            PrintWriter socketOutputStream = new PrintWriter(socket.getOutputStream(), true);

            while (socket.isConnected()) {
                String message = input.readLine();
                if (message.equals("quit")) {
                    System.out.println("Shutting down server...");
                    serverSocket.close();
                    socket.close();
                    break;
                }

                socketOutputStream.println(message);
            }
        } catch (Exception e) {
            System.out.println("Error handling user input: " + e.getMessage());
        }
    }

    public static void main(String[] args) {
        new Server2();
    }
}
