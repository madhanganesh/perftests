from quart import Quart, render_template, request, redirect, url_for, g
import aiosqlite
from datetime import datetime
import os

app = Quart(__name__)
DATABASE = 'todos.db'


async def get_db():
    if not hasattr(g, 'db'):
        g.db = await aiosqlite.connect(DATABASE)
        g.db.row_factory = aiosqlite.Row
    return g.db


@app.before_serving
async def init_db():
    db = await get_db()
    await db.execute("""CREATE TABLE IF NOT EXISTS todos (
                          id INTEGER PRIMARY KEY AUTOINCREMENT,
                          title TEXT NOT NULL,
                          created_at TEXT NOT NULL)""")
    await db.commit()


@app.after_serving
async def close_db():
    db = getattr(g, 'db', None)
    if db:
        await db.close()


@app.route('/')
async def index():
    return await render_template('about.html')


@app.route('/todos', methods=['GET', 'POST'])
async def todos():
    db = await get_db()
    if request.method == 'POST':
        form = await request.form
        title = form.get('title')
        if title:
            await db.execute(
                'INSERT INTO todos (title, created_at) VALUES (?, ?)',
                (title, datetime.utcnow().isoformat()),
            )
            await db.commit()
        return redirect(url_for('todos'))

    cursor = await db.execute('SELECT * FROM todos ORDER BY created_at DESC')
    todos = await cursor.fetchall()
    return await render_template('todos.html', todos=todos)


if __name__ == '__main__':
    import uvicorn

    uvicorn.run('app:app', host='0.0.0.0', port=8001, reload=True)
