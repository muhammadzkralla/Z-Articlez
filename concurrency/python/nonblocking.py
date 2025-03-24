from datetime import datetime
import time
import threading


class NonBlocking:
    def start(self):
        startTime = datetime.now()

        while True:
            print(f"Rendring UI on... ${threading.current_thread().name}")
            now = datetime.now()
            elapsed_time = (now - startTime).total_seconds()

            if elapsed_time >= 5 and elapsed_time <= 6:
                thread = threading.Thread(target=self.doRequest)
                thread.start()

            time.sleep(1)

    def doRequest(self):
        print(f"Initiated the request.... on  {threading.current_thread().name}")
        time.sleep(3)
        print(f"Finished the request.... on {threading.current_thread().name}")
