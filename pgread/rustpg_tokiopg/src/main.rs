use dotenv::dotenv;
use std::{env, sync::Arc};
use serde::{Deserialize, Serialize};
use axum::{extract::{Json, State}, routing::get, Router};
use tokio_postgres::{NoTls, Client};

#[derive(Debug, Serialize, Deserialize, Clone)]
struct Todo {
    title: Option<String>,
}

//#[tokio::main(flavor = "multi_thread")]
#[tokio::main()]
async fn main() {
    dotenv().ok();
    let database_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set in the .env file");

    let (client, connection) = tokio_postgres::connect(&database_url, NoTls)
        .await
        .expect("Failed to connect to the database");

    tokio::spawn(async move {
        if let Err(e) = connection.await {
            eprintln!("Database connection error: {}", e);
        }
    });

    let app_state = Arc::new(client);

    let app = Router::new()
        .route("/api/todo", get(get_todos))
        .with_state(app_state);

    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

async fn get_todos(State(client): State<Arc<Client>>) -> Json<Vec<Todo>> {
    let rows = client
        .query("SELECT title FROM todo", &[])
        .await
        .expect("Failed to execute query");

    let todos: Vec<Todo> = rows
        .iter()
        .map(|row| Todo {
            title: row.get("title"),
        })
        .collect();

    Json(todos)
}
