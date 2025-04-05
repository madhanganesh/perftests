import sqlite3
from datetime import datetime

from flask import Flask, render_template, request, redirect, url_for, g

app = Flask(__name__)
DATABASE = 'todos.db'


def get_db():
    db = getattr(g, '_database', None)
    if db is None:
        db = g._database = sqlite3.connect(DATABASE)
        db.row_factory = sqlite3.Row
    return db


@app.teardown_appcontext
def close_connection(_exception):
    db = getattr(g, '_database', None)
    if db is not None:
        db.close()


def init_db():
    with app.app_context():
        db = get_db()
        db.execute("""CREATE TABLE IF NOT EXISTS todos (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        title TEXT NOT NULL,
                        created_at TEXT NOT NULL)""")
        db.commit()


@app.route('/')
def index():
    return render_template('about.html')


@app.route('/todos', methods=['GET', 'POST'])
def todos():
    db = get_db()
    if request.method == 'POST':
        title = request.form.get('title')
        if title:
            db.execute(
                'INSERT INTO todos (title, created_at) VALUES (?, ?)',
                (title, datetime.utcnow().isoformat()),
            )
            db.commit()
        return redirect(url_for('todos'))

    cursor = db.execute('SELECT * FROM todos ORDER BY created_at DESC')
    todos = cursor.fetchall()
    return render_template('todos.html', todos=todos)


if __name__ == '__main__':
    init_db()
    app.run(debug=True)
