use hidapi::HidApi;

use crate::core::devices::a100;
pub mod devices;

pub const LIANLI_VENDOR_ID: u16 = 0x0cf2;

pub fn init() {
    match HidApi::new() {
        Ok(api) => {
            for device in api.device_list() {
                if device.vendor_id() == LIANLI_VENDOR_ID {
                    match device.product_id() {
                        0xa100 => a100::init(HidApi::new().unwrap()),
                        _ => println!("Unsupported controller found")
                    }
                }
            }
        },
        Err(e) => {
            eprintln!("Error initializing HidApi: {}", e);
        },
    }
}
