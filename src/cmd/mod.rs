use std::process::exit;

use clap::{Parser, Subcommand};
use colored::Colorize;

use crate::core::{modes::{breathing_mode, static_mode}, DEVICE_LIST};

#[derive(Debug, Parser)]
#[clap(name = "my-app", version)]
pub struct App {
    #[clap(subcommand)]
    command: Command,
}

#[derive(Debug, Subcommand)]
enum Command {
    /// Change controller lightning mode
    Light {
        /// Light mode
        mode: String,

        /// Red
        red: u8,

        /// Green
        green: u8,

        /// Blue
        blue: u8
    },
}

fn light(mode: String, red: u8, green: u8, blue: u8) {
    match mode.as_str() {
        "static" => {
            let devices_list = DEVICE_LIST.lock().unwrap();
            static_mode(&[red, green, blue], &devices_list[0]);
        },
        "breathing" => {
            let devices_list = DEVICE_LIST.lock().unwrap();
            breathing_mode(&[red, green, blue], &devices_list[0]);
        },
        _ => {
            println!("Unknown mode {}", mode.red());
            exit(1);
        }
    }
    println!("Changing mode to {} with colors {red}, {green}, {blue}", mode.purple());
} 

pub fn handle_args() {
    let args = App::parse();

    match args.command {
        Command::Light { mode, red, green, blue } => {
            light(mode, red, green, blue);
        }
    }
}
