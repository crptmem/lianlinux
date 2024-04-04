use cmd::handle_args;

extern crate hidapi;

mod core;
mod cmd;

fn main() {
    core::init();
    handle_args();
}
