class NonBlocking {

    start() {
        const startTime = new Date();

        setInterval(async () => {
            console.log(`Rendring UI on... ${process.pid}`);
            const now = new Date();

            if (now - startTime >= 5000 && now - startTime <= 6000) {
                await this.doRequest();
            }
        }, 1000);
    }

    async doRequest() {
        console.log(`Initiated the request.... on  ${process.pid}`);
        await this.sleep(3000);
        console.log(`Finished the request.... on ${process.pid}`);
    }

    async sleep(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }
}

export { NonBlocking };
