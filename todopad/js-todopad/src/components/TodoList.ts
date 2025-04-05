import { TodoItem } from './TodoItem';

export const TodoList = (
  todos: { id: number; title: string; completed: boolean }[]
) => `
  <div id="todo-list">
    ${todos.map((todo) => TodoItem(todo.id, todo.title, todo.completed)).join('')}
  </div>
`;
