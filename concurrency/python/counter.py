import asyncio
import threading


async def counter():
    for i in range(10):
        print(f"{i} on {threading.current_thread().name}")
        await asyncio.sleep(0)  # acts like yield


async def main():
    await asyncio.gather(counter(), counter())


if __name__ == "__main__":
    asyncio.run(main())
