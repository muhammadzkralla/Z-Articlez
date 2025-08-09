use std::{fs::File, io::Read, vec};

#[derive(Debug, Clone)]
pub enum CpInfo {
    Utf8 {
        length: u16,
        bytes: Vec<u8>,
    },
    Integer {
        bytes: u32,
    },
    Float {
        bytes: u32,
    },
    Long {
        high_bytes: u32,
        low_bytes: u32,
    },
    Double {
        high_bytes: u32,
        low_bytes: u32,
    },
    Class {
        name_index: u16,
    },
    String {
        string_index: u16,
    },
    Fieldref {
        class_index: u16,
        name_and_type_index: u16,
    },
    Methodref {
        class_index: u16,
        name_and_type_index: u16,
    },
    InterfaceMethodref {
        class_index: u16,
        name_and_type_index: u16,
    },
    NameAndType {
        name_index: u16,
        descriptor_index: u16,
    },
    MethodHandle {
        reference_kind: u8,
        reference_index: u16,
    },
    MethodType {
        descriptor_index: u16,
    },
    InvokeDynamic {
        bootstrap_method_attr_index: u16,
        name_and_type_index: u16,
    },
    // Placeholder for double/long entries (they take 2 slots)
    Empty,
}

struct ClassFile {
    magic: u32,
    minor: u16,
    major: u16,
    constant_pool_count: u16,
    constant_pool: Vec<CpInfo>,
    access_flags: u16,
    this_class: u16,
    super_class: u16,
    interfaces_count: u16,
    interfaces: Vec<u16>,
}

