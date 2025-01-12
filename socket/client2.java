import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;

class Client2 {
    private Socket socket;

    public Client2() {
        try {
            System.out.println("Client started.");
            socket = new Socket("localhost", 420);

            Thread serverListener = new Thread(() -> listenToServer());
            serverListener.start();

            Thread userInputHandler = new Thread(() -> handleUserInput());
            userInputHandler.start();

        } catch (Exception e) {
            System.out.println("Error: " + e.getMessage());
        }
    }

    private void listenToServer() {
        try {
            BufferedReader socketInputStream = new BufferedReader(new InputStreamReader(socket.getInputStream()));

            while (socket.isConnected()) {
                String serverMessage = socketInputStream.readLine();
                if (serverMessage == null) {
                    System.out.println("Server disconnected");
                    break;
                }

                System.out.println("Server: " + serverMessage);
            }
        } catch (Exception e) {
            System.out.println("Error listening to server: " + e.getMessage());
        }
    }

    private void handleUserInput() {
        try {
            BufferedReader input = new BufferedReader(new InputStreamReader(System.in));
            PrintWriter socketOutputStream = new PrintWriter(socket.getOutputStream(), true);

            while (socket.isConnected()) {
                String message = input.readLine();
                if (message.equals("quit")) {
                    System.out.println("Closing client...");
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
        new Client2();
    }
}
