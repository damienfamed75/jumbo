#[no_mangle]
pub unsafe extern "C" fn hello(name: *mut i8, name_len: usize) {
	// Create a usable format to manipulate the given string by making a slice
	// of the mutable c_char's with a maximum value.
	use std::slice;
	let new_str: &mut [i8] = slice::from_raw_parts_mut(name, name_len);

	// Change the first character's value for demo purposes.
	// new_str[0] = 'D' as raw::c_char;
	new_str[0] = 'D' as i8;

	new_str[name_len-1] = 'd' as i8;
}
