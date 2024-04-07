use std::process::exit;

use colored::Colorize;

use super::LightMethod;

/// # Initialize daemon client
///
/// Checks is daemon available
pub async fn init() {
    let response = reqwest::get("http://127.0.0.1:8471")
        .await;
    if response.is_ok() {
        println!("{} to daemon successfully", "Connected".green());
    } else {
        println!("{} to connect to daemon. Is it running?", "Failed".red());
        exit(1);
    }
}

/// # Change controller light mode by calling daemon
///
/// Sends a POST JSON request to localhost daemon which
/// contains `mode`, `red`, `blue` and `green`
async fn change(red: u8, blue: u8, green: u8, mode: String) {
    let client = reqwest::Client::new();

    let response = client
        .post("http://127.0.0.1:8471/light")
        .header("Content-Type", "application/json")
        .json(&LightMethod {
            mode,
            red,
            blue,
            green
        })
        .send()
        .await.unwrap();
    if response.status() != 200 {
        println!("{} to send POST request to daemon: {:?}", "Failed".red(), response.text().await);
    }
}

/// # Static lightning mode
pub async fn static_mode(red: u8, blue: u8, green: u8) {
    change(red, blue, green, "static".to_string()).await;
}

/// # Breathing lightning mode
pub async fn breathing_mode(red: u8, blue: u8, green: u8) {
    change(red, blue, green, "breathing".to_string()).await;
}

/// # Rainbow lightning mode
pub async fn rainbow_mode(red: u8, blue: u8, green: u8) {
    change(red, blue, green, "rainbow".to_string()).await;
}

/// # Morph lightning mode
pub async fn morph_mode(red: u8, blue: u8, green: u8) {
    change(red, blue, green, "morph".to_string()).await;
}
