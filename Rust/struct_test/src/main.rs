struct Human {
    name: String,
    age: u32,
}

impl Human {
    fn new (name: String, age: u32) -> Human {
        Human {name: name, age: age}
    }

    fn birthday (&mut self) -> &mut Self {
        self.age += 1;
        self
    }

    fn print (&mut self) -> &mut Self {
        println!("{}, {}", self.name, self.age);
        self
    }

}

fn main() {
    let mut human = Human::new ("Ryuta Sasaki".to_string(), 25);
    human.print()
        .birthday()
        .print()
        .birthday()
        .birthday()
        .print();
}
