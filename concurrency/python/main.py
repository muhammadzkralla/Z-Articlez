import asyncio
from async_coroutine import AsyncCoroutine


async def main():
    async_coroutine = AsyncCoroutine()
    await async_coroutine.start()


if __name__ == "__main__":
    asyncio.run(main())
