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
async fn change(red: u8, blue: u8, green: u8, red2: u8, green2: u8, blue2: u8, mode: String) {
    let client = reqwest::Client::new();

    let response = client
        .post("http://127.0.0.1:8471/light")
        .header("Content-Type", "application/json")
        .json(&LightMethod {
            mode,
            red,
            blue,
            green,
            red2,
            green2,
            blue2
        })
        .send()
        .await.unwrap();
    if response.status() != 200 {
        println!("{} to send POST request to daemon: {:?}", "Failed".red(), response.text().await);
    }
}

/// # Static lightning mode
pub async fn static_mode(red: u8, green: u8, blue: u8) {
    change(red, green, blue, 0, 0, 0, "static".to_string()).await;
}

/// # Breathing lightning mode
pub async fn breathing_mode(red: u8, green: u8, blue: u8) {
    change(red, green, blue, 0, 0, 0, "breathing".to_string()).await;
}

/// # Rainbow lightning mode
pub async fn rainbow_mode(red: u8, green: u8, blue: u8) {
    change(red, green, blue, 0, 0, 0, "rainbow".to_string()).await;
}

/// # Morph lightning mode
pub async fn morph_mode(red: u8, green: u8, blue: u8) {
    change(red, green, blue, 0, 0, 0, "morph".to_string()).await;
}

/// # Runway lightning mode
pub async fn runway_mode(red: u8, green: u8, blue: u8, red2: u8, green2: u8, blue2: u8) {
    change(red, green, blue, red2, green2, blue2, "runway".to_string()).await;
}

/// # Tide lightning mode
pub async fn tide_mode(red: u8, green: u8, blue: u8, red2: u8, green2: u8, blue2: u8) {
    change(red, green, blue, red2, green2, blue2, "tide".to_string()).await;
}
