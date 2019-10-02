#[no_mangle]
pub unsafe extern "C" fn hello(name: *mut i8, name_len: usize) {
	// Create a usable format to manipulate the given string by making a slice
	// of the mutable c_char's with a maximum value.
	use std::slice;
	let new_str: &mut [i8] = slice::from_raw_parts_mut(name, name_len);

	// Change the first character's value for demo purposes.
	new_str[0] = 'D' as i8;

	new_str[name_len-1] = 'd' as i8;
}

use std::error::Error;
// use std::path::Path;
// use std::io::Read;
// use std::marker::Sized;
use std::fs::File;
use std::io::prelude::*;
use std::io::{BufReader, BufRead};

pub unsafe extern "C" fn alphabet() -> std::io::Result<()> {
	let f = File::open("input")?;
	let mut reader = BufReader::new(f);

	let mut out = File::create("output")?;
	let mut line = String::new();

	loop {
		reader.read_line(&mut line)?;

		out.write(&mut line.as_bytes())?;
	}

	Ok(())
}