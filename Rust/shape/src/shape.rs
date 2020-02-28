use std::f64::consts::PI;

pub trait Shape {
    fn area2 (&self) -> f64;
}

pub struct Point {
    x: f64,
    y: f64,
}

impl Point {
    pub fn new (x: f64, y: f64) -> Point {
        Point { x:x, y:y }
    }
}

pub struct Circle {
    center: Point,
    radius: f64,
}

impl Circle {
    pub fn new (center: Point, radius: f64) -> Circle {
        Circle { center:center, radius:radius }
    }
}

impl Shape for Circle {
    fn area2 (&self) -> f64 {
        PI * self.radius * self.radius
    }
}

pub struct Rectangle {
    upper_left: Point,
    width: f64,
    height: f64,
}

impl Rectangle {
    pub fn new (upper_left: Point, width: f64, height: f64) -> Rectangle {
        Rectangle { upper_left:upper_left, width:width, height:height }
    }
}

impl Shape for Rectangle {
    fn area2 (&self) -> f64 {
        self.width * self.height
    }
}
