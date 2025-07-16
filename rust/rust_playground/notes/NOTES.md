# Notes

## longest.rs

`longest()` is a function that is valid over the lifecycle `'a` and takes two arguments, `x` which is a borrow of a string that is valid over the generic lifecycle `'a`, and `y` which is a borrow of a string that is valid over the generic lifecycle `'a`, and returns a string that is valid over the generic lifecycle `'a`.

Note that Rust infers the value of the generic lifecycle `'a` automatically and we don't need to worry about this part. We just suppose that it's the shortest lifecycle of the given arguments.

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

## 1 Mutable OR N Immutable Borrows

### Example 4

Example four shows that we can't create an immutable borrow after creating at least one mutable borrow and using that mutable borrow later in code after creating the immutable borrow. This is because that this behaviour will cause the actual data be modified after creating an immutable borrow that expects the object not to be altered during its lifecycle.


### Example 5

Example five addresses how can we avoid this issue, by declaring the immutable borrow after all the mutable borrows of the object are not used later in code anymore. This way, the immutable borrow ensures that the value will not be modified until its lifecycle drops.

### Example 6

Example six shows that even if we dropped the immutable borrow, we still can't use the mutable borrow. Rust's borrow checker says: "You still have r2 alive in this scope. I don't care if you call drop(r2), I won't let you use r1 mutably again in this block."

The `drop()` function drops the value at runtime, but does not shorten the borrow lifetime in the eyes of the compiler.

---

## Variables and Mutability

### Variables are Immutable by Default

All variables are immutable by default unless declared mutable using the `mut` keyword in specifying them.

### Immutable Variables vs Constants

There are some differences between Immutable variables and constants in Rust. Both are used to store values that will not change in the future, but there are some differences. With constants, type of the value must be annotated. Constants may be set only to a constant expression, not the result of a value that could only be computed at runtime. For example:

```rust
const THREE_HOURS_IN_SECONDS: u32 = 60 * 60 * 3;
```

### Shadowing

Shadowing is different from marking a variable as mut because we’ll get a compile-time error if we accidentally try to reassign to this variable without using the let keyword. By using let, we can perform a few transformations on a value but have the variable be immutable after those transformations have been completed. For example:

```rust
fn main() {
    let x = 5;

    let x = x + 1;

    {
        let x = x * 2;
        println!("The value of x in the inner scope is: {x}");
    }

    println!("The value of x is: {x}");
}
```

We can say that we make a variable mutable for some short amount of time by using shadowing.

The other difference between mut and shadowing is that because we’re effectively creating a new variable when we use the let keyword again, we can change the type of the value but reuse the same name. For example, say our program asks a user to show how many spaces they want between some text by inputting space characters, and then we want to store that input as a number:

```rust
    let spaces = "   ";
    let spaces = spaces.len();
```

The first `spaces` variable is a string type and the second `spaces` variable is a number type.

---

## Data Types

Data types can be scalar or compound.

A scalar type represents a single value. Rust has four primary scalar types: `integers`, `floating-point numbers`, `Booleans`, and `characters`.

Rust is a statically-typed language, and it can infer the type of variables sometimes and sometimes we need to specify the type explicitly like when converting a string to a numeric value for example:

```rust
let guess: u32 = "42".parse().expect("Not a number!");
```

## Scalar Data Types

### Integer Types in Rust

| Length | Signed | Unsigned |
| ------ | ------ | -------- |
| 8-bit  | i8     | u8       |
| 16-bit | i16    | u16      |
| 32-bit | i32    | u32      |
| 64-bit | i64    | u64      |
| 128-bit| i128   | u128     |
| arch   | isize  | usize    |

> [!NOTE]
>  Integer types default to i32.
> When you’re compiling in debug mode, Rust includes checks for integer overflow that cause your program to panic at runtime if this behavior occurs.
> When you’re compiling in release mode with the --release flag, Rust does not include checks for integer overflow that cause panics. Instead, if overflow occurs, Rust performs two’s complement wrapping.

### floating-point Types in Rust

Rust’s floating-point types are `f32` and `f64`. The default type is f64 because on modern CPUs, it’s roughly the same speed as f32 but is capable of more precision.

```rust
fn main() {
    let x = 2.0; // f64

    let y: f32 = 3.0; // f32
}
```

## Compound Data Types

Compound types can group multiple values into one type. Rust has two primitive compound types: tuples and arrays.

### Tuple Types in Rust

A tuple group different values of different types together, for example:

```rust
fn main() {
    let tup: (i32, f64, u8) = (500, 6.4, 1);
}
```

To access its values:

```rust
fn main() {
    let tup = (500, 6.4, 1);

    let (x, y, z) = tup;

    println!("The value of y is: {y}");
}
```

Or:

```rust
fn main() {
    let x: (i32, f64, u8) = (500, 6.4, 1);

    let five_hundred = x.0;

    let six_point_four = x.1;

    let one = x.2;
}
```

### Array Types in Rust

Arrays group fixed size of elements of the same type, for example:

```rust
fn main() {
    let a = [1, 2, 3, 4, 5];

    let a: [i32; 5] = [1, 2, 3, 4, 5];

    let a = [3; 5]; // equivalent to [3, 3, 3, 3, 3]
}
```

To access its values:

```rust
fn main() {
    let a = [1, 2, 3, 4, 5];

    let first = a[0];
    let second = a[1];
}
```
