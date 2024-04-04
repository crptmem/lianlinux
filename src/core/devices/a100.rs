use hidapi::HidApi;

use crate::core::LIANLI_VENDOR_ID;

extern crate hidapi;

const PRODUCT_ID: u16 = 0xa100;

pub fn init(api: HidApi) {
    let device = api.open(LIANLI_VENDOR_ID, PRODUCT_ID).unwrap();
    let ret = device.write(&[0xe0,0,0,0,0]);
    println!("Found a {} controller", device.get_product_string().unwrap().unwrap());
    if ret.is_err() {
        println!("Write failed with error {:?}", ret.err().unwrap());
    }
}
