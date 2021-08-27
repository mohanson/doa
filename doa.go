package doa

// Doa lets you test if a condition in your code returns true, if not, the program will panic.
func Doa(b bool) {
	if !b {
		panic("unreachable")
	}
}

// Nil lets you test if an error in your code is nil, if not, the program will panick.
func Nil(err error) {
	if err != nil {
		panic(err)
	}
}

// Try will give you the embedded value if there is no error returns. If instead error then it will panic.
func Try(a interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return a
}
