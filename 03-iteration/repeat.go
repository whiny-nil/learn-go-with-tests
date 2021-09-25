package iteration

func Repeat(text string, times int) string {
	var ret string

	for i := 0; i < times; i++ {
		ret += text
	}

	return ret
}
