package piscinego

func UltimateDivMod(a, b *int) {
	temp := *a
	*a = temp / *b
	*b = temp % *b
}
