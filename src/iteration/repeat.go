package iteration

func Repeat(ch string, times int) string {
	var rep string
	for i := 0; i < times; i++ {
		rep += ch
	}
	return rep
}
