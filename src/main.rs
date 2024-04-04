use cmd::handle_args;
use nix::unistd::Uid;

extern crate hidapi;

mod core;
mod cmd;
pub mod daemon;

#[tokio::main]
async fn main() {
    if Uid::effective().is_root() {
        core::init();
    } else {
        println!("Connecting to daemon...");
        daemon::client::init().await;
    }  
    handle_args().await;
}
