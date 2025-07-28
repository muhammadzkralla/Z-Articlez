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

---

## Functions in Rust

Function example with arguments:

```rust
fn main() {
    print_labeled_measurement(5, 'h');
}

fn print_labeled_measurement(value: i32, unit_label: char) {
    println!("The measurement is: {value}{unit_label}");
}
```

Functions in Rust are of two types:

- Statements: instructions that perform some action and do not return a value.
- Expressions: evaluate to a resultant value. Let’s look at some examples.

```rust
fn main() {
    let y = {
        let x = 3;
        x + 1
    };

    println!("The value of y is: {y}");
}
```

The value of the variable `y` here will be 4.

> [!NOTE]
> Rust is an expression-based language.

In Rust, the return value of the function is synonymous with the value of the final expression in the block of the body of a function. You can return early from a function by using the return keyword and specifying a value, but most functions return the last expression implicitly and it does not contain a semicolon.

---

## Control Flow in Rust

### If Conditions in Rust

If condition syntax in Rust:

```rust
fn main() {
    let number = 3;

    if number < 5 {
        println!("condition was true");
    } else {
        println!("condition was false");
    }
}
```

> [!NOTE]
> It’s also worth noting that the condition in this code must be a `bool`. If the condition isn’t a `bool`, we’ll get an error.

Using too many else if expressions can clutter your code, so if you have more than one, you might want to refactor your code. Chapter 6 describes a powerful Rust branching construct called match for these cases.

Because if is an expression, we can use it on the right side of a let statement to assign the outcome to a variable, something similar to the ternary operator, for example:

```rust
fn main() {
    let condition = true;
    let number = if condition { 5 } else { 6 };

    println!("The value of number is: {number}");
}
```

## Repetition with Loops in Rust

Rust has three kinds of loops: `loop`, `while`, and `for`.

### Repeating Code with Loop:

The `loop` keyword tells Rust to execute a block of code over and over again forever or until you explicitly tell it to stop, for example:

```rust
fn main() {
    loop {
        println!("again!");
    }
}
```

This will infinitely print "again!" in the terminal until you interrupt it with `SIGKILL_SIGNAL` manually.

> [!NOTE]
> You can avoid this by using the `break` keyword to step out of the loop programmatically.

### Loop in Rust

If you have loops within loops, break and continue apply to the innermost loop at that point. You can optionally specify a loop label on a loop that you can then use with break or continue to specify that those keywords apply to the labeled loop instead of the innermost loop.

Loop labels must begin with a single quote. Here’s an example with two nested loops:

```rust
fn main() {
    let mut count = 0;
    'counting_up: loop {
        println!("count = {count}");
        let mut remaining = 10;

        loop {
            println!("remaining = {remaining}");
            if remaining == 9 {
                break;
            }
            if count == 2 {
                break 'counting_up;
            }
            remaining -= 1;
        }

        count += 1;
    }
    println!("End count = {count}");
}
```

### While Loops in Rust

Here's an example syntax for the while loops in Rust:

```rust
fn main() {
    let mut number = 3;

    while number != 0 {
        println!("{number}!");

        number -= 1;
    }

    println!("LIFTOFF!!!");
}
```

### For Loops in Rust

Here's an example syntax for the "for each" loops in Rust:

```rust
fn main() {
    let a = [10, 20, 30, 40, 50];

    for element in a {
        println!("the value is: {element}");
    }
}
```

Here's an example syntax for the "for range" loops in Rust:

```rust
fn main() {
    for number in (1..4).rev() {
        println!("{number}!");
    }
    println!("LIFTOFF!!!");
}
```

---

## The Ownership Model in Rust

This Rust code:

```rust
    let mut s = "test";
```

Is equivalent to this C code:

```c
    const char* s = "test";
```

The string literals are stored in read-only memory in both languages. You can reassign their values but you can not modify their values. In code this looks like:

```rust
    let mut s = "test";
    s = "best"; // s is a &str reference; you're changing what it points to
```

And in C:

```c
    char *s = "test";
    s = "best"; // This just changes the pointer's value
```

But you can not do the following:

```rust
    let mut s = "test";
    s.replace_range(0..1, "B"); // compile error — &str is immutable
```

```c
    char *s = "test";
    s[0] = 'B'; // Undefined behavior — modifying read-only memory
```

TLDR: The pointer is stored on the stack while the string literal is stored in the read-only memory. You can reassign the reference of the pointer, but you can not modify the value of the string literal, just like in C.

While this Rust code:

```rust
    let mut sss = String::from("test2");
```

Is equivalent to this C code:

