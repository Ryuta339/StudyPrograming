pub mod shape;



#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        use crate::shape::Shape;
        let c =  super::shape::Point::new (0.0, 0.0);
        let circ = super::shape::Circle::new (c, 1.0);

        assert_eq! (circ.area2 (), std::f64::consts::PI);
    }
}
