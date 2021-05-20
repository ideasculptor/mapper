package mapper

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
	for pos, _ := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}
