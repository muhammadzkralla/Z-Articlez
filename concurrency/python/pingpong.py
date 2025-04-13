import asyncio
import threading


async def ping():
    while True:
        print(f"ping from {threading.current_thread().name}")
        await asyncio.sleep(1)


async def pong():
    while True:
        print(f"pong from {threading.current_thread().name}")
        await asyncio.sleep(1)


async def main():
    # more manual and easy to understand than using:
    # asyncio.gather(counter(), counter())
    task1 = asyncio.create_task(ping())
    task2 = asyncio.create_task(pong())

    await task1
    await task2


if __name__ == "__main__":
    asyncio.run(main())
