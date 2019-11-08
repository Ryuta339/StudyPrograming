const N: i32 = 2;
const TMAX: i32 = 800;
const PMAX:i32 = 4;




// extern crate approx;
extern crate nalgebra as na;


// fn calc (
//     J: &na::DMatrix<f64>,     // connectivity
//     dt: f64,            // dt
//     ps: &Vec<f64>,       // pump schedule
// ) -> Vec<f64> {
//     let mut x: Vec<na::DVector<f64>> = Vec::new ();
// 
//     let init = 
//     x.push (0.);
//     for idx in 0 .. ps.len() {
//         let c = x[idx]
//         x.push ( (-(1.-ps[idx]+c*c)*c + J*
//     }
//     x
// }
// 
// 

fn main() {
    let mut J = na::DMatrix::from_iterator (2,2,[
                                    0.,1.,
                                    1.,0.
    ].iter().cloned());
    let dt = 0.1;
    let t = (TMAX as f64/ dt) as i32;
    let mut ps: Vec<f64> = Vec::new ();
    for idx in 0..t {
        ps.push (dt * idx as f64 / TMAX as f64 * PMAX as f64);
    }

    // calc (&J, dt, &ps);

    println! ("{}", J*J);
}
