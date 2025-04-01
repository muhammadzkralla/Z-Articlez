function nonBlocking() {
    const startTime = Date.now();
    let count = 0;
    setInterval(async () => {
        count++;
        console.log(`Rendring UI on... ${process.pid}`);
        const currentTime = Date.now();

        if (currentTime - startTime >= 5000 && currentTime - startTime < 6000) {
            await doRequest();
        }

        console.log(`finished cb ${count}`)
    }, 1000);
}

async function doRequest() {
    console.log(`Initiated the request.... on  ${process.pid}`);
    await sleep(3000);
    console.log(`Finished the request.... on ${process.pid}`);
}

async function sleep(ms) {
    return new Promise((resolve) => { setTimeout(resolve, ms) });
}

nonBlocking();
