import asyncio
from datetime import datetime
import threading


class AsyncCoroutine:
    async def start(self):
        start_time = datetime.now()

        while True:
            print(f"Rendering UI on... {threading.current_thread().name}")
            now = datetime.now()
            elapsed_time = (now - start_time).total_seconds()

            if 5 <= elapsed_time <= 6:
                asyncio.create_task(self.doRequest())

            await asyncio.sleep(1)

    async def doRequest(self):
        print(f"Initiated the request.... on {threading.current_thread().name}")
        await asyncio.sleep(3)
        print(f"Finished the request.... on {threading.current_thread().name}")

