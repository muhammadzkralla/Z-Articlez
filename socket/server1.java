import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.ServerSocket;
import java.net.Socket;

class Server1 {
    public Server1() {
        try {
            ServerSocket serverSocket = new ServerSocket(420);
            Socket socket = serverSocket.accept();

            System.out.println("Client connected");

            PrintWriter socketOutputStream = new PrintWriter(socket.getOutputStream(), true);
            BufferedReader socketInputStream = new BufferedReader(new InputStreamReader(socket.getInputStream()));

            while (socket.isConnected()) {
                BufferedReader input = new BufferedReader(new InputStreamReader(System.in));

                System.out.println(socketInputStream.readLine());

                String message = input.readLine();

                if (message.equals("quit")) {
                    serverSocket.close();
                    socket.close();
                    break;
                }

                socketOutputStream.println(message);
            }
        } catch (Exception e) {
            System.out.println("Could not init server socket!" + e.getMessage());
        }
    }

    public static void main(String[] args) {
        new Server1();
    }
}
