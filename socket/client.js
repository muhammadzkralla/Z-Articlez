const net = require("net");
const readline = require("readline");

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
});

const socket = new net.Socket();

socket.connect(420, "localhost", () => {
    console.log("Client started.");
});

socket.on("data", (data) => {
    const message = data.toString();
    console.log(`Server: ${message}`);
});

rl.on("line", (line) => {
    if (line === "quit") {
        console.log("Closing client...");
        socket.destroy();
    } else {
        socket.write(line);
    }
});
