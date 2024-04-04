use std::process::exit;

use clap::{Parser, Subcommand};
use colored::Colorize;
use nix::unistd::Uid;

use crate::{core::{modes::{breathing_mode, morph_mode, rainbow_mode, static_mode}, DEVICE_LIST}, daemon};

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

    /// Run lianlinux as a daemon
    Daemon {}
}

async fn light(mode: String, red: u8, green: u8, blue: u8) {
    if Uid::effective().is_root() { 
        match mode.as_str() {
            "static" => {
                let devices_list = DEVICE_LIST.lock().unwrap();
                static_mode(&[red, green, blue], &devices_list[0]);
            },
            "breathing" => {
                let devices_list = DEVICE_LIST.lock().unwrap();
                breathing_mode(&[red, green, blue], &devices_list[0]);
            },
            "rainbow" => {
                let devices_list = DEVICE_LIST.lock().unwrap();
                rainbow_mode(&devices_list[0]);
            },
            "morph" => {
                let devices_list = DEVICE_LIST.lock().unwrap();
                morph_mode(&devices_list[0]);
            },
            _ => {
                println!("Unknown mode {}", mode.red());
                exit(1);
            }
        }
    } else {
        match mode.as_str() {
            "static" => {
                daemon::client::static_mode(red, blue, green).await;
            },
            "breathing" => {
                daemon::client::breathing_mode(red, blue, green).await;
            },
            "rainbow" => {
                daemon::client::rainbow_mode(red, blue, green).await;
            },
            "morph" => {
                daemon::client::morph_mode(red, blue, green).await;
            },
            _ => {
                println!("Unknown mode {}", mode.red());
                exit(1);
            }
        }
    }  
    println!("Changing mode to {} with colors {red}, {green}, {blue}", mode.purple());
} 

pub async fn handle_args() {
    let args = App::parse();

    match args.command {
        Command::Light { mode, red, green, blue } => {
            light(mode, red, green, blue).await;
        },
        Command::Daemon {  } => {
            daemon::run().await;
        }
    }
}