impl ClassFile {
    fn new() -> Self {
        ClassFile {
            magic: 0,
            minor: 0,
            major: 0,
            constant_pool_count: 0,
            constant_pool: Vec::new(),
            access_flags: 0,
            this_class: 0,
            super_class: 0,
            interfaces_count: 0,
            interfaces: Vec::new(),
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

    let mut offset = get_constant_pool(&buf, &mut class_file);

    get_access_flags(&buf, &mut class_file, &mut offset);

    get_this_class(&buf, &mut class_file, &mut offset);

    get_super_class(&buf, &mut class_file, &mut offset);

    get_interfaces_count(&buf, &mut class_file, &mut offset);

    get_interfaces(&buf, &mut class_file, &mut offset);

    println!("Magic: 0x{:X}", class_file.magic);

    println!("Minor: 0x{:X}", class_file.minor);

    println!("Major: 0x{:X}", class_file.major);

    println!("Constant Pool Count: {}", class_file.constant_pool_count);

    print_constant_pool(&class_file);

    print_access_flags(&class_file);

    println!("This Class: {}", class_file.this_class);

    println!("Super Class: {}", class_file.super_class);

    println!("Interfaces Count: {}", class_file.interfaces_count);

    print_interfaces(&class_file);

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

fn get_constant_pool(buf: &Vec<u8>, class_file: &mut ClassFile) -> usize {
    let mut offset = 10;
    let pool_count = class_file.constant_pool_count as usize;

    // Initialize with empty entries
    class_file.constant_pool = vec![CpInfo::Empty; pool_count];

    // Constant pool is 1-indexed
    let mut i = 1;
    while i < pool_count {
        // Take the byte of the tag and increment offset
        let tag = buf[offset];
        offset += 1;

        let entry = match tag {
            // CONSTANT_Utf8
            1 => {
                // Take the two bytes of the `length` field and increment offset
                let length = u16::from_be_bytes([buf[offset], buf[offset + 1]]);
                offset += 2;

                // Take the `length` bytes and increment offset
                let bytes = buf[offset..offset + length as usize].to_vec();
                offset += length as usize;

                // Return entry
                CpInfo::Utf8 { length, bytes }
            }
            // CONSTANT_Integer
            3 => {
                // Take the four bytes of the `bytes` field and increment offset
                let bytes = u32::from_be_bytes([
                    buf[offset],
                    buf[offset + 1],
                    buf[offset + 2],
                    buf[offset + 3],
                ]);
                offset += 4;

                CpInfo::Integer { bytes }
            }
            // CONSTANT_Float
            4 => {
                // Take the four bytes of the `bytes` field and increment offset
                let bytes = u32::from_be_bytes([
                    buf[offset],
                    buf[offset + 1],
                    buf[offset + 2],
                    buf[offset + 3],
                ]);
                offset += 4;

                CpInfo::Float { bytes }
            }
            // CONSTANT_Long
            5 => {
                // Take the four bytes of the `high_bytes` field and increment offset
                let high_bytes = u32::from_be_bytes([
                    buf[offset],
                    buf[offset + 1],
                    buf[offset + 2],
                    buf[offset + 3],
                ]);
                // Take the four bytes of the `low_bytes` field and increment offset
                let low_bytes = u32::from_be_bytes([
                    buf[offset + 4],
                    buf[offset + 5],
                    buf[offset + 6],
                    buf[offset + 7],
                ]);
                offset += 8;

                CpInfo::Long {
                    high_bytes,
                    low_bytes,
                }
            }
            // CONSTANT_Double
            6 => {
                // Take the four bytes of the `high_bytes` field and increment offset
                let high_bytes = u32::from_be_bytes([
                    buf[offset],
                    buf[offset + 1],
                    buf[offset + 2],
                    buf[offset + 3],
                ]);
                // Take the four bytes of the `low_bytes` field and increment offset
                let low_bytes = u32::from_be_bytes([
                    buf[offset + 4],
                    buf[offset + 5],
                    buf[offset + 6],
                    buf[offset + 7],
                ]);
                offset += 8;

                CpInfo::Double {
                    high_bytes,
                    low_bytes,
                }
            }
            // CONSTANT_Class
            7 => {
                // Take the two bytes of the `name_index` field and increment offset
                let name_index = u16::from_be_bytes([buf[offset], buf[offset + 1]]);
                offset += 2;

                CpInfo::Class { name_index }
            }
            // CONSTANT_String
            8 => {
                // Take the two bytes of the `string_index` field and increment offset
                let string_index = u16::from_be_bytes([buf[offset], buf[offset + 1]]);
                offset += 2;

                CpInfo::String { string_index }
            }
            // CONSTANT_Fieldref
            9 => {
                // Take the two bytes of the `class_index` field
                let class_index = u16::from_be_bytes([buf[offset], buf[offset + 1]]);

                // Take the two bytes of the `name_and_type_index` field
                let name_and_type_index = u16::from_be_bytes([buf[offset + 2], buf[offset + 3]]);

                // increment offset
                offset += 4;

                CpInfo::Fieldref {
                    class_index,
                    name_and_type_index,
                }
            }
            // CONSTANT_Methodref
            10 => {
                // Take the two bytes of the `class_index` field
                let class_index = u16::from_be_bytes([buf[offset], buf[offset + 1]]);

                // Take the two bytes of the `name_and_type_index` field
                let name_and_type_index = u16::from_be_bytes([buf[offset + 2], buf[offset + 3]]);

                // increment offset
                offset += 4;

                CpInfo::Methodref {
                    class_index,
                    name_and_type_index,
                }
            }
            // CONSTANT_InterfaceMethodref
            11 => {
                // Take the two bytes of the `class_index` field
                let class_index = u16::from_be_bytes([buf[offset], buf[offset + 1]]);

                // Take the two bytes of the `name_and_type_index` field
                let name_and_type_index = u16::from_be_bytes([buf[offset + 2], buf[offset + 3]]);

                // increment offset
                offset += 4;

                CpInfo::InterfaceMethodref {
                    class_index,
                    name_and_type_index,
                }
            }
            // CONSTANT_NameAndType
            12 => {
                // Take the two bytes of the `name_index` field
                let name_index = u16::from_be_bytes([buf[offset], buf[offset + 1]]);

                // Take the two bytes of the `descriptor_index` field
                let descriptor_index = u16::from_be_bytes([buf[offset + 2], buf[offset + 3]]);

                // increment offset
                offset += 4;

                CpInfo::NameAndType {
                    name_index,
                    descriptor_index,
                }
            }
            // CONSTANT_MethodHandle
            15 => {
                // Take the byte of the `reference_kind` field
                let reference_kind = buf[offset];

                // Take the two bytes of the `reference_index` field
                let reference_index = u16::from_be_bytes([buf[offset + 1], buf[offset + 2]]);

                // increment offset
                offset += 3;

                CpInfo::MethodHandle {
                    reference_kind,
                    reference_index,
                }
            }
            // CONSTANT_MethodType
            16 => {
                // Take the two bytes of the `descriptor_index` and increment offset
                let descriptor_index = u16::from_be_bytes([buf[offset], buf[offset + 1]]);
                offset += 2;

                CpInfo::MethodType { descriptor_index }
            }
            // CONSTANT_InvokeDynamic
            18 => {
                // Take the two bytes of the `bootstrap_method_attr_index`
                let bootstrap_method_attr_index =
                    u16::from_be_bytes([buf[offset], buf[offset + 1]]);

                // Take the two bytes of the `name_and_type_index`
                let name_and_type_index = u16::from_be_bytes([buf[offset + 2], buf[offset + 3]]);

                // increment offset
                offset += 4;

                CpInfo::InvokeDynamic {
                    bootstrap_method_attr_index,
                    name_and_type_index,
                }
            }

            // default
            _ => {
                panic!("Unknown constant pool tag: {}", tag);
            }
        };

        // Now that we parsed the entry, we need to store it in the `constant_pool` field
        // of the class_file
        // TODO: Why we need deep copying here??
        class_file.constant_pool[i] = entry.clone();

        // Long and Double entries take up two slots, so we need to assign the next
        // entry to them as empty and jump to the next entry
        if matches!(entry, CpInfo::Long { .. } | CpInfo::Double { .. }) {
            // Set the next entry as empty and skip one entry
            i += 2;
            class_file.constant_pool[i] = CpInfo::Empty;
        } else {
            // Move to process the next entry
            i += 1;
        }
    }

    offset
}

fn get_access_flags(buf: &Vec<u8>, class_file: &mut ClassFile, offset: &mut usize) {
    let access_flags = u16::from_be_bytes([buf[*offset], buf[*offset + 1]]);
    class_file.access_flags = access_flags;

    *offset += 2;
}

fn get_this_class(buf: &Vec<u8>, class_file: &mut ClassFile, offset: &mut usize) {
    let this_class = u16::from_be_bytes([buf[*offset], buf[*offset + 1]]);
    class_file.this_class = this_class;

    *offset += 2;
}

fn get_super_class(buf: &Vec<u8>, class_file: &mut ClassFile, offset: &mut usize) {
    let super_class = u16::from_be_bytes([buf[*offset], buf[*offset + 1]]);
    class_file.super_class = super_class;

    *offset += 2;
}

fn get_interfaces_count(buf: &Vec<u8>, class_file: &mut ClassFile, offset: &mut usize) {
    let interfaces_count = u16::from_be_bytes([buf[*offset], buf[*offset + 1]]);
    class_file.interfaces_count = interfaces_count;

    *offset += 2;
}

fn get_interfaces(buf: &Vec<u8>, class_file: &mut ClassFile, offset: &mut usize) {
    let interfaces_count = class_file.interfaces_count;

    if interfaces_count == 0 {
        return;
    }

    for _ in 0..interfaces_count {
        let current_interface_ref = u16::from_be_bytes([buf[*offset], buf[*offset + 1]]);
        class_file.interfaces.push(current_interface_ref);
        *offset += 2;
    }
}

fn print_constant_pool(class_file: &ClassFile) {
    println!("\nConstant Pool:");

    for (i, entry) in class_file.constant_pool.iter().enumerate() {
        // Constant pool is 1-indexed
        if i == 0 {
            continue;
        }

        match entry {
            CpInfo::Utf8 { length, bytes } => {
                let string = String::from_utf8_lossy(bytes);
                println!("  #{}: Utf8 [{}]", i, string);
            }
            CpInfo::Integer { bytes } => {
                println!("  #{}: Integer [{}]", i, *bytes as i32);
            }
            CpInfo::Float { bytes } => {
                let float_val = f32::from_bits(*bytes);
                println!("  #{}: Float [{}]", i, float_val);
            }
            CpInfo::Long {
                high_bytes,
                low_bytes,
            } => {
                // AS SPECIFIED BY THE SPECS:
                // ((long) high_bytes << 32) + low_bytes
                let long = ((*high_bytes as u64) << 32) + (*low_bytes as u64);
                println!("  #{}: Long [{}]", i, long as i64);
            }
            CpInfo::Double {
                high_bytes,
                low_bytes,
            } => {
                // AS SPECIFIED BY THE SPECS:
                // ((long) high_bytes << 32) + low_bytes
                let bits = ((*high_bytes as u64) << 32) + (*low_bytes as u64);
                let double = f64::from_bits(bits);
                println!("  #{}: Double [{}]", i, double);
            }
            CpInfo::Class { name_index } => {
                println!("  #{}: Class [name_index=#{}]", i, name_index);
            }
            CpInfo::String { string_index } => {
                println!("  #{}: String [string_index=#{}]", i, string_index);
            }
            CpInfo::Fieldref {
                class_index,
                name_and_type_index,
            } => {
                println!(
                    "  #{}: Fieldref [class_index=#{}, name_and_type_index=#{}]",
                    i, class_index, name_and_type_index
                );
            }
            CpInfo::Methodref {
                class_index,
                name_and_type_index,
            } => {
                println!(
                    "  #{}: Methodref [class_index=#{}, name_and_type_index=#{}]",
                    i, class_index, name_and_type_index
                );
            }
            CpInfo::InterfaceMethodref {
                class_index,
                name_and_type_index,
            } => {
                println!(
                    "  #{}: InterfaceMethodref [class_index=#{}, name_and_type_index=#{}]",
                    i, class_index, name_and_type_index
                );
            }
            CpInfo::NameAndType {
                name_index,
                descriptor_index,
            } => {
                println!(
                    "  #{}: NameAndType [name_index=#{}, descriptor_index=#{}]",
                    i, name_index, descriptor_index
                );
            }
            CpInfo::MethodHandle {
                reference_kind,
                reference_index,
            } => {
                println!(
                    "  #{}: MethodHandle [reference_kind={}, reference_index=#{}]",
                    i, reference_kind, reference_index
                );
            }
            CpInfo::MethodType { descriptor_index } => {
                println!(
                    "  #{}: MethodType [descriptor_index=#{}]",
                    i, descriptor_index
                );
            }
            CpInfo::InvokeDynamic {
                bootstrap_method_attr_index,
                name_and_type_index,
            } => {
                println!(
                    "  #{}: InvokeDynamic [bootstrap_method_attr_index={}, name_and_type_index=#{}]",
                    i, bootstrap_method_attr_index, name_and_type_index
                );
            }
            CpInfo::Empty => {
                println!("EMPTY ENTRY!")
            }
        }
    }
}

fn print_access_flags(class_file: &ClassFile) {
    println!("\nAccess Flags: 0x{:04X}", class_file.access_flags);

    let flags = class_file.access_flags;
    let mut flag_names = Vec::new();

    // Check each access flag bit according to JVM spec
    if flags & 0x0001 != 0 {
        flag_names.push("ACC_PUBLIC");
    }
    if flags & 0x0010 != 0 {
        flag_names.push("ACC_FINAL");
    }
    if flags & 0x0020 != 0 {
        flag_names.push("ACC_SUPER");
    }
    if flags & 0x0200 != 0 {
        flag_names.push("ACC_INTERFACE");
    }
    if flags & 0x0400 != 0 {
        flag_names.push("ACC_ABSTRACT");
    }
    if flags & 0x1000 != 0 {
        flag_names.push("ACC_SYNTHETIC");
    }
    if flags & 0x2000 != 0 {
        flag_names.push("ACC_ANNOTATION");
    }
    if flags & 0x4000 != 0 {
        flag_names.push("ACC_ENUM");
    }
    if flags & 0x8000 != 0 {
        flag_names.push("ACC_MODULE");
    }

    if flag_names.is_empty() {
        println!("  No access flags set");
    } else {
        println!("  Flags: {}", flag_names.join(", "));
    }

    println!()
}

fn print_interfaces(class_file: &ClassFile) {
    if class_file.interfaces.is_empty() {
        println!("Interfaces: None");
        return;
    }

    println!("Interfaces:");

    for (i, interface_ref) in class_file.interfaces.iter().enumerate() {
        println!("  [{}]: #{}", i, interface_ref);
    }
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
