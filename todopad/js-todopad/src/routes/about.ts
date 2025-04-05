import { Hono } from "hono";
import { Layout } from "../components/Layout";

const app = new Hono();

app.get("/", (c) => {
  return c.html(Layout("About", "<h2 class='text-lg'>This is a simple todo app using Hono, HTMX, and Tailwind.</h2>"));
});

export default app;
