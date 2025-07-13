use crate::user;

struct User {
    name: String,
}

impl User {
    fn new(name: String) -> Self {
        User { name }
    }
}

pub fn test_user() {
    let user1 = User::new("User1".to_owned());
    let user2 = user1;
    print_username(user1);
}

fn print_username(user3: User) {
    println!("Username is {}", user3.name)
}
