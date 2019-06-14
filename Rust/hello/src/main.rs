use std::io::Write;
use std::str::FromStr;

mod euc;        // in order to call function gcd

fn main() {
    let mut numbers = Vec::new ();          // Vec is like list in Python

    for arg in std::env::args (). skip(1) { // skipping the program name
        numbers.push (u64::from_str (&arg)
                      .expect ("error parsing argument"));
    }

    if numbers.len () == 0 {
        writeln!(std::io::stderr(), "Usage: gcd NUMBER ...").unwrap ();
        std::process::exit (1);
    }

    let mut d = numbers[0];
    for m in & numbers[1..] {
        d = euc::gcd (d, *m);
    }

    println! ("The greatest common divisor of {:?} is {}",
              numbers, d);

}
