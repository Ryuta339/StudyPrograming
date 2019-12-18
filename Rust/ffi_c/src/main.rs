extern crate libc;

extern {
//    fn hello ();
//    fn double_input (input: libc::c_double) -> libc::c_double;
//    fn matcreat () -> libc::c_int;
    fn doublesave (filename: *const libc::c_uchar,
                   arry: *const libc::c_double,
                   row: libc::size_t,
                   col: libc::size_t)
        -> libc::c_int;
}

fn main() {
//    unsafe { hello () };
//    println! ("Hello World! [from Rust]");
//
//
//    let input = 4.0;
//    let output = unsafe { double_input (input) };
//    println! ("{} * 2 = {}", input, output);

    let arry: [f64; 12] = [1., 2., 3., 4., 5., 6., 7., 8., 9., 10., 11., 12.];

    unsafe { doublesave ("db.mat".as_bytes ().as_ptr(), arry.as_ptr(), 3, 4) };
}
