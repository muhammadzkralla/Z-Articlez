const startTime = new Date();

console.log(`We are now on ${process.pid}`);
console.log("Starting non-blocking operation...");
const futureResult = performNonblockingOperation();

futureResult.then((res) => {
    console.log(res);
});

const endTime = new Date();
console.log(`Time taken: ${endTime - startTime} ms`);

async function performNonblockingOperation() {
    return new Promise(async (resolve, reject) => {
        console.log(`performing non-blocking operation on ${process.pid}`);
        try {
            await sleep(3000);
        } catch (error) {
            reject(error);
        }

        resolve('non-blocking operation completed');

    });
}

async function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
