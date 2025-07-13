# Notes

## longest.rs

`longest()` is a function that is valid over the lifecycle `'a` and takes two arguments, `x` which is a borrow that is valid over the generic lifecycle `'a` of type string, and `y` which is a borrow that is valid over the generic lifecycle `'a` of type string, and returns a string that is valid over the generic lifecycle `'a`.

Note that rust infers the value of the generic lifecycle `'a` automatically and we don't worry about this part. We just suppose that it's the shortest lifecycle of the given arguments.

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
