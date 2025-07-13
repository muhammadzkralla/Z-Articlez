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

// Example 4 ----------
pub fn test_user4() {
    let mut user1 = User::new("User1".to_owned());

    // This will fail as we borrowed user1 as an immutable object,
    // then we borrowed it as a mutable object,
    // then we modified the mutable borrow of user1
    // then we printed the immutable borrow of user1
    // this fails as immutable borrows are expected not to be mutated until dropped

    // COMMENTED TO AVOID COMPILE TIME ERRORS
    // let r1 = &mut user1;
    // let r2 = &user1;
    // update_username3(r1, "User2".to_owned());
    // print_username3(&r2);
}

fn print_username4(user3: &User) {
    println!("Username is {}", user3.name)
}

fn update_username4(user4: &mut User, name: String) {
    user4.name = name;
}
// -------------------

// Example 5 ----------
pub fn test_user5() {
    let mut user1 = User::new("User1".to_owned());

    // We solved the issue by modifying the mutable borrow before
    // creating the immutable borrow, now the immutable borrow
    // is sure that the value of user1 will not change before it drops
    let r1 = &mut user1;
    update_username5(r1, "User2".to_owned());
    let r2 = &user1;
    print_username5(r2);
}

fn print_username5(user3: &User) {
    println!("Username is {}", user3.name)
}

fn update_username5(user4: &mut User, name: String) {
    user4.name = name;
}
// -------------------

// Example 6 ----------
pub fn test_user6() {
    let mut user1 = User::new("User1".to_owned());

    let r1 = &mut user1;
    update_username6(r1, "User2".to_owned());
    let r2 = &user1;
    print_username6(r2);

    // COMMENTED TO AVOID COMPILE TIME ERRORS
    // drop(r2);
    // update_username6(r1, "User3".to_owned());
    // let r3 = &user1;
    // print_username6(r3);
}

fn print_username6(user3: &User) {
    println!("Username is {}", user3.name)
}

fn update_username6(user4: &mut User, name: String) {
    user4.name = name;
}
// -------------------
