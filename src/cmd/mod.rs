use std::process::exit;
use crate::daemon;
use clap::{Parser, Subcommand};
use colored::Colorize;

/// # App struct for clap
#[derive(Debug, Parser)]
#[clap(name = "my-app", version)]
pub struct App {
    #[clap(subcommand)]
    command: Command,
}

/// # Enum with subcommands
#[derive(Debug, Subcommand)]
enum Command {
    /// Change controller lightning mode
    Light {
        /// Light mode
        mode: String,

        /// Red
        red: Option<u8>,

        /// Green
        green: Option<u8>,

        /// Blue
        blue: Option<u8>
    },

    /// Run lianlinux as a daemon
    Daemon {}
}

/// # Change lightning mode by sending requests to daemon
async fn light(mode: String, red: Option<u8>, green: Option<u8>, blue: Option<u8>) {
    if (mode == "static" || mode == "breathing") && (red.is_none() || green.is_none() || blue.is_none()) {
        println!("Not enough arguments");
        exit(1);
    }
    match mode.as_str() {
        "static" => {
            daemon::client::static_mode(red.unwrap(), blue.unwrap(), green.unwrap()).await;
        },
        "breathing" => {
            daemon::client::breathing_mode(red.unwrap(), blue.unwrap(), green.unwrap()).await;
        },
        "rainbow" => {
            daemon::client::rainbow_mode(0, 0, 0).await;
        },
        "morph" => {
            daemon::client::morph_mode(0, 0, 0).await;
        },
        _ => {
            println!("Unknown mode {}", mode.red());
            exit(1);
        }
    }
    //println!("Changing mode to {} with colors {red}, {green}, {blue}", mode.purple());
} 

/// # Handle command line arguments
pub async fn handle_args() {
    let args = App::parse();

    match args.command {
        Command::Light { mode, red, green, blue } => {
            daemon::client::init().await;
            light(mode, red, green, blue).await;
        },
        Command::Daemon {  } => {
            daemon::run().await;
        }
    }
}
