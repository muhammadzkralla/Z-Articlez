class Blocking {

    start() {
        const startTime = new Date();

        while (true) {
            console.log(`Rendring UI on... ${process.pid}`);
            const now = new Date();

            if (now - startTime >= 5000 && now - startTime <= 6000) {
                this.doRequest();
            }

            this.sleep(1000);
        }
    }

    doRequest() {
        console.log(`Initiated the request.... on  ${process.pid}`);
        this.sleep(3000);
        console.log(`Finished the request.... on ${process.pid}`);
    }

    sleep(ms) {
        var start = new Date().getTime(), expire = start + ms;
        while (new Date().getTime() < expire) { }
        return;
    }
}

export { Blocking };
