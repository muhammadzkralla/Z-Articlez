async function counter() {
    for (let i = 0; i < 10; ++i) {
        console.log(`${i} from ${process.pid}`);
        await sleep(0); // acts like yield
    }
}

function main() {
    counter();
    counter();
}

async function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

main();
