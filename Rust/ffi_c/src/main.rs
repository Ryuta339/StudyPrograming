extern crate libc;

extern {
    fn hello ();
    fn double_input (input: libc::c_double) -> libc::c_double;
    fn matcreat () -> libc::c_int;
}

fn main() {
    unsafe { hello () };
    println! ("Hello World! [from Rust]");


    let input = 4.0;
    let output = unsafe { double_input (input) };
    println! ("{} * 2 = {}", input, output);


    unsafe { matcreat () };
}
