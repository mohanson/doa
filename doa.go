package doa

func Try1(err error) {
	if err != nil {
		panic(err)
	}
}

func Try2(a interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return a
}

func Try3(a interface{}, b interface{}, err error) (interface{}, interface{}) {
	if err != nil {
		panic(err)
	}
	return a, b
}

func Try4(a interface{}, b interface{}, c interface{}, err error) (interface{}, interface{}, interface{}) {
	if err != nil {
		panic(err)
	}
	return a, b, c
}
