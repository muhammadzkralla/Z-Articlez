import asyncio
import threading


async def counter():
    for i in range(10):
        print(f"{i} on {threading.current_thread().name}")
        await asyncio.sleep(0)  # acts like yield


async def main():
    # more manual and easy to understand than using:
    # asyncio.gather(counter(), counter())
    task1 = asyncio.create_task(counter())
    task2 = asyncio.create_task(counter())

    await task1
    await task2

    asyncio.gather


if __name__ == "__main__":
    asyncio.run(main())
