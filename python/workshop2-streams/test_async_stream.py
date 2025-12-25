import pytest
import asyncio
from async_stream import AsyncStream


@pytest.mark.asyncio
async def test_from_iterable():
    items = [1, 2, 3, 4, 5]
    result = []

    async for item in AsyncStream.from_iterable(items):
        result.append(item)

    assert result == items


@pytest.mark.asyncio
async def test_map_sync_mapper():
    stream = AsyncStream.from_iterable([1, 2, 3, 4])
    result = []

    async for value in AsyncStream.map(stream, lambda x: x * 2):  # type: ignore[arg-type]
        result.append(value)

    assert result == [2, 4, 6, 8]


@pytest.mark.asyncio
async def test_map_async_mapper():
    async def async_double(x: int) -> int:
        await asyncio.sleep(0.01)
        return x * 2

    stream = AsyncStream.from_iterable([1, 2, 3])
    result = []

    async for value in AsyncStream.map(stream, async_double):  # type: ignore[arg-type]
        result.append(value)

    assert result == [2, 4, 6]


@pytest.mark.asyncio
async def test_map_with_concurrency():
    async def async_operation(x: int) -> int:
        await asyncio.sleep(0.01)
        return x * 2

    stream = AsyncStream.from_iterable([1, 2, 3, 4])
    result = []

    async for value in AsyncStream.map(stream, async_operation, concurrency=2):  # type: ignore[arg-type]
        result.append(value)

    assert result == [2, 4, 6, 8]


@pytest.mark.asyncio
async def test_map_empty_stream():
    stream = AsyncStream.from_iterable([])
    result = []

    async for value in AsyncStream.map(stream, lambda x: x * 2):  # type: ignore[arg-type]
        result.append(value)

    assert result == []


@pytest.mark.asyncio
async def test_map_string_transformation():
    stream = AsyncStream.from_iterable(["hello", "world"])
    result = []

    async for value in AsyncStream.map(stream, str.upper):  # type: ignore[arg-type]
        result.append(value)

    assert result == ["HELLO", "WORLD"]


@pytest.mark.asyncio
async def test_map_preserves_order():
    async def reverse_delay(x: int) -> int:
        await asyncio.sleep(0.001 * (5 - x))
        return x

    items = [1, 2, 3, 4, 5]
    stream = AsyncStream.from_iterable(items)
    result = []

    async for value in AsyncStream.map(stream, reverse_delay, concurrency=5):  # type: ignore[arg-type]
        result.append(value)

    assert result == items
