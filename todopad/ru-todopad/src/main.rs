use anyhow::{Context, Result};
use askama::Template;
use axum::{Extension, Router, response::Html, routing::get};
use sqlx::{
    query_as,
    sqlite::{SqliteConnectOptions, SqlitePool, SqlitePoolOptions},
};
use std::str::FromStr;
use std::thread;
use tracing::error;

#[tokio::main]
async fn main() -> Result<()> {
    tracing_subscriber::fmt::init();

    let database_url = "sqlite://todos.db";
    let pool = get_db(database_url).await?;

    let routes = Router::new().route("/", get(index)).layer(Extension(pool));

    let listener = tokio::net::TcpListener::bind("0.0.0.0:8080").await?;
    axum::serve(listener, routes).await?;

    Ok(())
}

#[derive(sqlx::FromRow)]
#[allow(dead_code)]
struct Todo {
    id: i64,
    title: String,
}

#[derive(Template)]
#[template(path = "index.html")]
struct IndexTemplate {
    todos: Vec<Todo>,
}

async fn index(Extension(pool): Extension<SqlitePool>) -> Html<String> {
    let todos = match get_todos(&pool).await {
        Ok(todos) => todos,
        Err(e) => return handle_error(e),
    };

    let result = IndexTemplate { todos }.render().context("error rendering");
    match result {
        Ok(html) => Html(html),
        Err(e) => handle_error(e),
    }
}

async fn get_todos(pool: &SqlitePool) -> Result<Vec<Todo>> {
    query_as!(Todo, "SELECT id, title FROM todos")
        .fetch_all(pool)
        .await
        .context("Failed to fetch todos")
}

fn handle_error(error: anyhow::Error) -> Html<String> {
    error!("{:?}", error);
    Html("<h1>Internal Server Error</h1><p>Something went wrong.</p>".to_string())
}

async fn get_db(url: &str) -> Result<SqlitePool, sqlx::Error> {
    let options = SqliteConnectOptions::from_str(url)?
        .create_if_missing(true)
        .foreign_keys(true);

    let pool = SqlitePoolOptions::new().connect_with(options).await?;

    sqlx::migrate!("./migrations").run(&pool).await?;

    Ok(pool)
}
