const startTime = new Date();

console.log("Starting blocking operation...");
const reslut = performBlockingOperation();
console.log(`blocking operation result: ${reslut}`);

const endTime = new Date();
console.log(`Time taken: ${endTime - startTime} ms`);

function performBlockingOperation() {
    const start = Date.now();
    while (Date.now() - start < 3000) {
        // Busy-wait for 3 seconds
    }

    return "Operation completed";
}
