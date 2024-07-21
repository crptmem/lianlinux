use std::process::exit;
use crate::daemon::{self};
use clap::{Parser, Subcommand};

/// # App struct for clap
#[derive(Debug, Parser)]
#[clap(name = "my-app", version)]
pub struct App {
    #[clap(subcommand)]
    command: Command,
}

struct LightParams {
    mode: String,
    red: Option<u8>,
    green: Option<u8>,
    blue: Option<u8>,
    red2: Option<u8>,
    green2: Option<u8>,
    blue2: Option<u8>
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
        blue: Option<u8>,

        /// Second red
        red2: Option<u8>,

        /// Second green
        green2: Option<u8>,

        /// Second blue
        blue2: Option<u8>
    },

    /// Run lianlinux as a daemon
    Daemon {}
}

/// # Change lightning mode by sending requests to daemon
async fn light(params: LightParams) {
    if (params.mode == "static" || params.mode == "breathing") && (params.red.is_none() || params.green.is_none() || params.blue.is_none()) {
        println!("Not enough arguments");
        exit(1);
    }
    match params.mode.as_str() {
        "static" => {
            daemon::client::static_mode(params.red.unwrap(), params.green.unwrap(), params.blue.unwrap()).await;
        },
        "breathing" => {
            daemon::client::breathing_mode(params.red.unwrap(), params.green.unwrap(), params.green.unwrap()).await;
        },
        "rainbow" => {
            daemon::client::rainbow_mode(0, 0, 0).await;
        },
        "morph" => {
            daemon::client::morph_mode(0, 0, 0).await;
        },
        "runway" => {
            daemon::client::runway_mode(params.red.unwrap(), params.green.unwrap(), params.blue.unwrap(), params.red2.unwrap(), params.green2.unwrap(), params.blue2.unwrap()).await;
        },
        "tide" => {
            daemon::client::tide_mode(params.red.unwrap(), params.green.unwrap(), params.blue.unwrap(), params.red2.unwrap(), params.green2.unwrap(), params.blue2.unwrap()).await;
        },
        _ => {
            println!("Unknown mode {}", params.mode);
            exit(1);
        }
    }
    //println!("Changing mode to {} with colors {red}, {green}, {blue}", mode.purple());
} 

/// # Handle command line arguments
pub async fn handle_args() {
    let args = App::parse();

    match args.command {
        Command::Light { mode, red, green, blue, red2, green2, blue2 } => {
            daemon::client::init().await;
            light(LightParams {
                mode, red, green, blue, red2, green2, blue2
            }).await;
        },
        Command::Daemon {  } => {
            let _ = daemon::run().await;
        }
    }
}
