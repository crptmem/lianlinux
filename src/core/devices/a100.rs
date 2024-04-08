use std::{process::exit, vec};

use colored::Colorize;
use hidapi::{HidApi, HidDevice};

use crate::core::{DEVICE_LIST, LIANLI_VENDOR_ID};

extern crate hidapi;

const PRODUCT_ID: u16 = 0xa100;

pub const REPORT_BYTE: u8 = 0xE0;

/// # Set preset controller mode
/// 
/// The lian li controller typically have 4 ports.
/// We need to iterate through them and set their modes.
/// Preset modes is stored in the controller itself.
pub fn set_mode(color: &[u8], device: &HidDevice) {
    for i in 0u8..4u8 {
        let mut packet: Vec<u8> = vec![REPORT_BYTE, 0x10 + i];

        packet.append(&mut color.to_vec());

        let _ = device.write(&packet[..]);
    }
}

/// # Set custom colors controller mode
///
/// The lian li controller typically have 4 ports.
/// We need to iterate through them and set their modes.
pub fn set_rgb_mode(color: &[u8], mode: u8, device: &HidDevice) {
    for i in 0u8..4u8 { 
        let mode_packet: Vec<u8> = vec![REPORT_BYTE, 0x10 + i, mode];
        let mut rgb: Vec<u8> = vec![];

        rgb.append(&mut color.repeat(128));

        let mut packet: Vec<u8> = vec![REPORT_BYTE, 0x30 + i];
        packet.append(&mut rgb);
        let _ = device.write(&mode_packet[..]);
        let _ = device.write(&packet[..]);
    }
}


/// # Initialize controller driver
///
/// Test if a controller handles write operations ok
pub fn init(api: HidApi) {
    let mut devices_list = DEVICE_LIST.lock().unwrap();
    let device = api.open(LIANLI_VENDOR_ID, PRODUCT_ID).unwrap();
    let ret = device.write(&[0xe0, 0, 0, 0, 0]); // Test write
    println!("Found a {} controller", device.get_product_string().unwrap().unwrap().green());
    if ret.is_err() {
        println!("Write {} with error {:?}", "failed".red(), ret.err().unwrap());
        exit(1);
    }
    devices_list.push(device);
}
