//
// 参考
// https://www.mathgram.xyz/entry/rust/gd
//

extern crate gnuplot;
extern crate nalgebra;

use std::f64;
use std::vec::Vec;

use gnuplot::{Figure, Color};


struct NormalDistribution {
    mu: f64,
    sigma: f64,
}

impl NormalDistribution {
    fn new (mu: f64, sigma: f64) -> NormalDistribution {
        NormalDistribution { mu: mu, sigma: sigma }
    }

    fn calc (&mut self, x: &Vec<f64>) -> Vec<f64> {
        let mut y: Vec<f64> = vec![];
        for e in x {
            let a = - (e - self.mu).powi(2) / (2.0f64 * self.sigma.powi (2));
            let b = (2.0_f64 * f64::consts::PI).sqrt () * self.sigma;
            y.push (a.exp () / b);
        }

        y
    }

    fn show_info (&mut self) {
        println!("------- Information -------");
        println!("Average: {}, Variance: {}", self.mu, self.sigma);
        println!("---------------------------");
    }

    fn plot (&mut self, x: &Vec<f64>) {
        let mut fg = Figure::new ();
        let y: Vec<f64> = self.calc (&x);

        fg.axes2d().lines (x, &y, &[Color("blue")]);

        fg.set_terminal ("png", "norm.png");
        fg.show ();
    }
}

fn main() {
    let mut norm = NormalDistribution::new (0.0f64, 1.0f64);
    let mut x: Vec<f64> = vec![];
    let mut n = -3.0f64;

    for _ in 0..60 {
        x.push (n);
        n += 0.1;
    }

    norm.show_info ();

    // plot
    norm.plot (&x);
}

