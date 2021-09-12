package fcz

type Factor struct {
	Factor uint64
	Degree uint8
}

type Factors []Factor

func Factorize(num uint64) (answer Factors) {
	answer = append(answer, Factor{Factor: num, Degree: 1})
	return
}
