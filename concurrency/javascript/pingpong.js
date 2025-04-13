async function ping() {
    while (true) {
        console.log(`ping from ${process.pid}`);
        await sleep(1000);
    }
}

async function pong() {
    while (true) {
        console.log(`pong from ${process.pid}`);
        await sleep(1000);
    }
}

function main() {
    ping();
    pong();
}

async function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

main();
