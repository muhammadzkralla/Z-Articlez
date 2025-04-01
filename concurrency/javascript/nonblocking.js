class NonBlocking {

    start() {
        const startTime = new Date();

        setInterval(() => {
            console.log(`Rendring UI on... ${process.pid}`);
            const now = new Date();

            if (now - startTime >= 5000 && now - startTime <= 6000) {
                this.doRequest();
            }
        }, 1000);

        // if you uncomment this line, the set interval will never call back
        //this.block();
    }

    async doRequest() {
        console.log(`Initiated the request.... on  ${process.pid}`);
        await this.sleep(3000);
        console.log(`Finished the request.... on ${process.pid}`);
    }

    async sleep(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

    block(ms) {
        const startTime = new Date().getTime();
        while (true) {
            const currentTime = new Date().getTime();
            if (currentTime >= startTime + ms) {
                return;
            }
        }
    }

}

export { NonBlocking };
