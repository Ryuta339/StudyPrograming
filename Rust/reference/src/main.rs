// Box<T> やString 値や Vec 値の内部にあるポインタは所有権を持つポインタ
//
// 所有権を持たないポインタ型 （参照、reference）がある
// 参照は、参照先の生存期間に何の影響も持たない
//
// ある値に対する参照を作ることを借用と呼ぶ。


use std::collections::HashMap;

type Table = HashMap <String, Vec<String>>;

fn show (table: Table) {
    for (artist, works) in table { // table の所有権は for ループへ
        println!("work by {}:", artist);
        for work in works {
            println!("\t{}", work);
        }
    }
}

fn show2 (table: &Table) { // 共有参照で引数をとる
    for (artist, works) in table {
        println!("work by {}:", artist);
        for work in works {
            println!("\t{}", work);
        }
    }
}

fn sort_works (table: &mut Table) { // ソートするためには可変参照でないといけない
    for (_artist, works) in table {
        works.sort ();
    }
}

fn main() {
    let mut table = Table::new ();
    table.insert ("Gesualdo".to_string (),
                  vec!["many madrigals".to_string (),
                       "Tenebrae Responsoria".to_string ()]);
    table.insert ("Caravaggio".to_string (),
                  vec!["The Musicians".to_string (),
                       "The Calling of St. Matthew".to_string ()]);
    table.insert ("Cellini".to_string (),
                  vec!["Perseus with the head of <edusa".to_string (),
                       "a salt cellar".to_string ()]);

    show2 (&table);  // 共有参照なので所有権は移らない
    sort_works (&mut table);
    show (table); // table の所有権はshow へ




    let x = 10;
    let y = 10;

    let rx = &x;
    let ry = &y;

    let rrx = &rx;
    let rry = &ry;

    assert! (rrx == rry);       // 10 == 10
    assert! (!std::ptr::eq(rx,ry));   // アドレスは別
}

