const net = require("net");
const readline = require("readline");

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
});

const server = net.createServer((socket) => {
    console.log("Client connected.");

    socket.on("data", (data) => {
        const message = data.toString();
        console.log(`Client: ${message}`);
    });

    rl.on("line", (line) => {
        if (line === "quit") {
            console.log("Shutting down server...");
            socket.destroy();
        } else {
            socket.write(line);
        }
    });
});

server.listen(420, "localhost", () => {
    console.log("Server is listening on port 420");
});
