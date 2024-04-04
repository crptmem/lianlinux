use colored::Colorize;

use super::LightMethod;

pub async fn init() {
    let response = reqwest::get("http://127.0.0.1:8471")
        .await.unwrap()
        .status();
    if response == 200 {
        println!("{} to daemon successfully", "Connected".green());
    }
}

pub async fn static_mode(red: u8, blue: u8, green: u8) {
    let client = reqwest::Client::new();

    let response = client
        .post("http://127.0.0.1:8471/light")
        .header("Content-Type", "application/json")
        .json(&LightMethod {
            mode: "static".to_string(),
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
