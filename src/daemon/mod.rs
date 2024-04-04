use axum::{
    http::StatusCode, routing::{get, post}, Json, Router
};
use serde::{Deserialize, Serialize};

use crate::core::{modes::{breathing_mode, static_mode}, DEVICE_LIST};

pub mod client;

#[derive(Deserialize, Serialize, Debug)]
pub struct LightMethod {
    pub mode: String,
    pub red: u8,
    pub blue: u8,
    pub green: u8
}

#[derive(Serialize, Debug)]
struct Response {
    status: isize,
    message: String
}

pub async fn run() {
    let app = Router::new()
        .route("/light", post(light))
        .route("/", get(root));

    let listener = tokio::net::TcpListener::bind("127.0.0.1:8471").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

async fn light(Json(payload): Json<LightMethod>) -> (StatusCode, Json<Response>) {
    match payload.mode.as_str() {
        "static" => {
            static_mode(&[payload.red, payload.blue, payload.green], &DEVICE_LIST.lock().unwrap()[0]);
             (StatusCode::OK, Json(Response {
                status: 200,
                message: "ok".to_string()
            }))
        },
        "breathing" => {
            breathing_mode(&[payload.red, payload.blue, payload.green], &DEVICE_LIST.lock().unwrap()[0]);
             (StatusCode::OK, Json(Response {
                status: 200,
                message: "ok".to_string()
            }))
        },
        _ => {
            (StatusCode::BAD_REQUEST, Json(Response {
                status: 400,
                message: "invalid mode".to_string()
            }))
        }
    } 
}

async fn root() -> &'static str {
    "Hello, World!"
}
