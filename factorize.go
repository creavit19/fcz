package fcz

import "math"

type Factor struct {
	Factor uint64
	Degree uint8
}

type Factors []Factor

func Factorize(num uint64) (answer Factors) {

	if(num == 0){answer = append(answer, Factor{Factor: 0, Degree: 1});return}
	if(num == 1){answer = append(answer, Factor{Factor: 1, Degree: 1});return}

	sqrtCeil := func (n uint64) (sq uint64) {
		sq = uint64(math.Ceil(math.Sqrt(float64(n))))
		return
	}

	remain := num
	sqrtRemain := sqrtCeil(num)
	currentNum := uint64(1)
	primeNumbers := [8]uint64{2, 3, 5, 7, 11, 13, 17, 19}

	recInAnswer := func () bool {
		degree := uint8(0)
		for remain % currentNum == 0 {
			remain /= currentNum
			degree++
		}
		answer = append(answer, Factor{Factor: currentNum, Degree: degree})
		if (remain == 1) { return true }
		sqrtRemain = sqrtCeil(remain)
		return false
	}

	for _, currentNum = range primeNumbers {
		if (remain % currentNum == 0) {
			if (recInAnswer()) {
				return
			}
		}
	}

	for sqrtRemain >= currentNum {

		currentNum += 2

		if (currentNum % 3 == 0) { currentNum += 2 }
		if (currentNum % 5 == 0) { continue }
		if (currentNum % 7 == 0) { continue }
		if (currentNum % 11 == 0) { continue }
		if (currentNum % 13 == 0) { continue }
		if (currentNum % 17 == 0) { continue }
		if (currentNum % 19 == 0) { continue }

		if (remain % currentNum != 0) { continue }
		if (recInAnswer()) { return }

	}

	answer = append(answer, Factor{Factor: remain, Degree: 1})

	return

}
