enum Message {
    Quit,
    ChangeColor(i32,i32,i32),
    Move { x: i32, y: i32},
    Write (String),
}

fn quit () { /* ... */ }
fn change_color (r: i32, g: i32, b: i32) { /* ... */ }
fn move_cursor (x: i32, y: i32) { /* ... */ }

fn process_message (msg:Message) {
    match msg {
        Message::Quit => quit (),
        Message::ChangeColor (r,g,b) => change_color (r,g,b),
        Message::Move {x: x, y: y} => move_cursor (x, y),
        Message::Write (s) => println!("{}",s),
    };
}





trait HasArea {
    fn area (&self) -> f64;
}

struct Circle {
    x: f64,
    y: f64,
    radius: f64,
}

impl HasArea for Circle {
    fn area (&self) -> f64 {
        std::f64::consts::PI * (self.radius * self.radius)
    }
}

struct Square {
    x: f64,
    y: f64,
    side: f64,
}

impl HasArea for Square {
    fn area (&self) -> f64 {
        self.side * self.side
    }
}

fn print_area <T: HasArea> (shape: T) {
    println! ("This shape has an area of {}", shape.area ());
}





trait Foo {
    fn foo (&self);
}

trait FooBar : Foo {
    fn foobar (&self);
}

struct Baz;

impl Foo for Baz{
    fn foo (&self) { println!("foo"); }
}

impl FooBar for Baz {
    fn foobar (&self) { println!("foobar"); }
}




fn main() {
//    diverges ();


    let c = Circle {
        x: 0.0f64,
        y: 0.0f64,
        radius: 1.0f64,
    };

    let s = Square {
        x: 0.0f64,
        y: 0.0f64,
        side: 1.0f64,
    };

    print_area (c);
    print_area (s);
}





fn diverges () -> ! {
    panic! ("This function never returns!");
}
