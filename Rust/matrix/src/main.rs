const N: usize = 100;

fn matadd (a: [[f64;N];N], b: [[f64;N];N]) -> [[f64;N];N] {
    let mut retval: [[f64;N];N] = [[0.;N];N];
    for col in 0..N {
        for row in 0..N {
            retval[col][row] = a[col][row] + b[col][row];
        }
    }
    retval

}

fn matmul (a: [[f64;N];N], b: [[f64;N];N]) -> [[f64;N];N] {
    let mut retval: [[f64;N];N] = [[0.;N];N];
    for col in 0..N {
        for row in 0..N {
            for k in 0..N {
                retval[col][row] += a[col][k] * b[k][row];
            }
        }
    }
    retval
}


fn linear_regression (x: [f64;N], y: [f64;N]) -> (f64, f64) {
    let mean_x: f64 = x.iter().fold(0., |sum, i| sum + i)/N as f64;
    let mean_y: f64 = y.iter().fold(0., |sum, i| sum + i)/N as f64;
    let var_x: f64 = x.iter().fold(0., |sum, i| sum+i*i)/N as f64 - mean_x*mean_x;
    
    let cv = x.iter().zip (y.iter ());
    let cov: f64 = cv.fold(0., |sum, i| sum+i.0*i.1)/N as f64 - mean_x*mean_y;
    let b = cov/var_x;
    (b, mean_y-b*mean_x)
}

extern crate rand;

use rand::prelude::*;
use rand::Rng;
use rand::distributions::Uniform;
use rand::distributions::Normal;


fn main() {
    let mut rng = rand::thread_rng ();
    let side = Uniform::new (-10.0, 10.0);
    let nsd = Normal::new (0.0,0.5);
    let mut a: [f64;N] = [0.;N];
    let mut b: [f64;N] = [0.;N];
    let beta = 2.0;
    let beta0 = 1.0;
    for row in 0..N {
        a[row] = rng.sample (side);
        b[row] = beta*a[row] + beta0 + rng.sample(nsd);
    }
    let c = linear_regression (a,b);
//    for col in 0..N {
//        println!("{:?}\t{:?}", a[col], b[col]);
//    }
    println!("{:}\t{:}",c.0, c.1);
}
