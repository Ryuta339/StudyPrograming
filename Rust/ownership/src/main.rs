//fn main() {
//    struct Person { name: String, birth: i32 }
//
//    let mut composers = Vec::new ();
//    composers.push (Person { name: "Palestrina".to_string(),
//                             birth: 1525});
//    composers.push (Person { name: "Dowland".to_string(),
//                             birth: 1563});
//    composers.push (Person { name: "Lully".to_string(),
//                             birth: 1632});
//    
//    // for composer in composers { // これだと所有権がdrop してしまう
//    for composer in &composers {
//        println!("{}, born {}", composer.name, composer.birth);
//    }
//    for composer in &composers {
//        println!("{}, born {}", composer.name, composer.birth);
//    }
//
//    let mut composers2 = &composers;
//    for composer in composers2 {
//        println!("{}, born {}", composer.name, composer.birth);
//    }
//    for composer in &composers {
//        println!("{}, born {}", composer.name, composer.birth);
//    }
//    let a = 3.14;
//    let b = a;
//    println!("{}", a);
//}


use std::rc::Rc;

fn main () {
    // Rust can infer all these types; written out for clarity
    let s: Rc <String> = Rc::new ("shirataki".to_string());
    let t: Rc <String> = s.clone ();
    let u: Rc <String> = s.clone ();
    println!("{}", &s);
    println!("{}", &t);
    println!("{}", &u);
}
