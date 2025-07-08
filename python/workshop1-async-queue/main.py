import asyncio
import aiosqlite
import time

concurrency = 2
delay = 2


async def setup_database():
    async with aiosqlite.connect("items.db") as db:
        await db.execute("""
            CREATE TABLE IF NOT EXISTS items (
                id INTEGER PRIMARY KEY,
                name TEXT
            )
        """)
        
        await db.execute("DELETE FROM items")
        
        items_data = [
            (1, "Item One"),
            (2, "Item Two"),
            (3, "Item Three"),
            (4, "Item Four"),
            (5, "Item Five"),
            (6, "Item Six"),
            (7, "Item Seven"),
            (8, "Item Eight"),
            (9, "Item Nine"),
            (10, "Item Ten")
        ]
        
        await db.executemany("INSERT INTO items (id, name) VALUES (?, ?)", items_data)
        await db.commit()

async def load_items():
    async with aiosqlite.connect("items.db") as db:
        cursor = await db.execute("SELECT id, name FROM items")
        rows = await cursor.fetchall()
        return [{"id": row[0], "name": row[1]} for row in rows]

async def task_callback(item):
    await asyncio.sleep(delay)
    return f"Processed: {item['name']} (ID: {item['id']})"

async def main():
    await setup_database()
    
    items = await load_items()
    
    queue = asyncio.Queue()
    
    for item in items:
        await queue.put(item)
    
    results = []
    first_task_duration = None
    start_time = time.time()
    
    async def worker():
        nonlocal first_task_duration
        while not queue.empty():
            try:
                item = await queue.get()
                task_start = time.time()
                result = await task_callback(item)
                task_end = time.time()
                
                if first_task_duration is None:
                    first_task_duration = task_end - task_start
                    total_items = len(items)
                    estimated_total_time = (first_task_duration * total_items) / concurrency
                    print(f"First task completed in {first_task_duration/60:.2f} minutes")
                    print(f"Estimated total time: {estimated_total_time/60:.2f} minutes for {total_items} items with {concurrency} workers")
                
                print(result)
                results.append(result)
                queue.task_done()
            except asyncio.QueueEmpty:
                break
    
    workers = [asyncio.create_task(worker()) for _ in range(concurrency)]
    await asyncio.gather(*workers)
    
    end_time = time.time()
    actual_total_time = end_time - start_time
    print(f"\nActual total time: {actual_total_time/60:.2f} minutes")
    
    for result in results:
        print(result)

if __name__ == "__main__":
    asyncio.run(main())
