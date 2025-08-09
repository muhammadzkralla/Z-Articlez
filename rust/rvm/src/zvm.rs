use std::{fs::File, io::Read, vec};

struct ClassFile {
    magic: u32,
    minor: u16,
    major: u16,
    constant_pool_count: u16,
}

impl ClassFile {
    fn new() -> Self {
        ClassFile {
            magic: 0,
            minor: 0,
            major: 0,
            constant_pool_count: 0,
        }
    }
}

pub fn zvm() {
    let mut class_file = ClassFile::new();
    let mut file = File::open("Main.class").unwrap();

    let mut buf = Vec::new();
    let _ = file.read_to_end(&mut buf).unwrap();

    get_magic(&buf, &mut class_file);

    get_minor(&buf, &mut class_file);

    get_major(&buf, &mut class_file);

    get_constant_pool_count(&buf, &mut class_file);

    println!("Magic: {:X}", class_file.magic);

    println!("Minor: {:X}", class_file.minor);

    println!("Major: {:X}", class_file.major);

    println!("Constant Pool Count: {:X}", class_file.constant_pool_count);

    // print_hex(&buf);
    //
    // println!();
    //
    // print_bin(&buf);
}

fn get_magic(buf: &Vec<u8>, class_file: &mut ClassFile) {
    let magic = u32::from_be_bytes([buf[0], buf[1], buf[2], buf[3]]);
    class_file.magic = magic;
}

fn get_minor(buf: &Vec<u8>, class_file: &mut ClassFile) {
    let minor = u16::from_be_bytes([buf[4], buf[5]]);
    class_file.minor = minor;
}

fn get_major(buf: &Vec<u8>, class_file: &mut ClassFile) {
    let major = u16::from_be_bytes([buf[6], buf[7]]);
    class_file.major = major;
}

fn get_constant_pool_count(buf: &Vec<u8>, class_file: &mut ClassFile) {
    let constant_pool_count = u16::from_be_bytes([buf[8], buf[9]]);
    class_file.constant_pool_count = constant_pool_count;
}

fn print_hex(buf: &Vec<u8>) {
    for (i, byte) in buf.iter().enumerate() {
        print!("{:02X} ", byte);

        if (i + 1) % 8 == 0 {
            print!(" ");
        }

        if (i + 1) % 16 == 0 {
            println!();
        }
    }
}

fn print_bin(buf: &Vec<u8>) {
    for (i, byte) in buf.iter().enumerate() {
        print!("{:08b} ", byte);

        if (i + 1) % 6 == 0 {
            println!();
        }
    }
}
