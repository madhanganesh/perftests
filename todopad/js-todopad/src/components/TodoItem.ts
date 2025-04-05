export const TodoItem = (id: number, title: string, completed: boolean) => `
  <div class="flex items-center space-x-2 p-2 border-b">
    <input type="checkbox" 
           class="h-4 w-4"
           hx-post="/todos/toggle/${id}" 
           hx-trigger="change"
           hx-target="#todo-${id}"
           ${completed ? "checked" : ""}>
    <span id="todo-${id}" class="${completed ? "line-through text-gray-500" : ""}">
      ${title}
    </span>
    <button class="ml-auto text-red-500" 
            hx-delete="/todos/delete/${id}" 
            hx-target="#todo-list">
      ✖
    </button>
  </div>
`;
