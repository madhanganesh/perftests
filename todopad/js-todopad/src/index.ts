import { Hono } from 'hono';
import todosRoute from './routes/todos';
import aboutRoute from './routes/about';

const app = new Hono();

app.route('/', todosRoute);
app.route('/about', aboutRoute);

export default app;
