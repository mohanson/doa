package doa

// Doa1 test if a condition returns true
func Doa1(a bool) {
	if !a {
		panic("assertion error")
	}
}

// Doa2 likes Doa1.
func Doa2(a bool, b bool) {
	if !(a || b) {
		panic("assertion error")
	}
}

// Try1 panics when it receives a error.
func Try1(err error) {
	if err != nil {
		panic(err)
	}
}

// Try2 panics when it receives a error.
func Try2(a interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return a
}

// Try3 panics when it receives a error.
func Try3(a interface{}, b interface{}, err error) (interface{}, interface{}) {
	if err != nil {
		panic(err)
	}
	return a, b
}

// Try4 panics when it receives a error.
func Try4(a interface{}, b interface{}, c interface{}, err error) (interface{}, interface{}, interface{}) {
	if err != nil {
		panic(err)
	}
	return a, b, c
}
