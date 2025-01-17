const net = require('net');

class FrontController {
    constructor() {
        this.routes = new Map();
    }

    start(port) {
        const server = net.createServer(socket => this.handleClient(socket));

        server.listen(port, () => {
            console.log(`Server started on port ${port}`);
        });
    }

    addRoute(path, handler) {
        this.routes.set(path, handler);
    }

    handleClient(socket) {
        let requestData = '';
        socket.on('data', data => {
            requestData += data.toString();

            if (requestData.includes('\r\n\r\n')) {
                this.processRequest(socket, requestData);
            }
        });

        socket.on('error', err => {
            console.error(`Socket error: ${err.message}`);
        });
    }

    processRequest(socket, requestData) {
        const [requestLine] = requestData.split('\r\n');
        const [method, path] = requestLine.split(' ');

        console.log(`Incoming request: ${method} ${path}`);

        const headers = {};
        let body = '';

        const headerEndIndex = requestData.indexOf('\r\n\r\n');
        if (headerEndIndex !== -1) {
            const rawHeaders = requestData.slice(requestLine.length + 2, headerEndIndex).split('\r\n');
            rawHeaders.forEach(line => {
                const [key, value] = line.split(': ');
                headers[key] = value;
            });

            body = requestData.slice(headerEndIndex + 4);
        }

        const handler = this.routes.get(path);
        if (!handler) {
            this.sendResponse(socket, 404, "Not Found");
            return;
        }

        try {
            const response = handler(method, body);
            this.sendResponse(socket, response.statusCode, response.body);
        } catch (err) {
            console.error(`Error processing request: ${err.message}`);
            this.sendResponse(socket, 500, "Internal Server Error");
        }
    }

    sendResponse(socket, statusCode, body) {
        const statusMessage = this.getStatusMessage(statusCode);
        const response = `HTTP/1.1 ${statusCode} ${statusMessage}\r\n` +
            `Content-Length: ${Buffer.byteLength(body)}\r\n` +
            `Content-Type: text/plain\r\n\r\n` +
            body;

        socket.write(response, () => {
            socket.end();
        });
    }

    getStatusMessage(statusCode) {
        const messages = {
            200: "OK",
            201: "Created",
            400: "Bad Request",
            404: "Not Found",
            405: "Method Not Allowed",
            500: "Internal Server Error",
        };
        return messages[statusCode] || "Unknown Status";
    }
}

const framework = new FrontController();

framework.addRoute('/hello', (method, body) => {
    if (method === 'GET') {
        return {
            statusCode: 200,
            body: "Hello, World!"
        }
    } else if (method === 'POST') {
        return {
            statusCode: 201,
            body: `You posted: ${body}`
        }
    }
    return {
        statusCode: 405,
        body: "Method Not Allowed"
    }
});

framework.start(420);
