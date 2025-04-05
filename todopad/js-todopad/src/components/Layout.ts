export const Layout = (title: string, content: string) => `
  <!DOCTYPE html>
  <html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>${title}</title>
    <script src="https://unpkg.com/htmx.org@1.9.10"></script>
    <script src="https://cdn.tailwindcss.com"></script>
  </head>
  <body class="bg-gray-100 p-6">
    <div class="max-w-xl mx-auto bg-white p-6 rounded-lg shadow">
      <nav class="mb-4 flex justify-between">
        <a href="/" class="text-blue-600">Todos</a>
        <a href="/about" class="text-blue-600">About</a>
      </nav>
      ${content}
    </div>
  </body>
  </html>
`;
