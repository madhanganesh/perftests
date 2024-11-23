use dotenv::dotenv;
use std::{env, sync::Arc};
use serde::{Deserialize, Serialize};
use axum::{extract::{Json, State}, routing::get, Router};
use sqlx::{postgres::PgPoolOptions, Pool, prelude::FromRow, Postgres};

#[derive(Debug, Serialize, Deserialize, Clone, FromRow)]
struct Todo {
    title: Option<String>,
}

#[tokio::main(flavor = "multi_thread")]
async fn main() {
    dotenv().ok();
    let database_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set in the .env file");

    let pool = PgPoolOptions::new()
        .max_connections(50)
        .connect(&database_url)
        .await
        .unwrap();

    let app = Router::new()
        .route("/api/todo", get(get_todos))
        .with_state(Arc::new(pool));

    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

async fn get_todos(State(pool): State<Arc<Pool<Postgres>>>) -> Json<Vec<Todo>>  {
    let todos: Vec<Todo> = sqlx::query_as!(Todo, "SELECT title FROM todo")
        .fetch_all(&*pool)
        .await
        .unwrap(); 

    Json(todos)
}
