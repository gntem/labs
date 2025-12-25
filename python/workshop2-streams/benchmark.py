"""
Simple benchmark showing the power of concurrent async sleep operations.
"""

import asyncio
import time
from async_stream import AsyncStream


async def sleep_sequential(n: int = 100, delay: float = 0.1):
    """Process items sequentially with sleep"""
    print(f"Sequential: Processing {n} items with {delay}s sleep each...")

    async def slow_task(x: int) -> int:
        await asyncio.sleep(delay)
        return x * 2

    start = time.perf_counter()
    stream = AsyncStream.from_iterable(range(n))
    result: list[int] = []
    value: int
    async for value in AsyncStream.map(stream, slow_task, concurrency=1):  # type: ignore[arg-type]
        result.append(value)
    elapsed = time.perf_counter() - start

    print(f"  Time: {elapsed:.2f}s")
    print(f"  Result: {result[:5]}")
    return elapsed


async def sleep_concurrent(n: int = 100, delay: float = 0.1, concurrency: int = 10):
    """Process items concurrently with sleep"""
    print(
        f"\nConcurrent (concurrency={concurrency}): Processing {n} items with {delay}s sleep each..."
    )

    async def slow_task(x: int) -> int:
        await asyncio.sleep(delay)
        return x * 2

    start = time.perf_counter()
    stream = AsyncStream.from_iterable(range(n))
    result: list[int] = []
    value: int
    async for value in AsyncStream.map(stream, slow_task, concurrency=concurrency):  # type: ignore[arg-type]
        result.append(value)
    elapsed = time.perf_counter() - start

    print(f"  Time: {elapsed:.2f}s")
    print(f"  Result: {result[:5]}")
    return elapsed


async def main():
    """Run benchmark"""
    print("=" * 60)
    print("ASYNC STREAM CONCURRENCY BENCHMARK")
    print("=" * 60)
    print("\nScenario: 100 items, each with 0.1s sleep (simulating I/O)\n")

    seq_time = await sleep_sequential(100, 0.1)
    conc_time = await sleep_concurrent(100, 0.1, concurrency=10)

    print(f"\n{'=' * 60}")
    print(f"Sequential time:  {seq_time:.2f}s")
    print(f"Concurrent time:  {conc_time:.2f}s")
    print(f"Speedup:          {seq_time / conc_time:.2f}x faster")
    print("=" * 60)


if __name__ == "__main__":
    asyncio.run(main())