```c
    char* sss = malloc(strlen("test2") + 1);
    strcpy(sss, "test2");
```

In order to support a mutable, growable piece of text, we need to allocate an amount of memory on the heap, unknown at compile time, to hold the contents. This means:

- The memory must be requested from the memory allocator at runtime.
- We need a way of returning this memory to the allocator when we’re done with our String.

That first part is done by us: when we call String::from, its implementation requests the memory it needs. This is pretty much universal in programming languages.

However, the second part is different. In languages with a garbage collector (GC), the GC keeps track of and cleans up memory that isn’t being used anymore, and we don’t need to think about it. In most languages without a GC, it’s our responsibility to identify when memory is no longer being used and to call code to explicitly free it, just as we did to request it. Doing this correctly has historically been a difficult programming problem. If we forget, we’ll waste memory. If we do it too early, we’ll have an invalid variable. If we do it twice, that’s a bug too. We need to pair exactly one allocate with exactly one free.

Rust takes a different path: the memory is automatically returned once the variable that owns it goes out of scope.

When a variable goes out of scope, Rust calls a special function for us. This function is called drop, and it’s where the author of String can put the code to return the memory. Rust calls drop automatically at the closing curly bracket. Similar to the Resource Acquisition Is Initialization principle in C++.

### Variables and Data Interacting with Move

This will not move ownership of x to y:

```rust
    let x = 5;
    let y = x;
```

While this will move the ownership of s1 to s2:

```rust
    let s1 = String::from("hello");
    let s2 = s1; // s1 now is not valid anymore
```

This happens to avoid the very famous memory-safety bug, the double freeing issue.

In addition, there’s a design choice that’s implied by this: Rust will never automatically create “deep” copies of your data. Therefore, any automatic copying can be assumed to be inexpensive in terms of runtime performance.

### Scope and Assignment

For this Rust code:

```rust
    let mut s = String::from("hello");
    s = String::from("ahoy");

    println!("{s}, world!");
```

The "hello" string literal is dropped as nothing is pointing to it anymore.

### Variables and Data Interacting with Clone

If we do want to deeply copy the heap data of the String, not just the stack data, we can use a common method called clone.

For example:

```rust
    let s1 = String::from("hello");
    let s2 = s1.clone();

    println!("s1 = {s1}, s2 = {s2}");
```

This performs a heap copy too, so each variable of `s1` and `s2` are pointing to completely different parts of the heap with the exact same values. Please note that this is an expensive operation.

### Stack-Only Data: Copy

For this Rust code:

```rust
    let x = 5;
    let y = x;

    println!("x = {x}, y = {y}");
```

The variable `x` is not moved to the variable `y`, and instead the variable `y` gets a copy of the value of the variable `x` which is 5, this means that the value of `x` will be 5 and it's a valid variable and the value of `y` will be 5 and it's a valid variable too.

The reason is that types such as integers that have a known size at compile time are stored entirely on the stack, so copies of the actual values are quick to make.

That means there’s no reason we would want to prevent x from being valid after we create the variable y. In other words, there’s no difference between deep and shallow copying here, so calling clone wouldn’t do anything different from the usual shallow copying, and we can leave it out.

### The Copy Trait

Rust has a special annotation called the Copy trait that we can place on types that are stored on the stack, as integers are (we’ll talk more about traits in Chapter 10). If a type implements the Copy trait, variables that use it do not move, but rather are trivially copied, making them still valid after assignment to another variable.

Rust won’t let us annotate a type with Copy if the type, or any of its parts, has implemented the Drop trait.

So, what types implement the Copy trait? You can check the documentation for the given type to be sure, but as a general rule, any group of simple scalar values can implement Copy, and nothing that requires allocation or is some form of resource can implement Copy. Here are some of the types that implement Copy:

- All the integer types
- The boolean types
- All the floating-point types
- The character types
- Tuples, if they only contain types that also implement copy.

### Return Values and Scope

The ownership of a variable follows the same pattern every time: assigning a value to another variable moves it. When a variable that includes data on the heap goes out of scope, the value will be cleaned up by drop unless ownership of the data has been moved to another variable.

While this works, taking ownership and then returning ownership with every function is a bit tedious. What if we want to let a function use a value but not take ownership? It’s quite annoying that anything we pass in also needs to be passed back if we want to use it again, in addition to any data resulting from the body of the function that we might want to return as well.

Rust does let us return multiple values using a tuple:

```rust
fn main() {
    let s1 = String::from("hello");

    let (s2, len) = calculate_length(s1);

    println!("The length of '{s2}' is {len}.");
}

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len();

    (s, length)
}
```
