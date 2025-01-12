import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;

class Client1 {
    public Client1() {
        try {
            System.out.println("Client started.");
            Socket socket = new Socket("localhost", 420);

            PrintWriter socketOutputStream = new PrintWriter(socket.getOutputStream(), true);
            BufferedReader socketInputStream = new BufferedReader(new InputStreamReader(socket.getInputStream()));

            while (socket.isConnected()) {
                BufferedReader input = new BufferedReader(new InputStreamReader(System.in));

                String message = input.readLine();

                if (message.equals("quit")) {
                    socket.close();
                    break;
                }

                socketOutputStream.println(message);

                System.out.println(socketInputStream.readLine());
            }

        } catch (Exception e) {
            System.out.println("Could not init client socket!" + e.getMessage());
        }
    }

    public static void main(String[] args) {
        new Client1();
    }
}
