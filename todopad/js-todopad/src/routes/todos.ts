import { Hono } from 'hono';
import { Database } from 'bun:sqlite';

import { Layout } from '../components/Layout';
import { TodoList } from '../components/TodoList';
import { TodoInput } from '../components/TodoInput';
import { TodoItem } from '../components/TodoItem';

import { Todo } from '../models/todo';

const app = new Hono();
const db = new Database('todopad.db');
/*const query = db.query(
  `insert into todos (title) values ('todo-1'), ('todo-2'), ('todo-3')`
);
query.run();*/

app.get('/', (c) => {
  const todos = db.query(`select * from todos`).as(Todo).all();
  return c.html(Layout('Todos', TodoInput() + TodoList(todos)));
});

app.post('/todos/add', async (c) => {
  const body = await c.req.parseBody();
  const title = body['title']?.toString();
  if (title) {
    const query = db.query(`insert into todos (title) values ($title)`);
    const result = query.run({ $title: title });
    return c.html(TodoItem(Number(result.lastInsertRowid), title, false));
  }
});

app.post('/todos/toggle/:id', (c) => {
  const id = Number(c.req.param('id'));
  const todo = db.run(
    `UPDATE todos SET completed = !completed WHERE id = ?`,
    id
  );
  return c.html(TodoItem(todo!.id, todo!.title, todo!.completed));
});

app.delete('/todos/delete/:id', (c) => {
  const id = Number(c.req.param('id'));
  const index = todos.findIndex((t) => t.id === id);
  if (index !== -1) {
    todos.splice(index, 1);
  }
  return c.html(TodoList(todos));
});

export default app;
