export const TodoInput = () => `
  <form hx-post="/todos/add" hx-target="#todo-list" class="mb-4 flex">
    <input type="text" name="title" required 
           class="flex-1 border p-2 rounded-l"
           placeholder="New Todo...">
    <button type="submit" class="bg-blue-600 text-white px-4 rounded-r">
      Add
    </button>
  </form>
`;
