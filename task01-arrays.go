package homework

func average(input [15]float32) (result float32) {
	var average float32 = 0
	l := len(input)
	for i := 0; i < l; i++ {
		average += input[i]
	}
	return (average / float32(l))
}
