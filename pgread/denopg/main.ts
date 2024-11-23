import { Application, Router } from "https://deno.land/x/oak/mod.ts";
import { Client } from "https://deno.land/x/postgres@v0.14.0/mod.ts";

const client = await connectDB();

const app = new Application();
const router = new Router();


router.get("/api/todo", async (context) => {
    const todos = await client.queryArray("select title from todo");
    context.response.body = todos;
    context.response.headers.set("Content-Type", "application/json");
});

app.use(router.routes());
app.use(router.allowedMethods());

console.log("Server running on http://localhost:8000");
await app.listen({ port: 8000 });


async function connectDB() {
  const client = new Client({
    user: "postgres",
    password: "zenith",
    database: "rants",
    hostname: "localhost", // or the hostname of your PostgreSQL server
    port: 5432,            // default PostgreSQL port
  });

  await client.connect();
  console.log("Connected to the database!");

  return client;
}
