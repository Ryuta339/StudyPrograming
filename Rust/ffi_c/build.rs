extern crate cc;


fn main () {
    println!("cargo:rustc-link-search=native=/Applications/MATLAB_R2018a.app/bin/maci64");
    println!("cargo:rustc-link-lib=mx");
    println!("cargo:rustc-link-lib=mat");

    cc::Build::new()
        // .cpp(true)  // in the case of C++ 
        .files (vec!["src/hello.c","src/double.c","src/matcreat.c", "src/doublesave.c"])
//        .object ("/Applications/MATLAB_R2018a.app/bin/maci64/libmx.dylib")
//        .object ("/Applications/MATLAB_R2018a.app/bin/maci64/libmat.dylib")
        .include ("/Applications/MATLAB_R2018a.app/extern/include")
        .compile ("libfoo.a");
}
