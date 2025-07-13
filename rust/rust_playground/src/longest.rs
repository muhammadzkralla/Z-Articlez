fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() { x } else { y }
}

pub fn test_longest() {
    let x = String::from("hello");
    let y = String::from("hell");

    let z = longest(&x, &y);

    println!("Longest is {z}");
}
