use std::vec;

use hidapi::HidDevice;

use super::devices::a100::{set_mode, set_rgb_mode};

/// # Set static mode by sending USB packet
pub fn static_mode(color: &[u8], device: &HidDevice) {
    set_rgb_mode(color, 0x03, &device);
    
    let mut packet = vec![0x01, 0x02];
    packet.append(&mut [0x00].repeat(0x250).to_vec());
    set_mode(&packet[..], &device)
}

/// # Set breathing mode by sending USB packet
pub fn breathing_mode(color: &[u8], device: &HidDevice) {
    set_rgb_mode(color, 0x03, &device);
    
    let mut packet = vec![0x02, 0xff];
    packet.append(&mut [0x00].repeat(0x500).to_vec());
    set_mode(&packet[..], &device)
}

/// # Set rainbow mode by sending USB packet
pub fn rainbow_mode(device: &HidDevice) {
    let mut packet = vec![0x05, 0xff];
    packet.append(&mut [0x00].repeat(0x250).to_vec());
    set_mode(&packet[..], &device)
}

/// # Set morph mode by sending USB packet
pub fn morph_mode(device: &HidDevice) {
    let mut packet = vec![0x04, 0xff];
    packet.append(&mut [0x00].repeat(0x250).to_vec());
    set_mode(&packet[..], &device)
}
