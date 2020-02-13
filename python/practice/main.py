import asyncio


async def main():
    await asyncio.sleep(3)
    print(42)


asyncio.run(main())
