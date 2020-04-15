package reverse

func ReverseInt(a int) (b int) {
	for a > 0 {
		k := a % 10
		b = b*10 + k
		a = a / 10
	}
	return
}
