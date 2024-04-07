use axum::{
    http::StatusCode, routing::{get, post}, Json, Router
};
use serde::{Deserialize, Serialize};

use crate::core::{modes::{breathing_mode, morph_mode, rainbow_mode, static_mode}, DEVICE_LIST};

pub mod client;

/// # Structure containing client request to change lightning mode
#[derive(Deserialize, Serialize, Debug)]
pub struct LightMethod {
    pub mode: String,
    pub red: u8,
    pub blue: u8,
    pub green: u8
}

/// # Structure containing server response
#[derive(Serialize, Debug)]
struct Response {
    status: isize,
    message: String
}

/// # Run daemon
///
/// This function listens local HTTP server on port 8471
pub async fn run() {
    let app = Router::new()
        .route("/light", post(light))
        .route("/", get(root));

    let listener = tokio::net::TcpListener::bind("127.0.0.1:8471").await.unwrap();
    println!("Listening on 127.0.0.1:8471");
    axum::serve(listener, app).await.unwrap();
}

/// # Handle lightning mode
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
        "rainbow" => {
            rainbow_mode(&DEVICE_LIST.lock().unwrap()[0]);
             (StatusCode::OK, Json(Response {
                status: 200,
                message: "ok".to_string()
            }))
        },
        "morph" => {
            morph_mode(&DEVICE_LIST.lock().unwrap()[0]);
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

/// # Root route method
///
/// Just returns "Hello, World!"
async fn root() -> &'static str {
    "Hello, World!"
}
