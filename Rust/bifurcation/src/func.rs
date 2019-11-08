#![crate_type = "cdylib"]

#[no_mangle]
pub extern fn func () {
    println!("Hello, world! [from Rust]");
}
