# Notes

## longest.rs

`longest()` is a function that is valid over the lifecycle `'a` and takes two arguments, `x` which is a borrow of a string that is valid over the generic lifecycle `'a`, and `y` which is a borrow of a string that is valid over the generic lifecycle `'a`, and returns a string that is valid over the generic lifecycle `'a`.

Note that rust infers the value of the generic lifecycle `'a` automatically and we don't need to worry about this part. We just suppose that it's the shortest lifecycle of the given arguments.

## Rust's Lifetime Elision Rules

### Rule 1: Each parameter gets its own lifetime

```rust
fn foo(x: &str); // treated as fn foo<'a>(x: &'a str)
```

### Rule 2: If there’s only one input reference, the output gets its lifetime

```rust
fn identity(x: &str) -> &str; // becomes fn identity<'a>(x: &'a str) -> &'a str
```

### Rule 3: If multiple inputs, and one is &self or &mut self, use its lifetime

```rust
impl MyStruct {
    fn get_name(&self) -> &str; // becomes fn get_name<'a>(&'a self) -> &'a str
}
```

## Dropped Lifecycle Example

```rust
fn main() {
    let string1 = String::from("abc"); // string1 lives until end of main
    let result;
    {
        let string2 = String::from("abcdef"); // string2 lives only in this block
        result = longest(&string1, &string2); // ERROR!
    }
    println!("{}", result); // may be referencing dropped data
}
```

## TLDR

- Rust infers lifetimes when it can (thanks to elision rules).
- When returning references involving multiple input lifetimes, you must write them explicitly.
- When in doubt, write `'a`, the compiler will guide you.

---

## user.rs

We mocked a simple struct called `User` with only one field called `name` that's a string. We implement only one function called `new` that acts like a constructor to create a new `User` object.

## Immutable vs Mutable Borrows

### Example 1

Example one shows how ownership moves from `user1` to `user2` indicating the expiration of the lifecycle of the `user1` variable lifecycle meaning that we can't use it anymore later in code. Similarly, when we called the function `print_username1`, the ownership moved from `user2` to `user3` indicating the expiration of the lifecycle of the `user2` variable lifecycle meaning that we can't use it anymore later in code.

### Example 2

Example two shows how we can use the `user2` variable again in code by passing an immutable borrow of the `user2` variable to the `print_username1` function instead of the actual `user2` variable. This allows us to use the `user2` again later in the code without moving the ownership to `user3`. Note that this is applicable for read operations only.

### Example 3

Example three shows how we can use the `user2` variable again in code, but this time, we want to modify its value (perform a write operation). This can be done by passing a mutable borrow of the `user2` variable to the `update_username3` function that writes on the `name` field of the `user` object.
