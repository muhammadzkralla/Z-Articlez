use crate::user;

struct User {
    name: String,
}

impl User {
    fn new(name: String) -> Self {
        User { name }
    }
}

// Example 1 -----------
pub fn test_user1() {
    let user1 = User::new("User1".to_owned());
    let user2 = user1; // ownership moved from user1 to user2, user1 is not valid anymore
    print_username1(user2); // ownership moved from user2 to user3, user2 is not valid anymore
}

fn print_username1(user3: User) {
    println!("Username is {}", user3.name)
}
// --------------------

// Example 2 ----------
pub fn test_user2() {
    let user1 = User::new("User1".to_owned());
    let user2 = user1; // ownership moved from user1 to user2, user1 is not valid anymore
    //

    // Here, we pass immutable borrows of the user2 object to read
    // This avoids moving the ownership from user2 to user3 and allows us to use user2 again later,
    // unlike in example 1
    print_username2(&user2);
    print_username2(&user2);
}

fn print_username2(user3: &User) {
    println!("Username is {}", user3.name)
}
// -------------------

// Example 3 ----------
pub fn test_user3() {
    let user1 = User::new("User1".to_owned());
    let mut user2 = user1;

    // Here, we pass immutable borrows of the user2 object to read
    // And we pass mutable borrow of the user2 object to write
    print_username3(&user2);
    update_username3(&mut user2, "User2".to_owned());
    print_username3(&user2);
}

fn print_username3(user3: &User) {
    println!("Username is {}", user3.name)
}

fn update_username3(user4: &mut User, name: String) {
    user4.name = name;
}
// -------------------
