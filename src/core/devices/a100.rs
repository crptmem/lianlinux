use colored::Colorize;
use hidapi::HidApi;

use crate::core::{packet, LIANLI_VENDOR_ID};

extern crate hidapi;

const PRODUCT_ID: u16 = 0xa100;

pub fn init(api: HidApi) {
    let device = api.open(LIANLI_VENDOR_ID, PRODUCT_ID).unwrap();
    let ret = device.write(&[0xe0, 0, 0, 0, 0]);
    println!("Found a {} controller", device.get_product_string().unwrap().unwrap().green());
    if ret.is_err() {
        println!("Write {} with error {:?}", "failed".red(), ret.err().unwrap());
    }
    packet::set_rgb_mode(&[255, 255, 255], 0x03, device);
}
