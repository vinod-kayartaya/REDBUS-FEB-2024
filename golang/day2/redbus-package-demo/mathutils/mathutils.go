package mathutils

func Factorial(num int) int {
	if num < 0 {
		panic("mathutils.Factorial does not work with negative numbers")
	}

	f := 1
	for num > 0 {
		f *= num
		num--
	}
	return f
}
