use cmd::handle_args;
use nix::unistd::Uid;

extern crate hidapi;

mod core;
mod cmd;
pub mod daemon;

/// # Main `lianlinux` function
///
/// Initializes `core` module if ran as root and handles 
/// command line arguments asynchrously
#[tokio::main]
async fn main() {
    if Uid::effective().is_root() {
        core::init();
    }
    handle_args().await;
}
