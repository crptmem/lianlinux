use hidapi::HidDevice;

pub const REPORT_BYTE: u8 = 0xE0;

/// # Set preset controller mode
/// 
/// The lian li controller typically have 4 ports.
/// We need to iterate through them and set their modes.
/// Preset modes is stored in the controller itself.
pub fn set_mode(color: &[u8], device: HidDevice) {
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
pub fn set_rgb_mode(color: &[u8], mode: u8, device: HidDevice) {
    for i in 0u8..4u8 { 
        let mode_packet: Vec<u8> = vec![REPORT_BYTE, 0x10 + i, mode];
        let mut rgb: Vec<u8> = vec![];

        rgb.append(&mut color.repeat(38));

        let mut packet: Vec<u8> = vec![REPORT_BYTE, 0x30 + i];
        packet.append(&mut rgb);
        let _ = device.write(&mode_packet[..]);
        let _ = device.write(&packet[..]);
    }
}
