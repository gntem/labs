import asyncio
from typing import AsyncIterator, Callable, TypeVar, Iterable, Union
from collections.abc import Awaitable

T = TypeVar("T")
U = TypeVar("U")


class AsyncStream:
    """A utility class for async stream operations similar to Node.js Readable.map()"""

    @staticmethod
    async def from_iterable(items: Iterable[T]) -> AsyncIterator[T]:
        """Create an async stream from an iterable"""
        for item in items:
            yield item

    @staticmethod
    async def map(
        stream: AsyncIterator[T],
        mapper: Callable[[T], Union[U, Awaitable[U]]],
        concurrency: int = 1,
    ) -> AsyncIterator[U]:
        """
        Map over an async stream with optional concurrency control.

        Args:
            stream: The input async iterator
            mapper: A sync or async function to apply to each item
            concurrency: Maximum number of concurrent operations (for async mappers)
        """
        if concurrency == 1:
            async for item in stream:
                result = mapper(item)
                if asyncio.iscoroutine(result):
                    yield await result  # type: ignore
                else:
                    yield result  # type: ignore
        else:
            semaphore = asyncio.Semaphore(concurrency)
            tasks = []

            async def process_item(item: T, index: int) -> tuple[int, U]:
                async with semaphore:
                    result = mapper(item)
                    if asyncio.iscoroutine(result):
                        return (index, await result)  # type: ignore
                    return (index, result)  # type: ignore

            index = 0
            async for item in stream:
                tasks.append(asyncio.create_task(process_item(item, index)))
                index += 1

            results = await asyncio.gather(*tasks)
            results.sort(key=lambda x: x[0])
            for _, value in results:
                yield value

    @staticmethod
    async def for_each(
        stream: AsyncIterator[T],
        func: Callable[[T], Union[None, Awaitable[None]]],
        concurrency: int = 1,
    ) -> None:
        """
        Apply a function to each item in the async stream with optional concurrency.

        Args:
            stream: The input async iterator
            func: A sync or async function to apply to each item
            concurrency: Maximum number of concurrent operations (for async functions)
        """
        async for _ in AsyncStream.map(stream, func, concurrency):
            pass

    @staticmethod
    async def to_list(stream: AsyncIterator[T]) -> list[T]:
        """Collect all items from the async stream into a list"""
        result: list[T] = []
        async for item in stream:
            result.append(item)
        return result
