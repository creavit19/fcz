package fcz

import (
	"math"
	"math/big"
)

type Factor struct {
	Factor uint64
	Degree uint8
}

type Factors []Factor

func Factorize(num uint64) (answer Factors) {

	if num == 0 {answer = append(answer, Factor{Factor: 0, Degree: 1});return}
	if num == 1 {answer = append(answer, Factor{Factor: 1, Degree: 1});return}

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
		if remain == 1 { return true }
		sqrtRemain = sqrtCeil(remain)
		return false
	}

	for _, currentNum = range primeNumbers {
		if remain % currentNum == 0 {
			if recInAnswer() {
				return
			}
		}
	}

	for sqrtRemain >= currentNum {

		currentNum += 2

		if currentNum % 3 == 0 { currentNum += 2 }
		if currentNum % 5 == 0 { continue }
		if currentNum % 7 == 0 { continue }
		if currentNum % 11 == 0 { continue }
		if currentNum % 13 == 0 { continue }
		if currentNum % 17 == 0 { continue }
		if currentNum % 19 == 0 { continue }

		if remain % currentNum != 0 { continue }
		if recInAnswer() { return }

	}

	answer = append(answer, Factor{Factor: remain, Degree: 1})

	return

}

type BFactor struct {
	Factor *big.Int
	Degree uint16
}

type BFactors []BFactor

func FactorizeBig(num *big.Int) (answer BFactors) {

	if num.ProbablyPrime(16) {
		answer = append(answer, BFactor{Factor: big.NewInt(0).Abs(num), Degree: 1})
		return
	}

	Int0 := big.NewInt(0)
	Int1 := big.NewInt(1)

	if num.Cmp(Int0) == 0 {answer = append(answer, BFactor{Factor: Int0, Degree: 1});return}
	if num.Cmp(Int1) == 0 {answer = append(answer, BFactor{Factor: Int1, Degree: 1});return}

	Int2 := big.NewInt(2)
	Int3 := big.NewInt(3)
	Int5 := big.NewInt(5)
	Int7 := big.NewInt(7)
	Int11 := big.NewInt(11)
	Int13 := big.NewInt(13)
	Int17 := big.NewInt(17)
	Int19 := big.NewInt(19)

	remain := big.NewInt(0)
	remain.Abs(num)
	sqrtRemain := big.NewInt(0)
	sqrtRemain.Sqrt(remain)
	currentNum := big.NewInt(19)
	cN := currentNum
	primeNumbers := [8]*big.Int{Int2, Int3, Int5, Int7, Int11, Int13, Int17, Int19}
	mod := big.NewInt(0)

	recInAnswer := func () bool {
		degree := uint16(0)
		for mod.Mod(remain, currentNum).Cmp(Int0) == 0 {
			remain.Div(remain, currentNum)
			degree++
		}
		answer = append(answer, BFactor{Factor: big.NewInt(0).Abs(currentNum), Degree: degree})
		if remain.Cmp(Int1) == 0 { return true }
		sqrtRemain.Sqrt(remain)
		return false
	}

	for _, currentNum = range primeNumbers {
		if mod.Mod(remain, currentNum).Cmp(Int0) == 0 {
			if recInAnswer() {
				return
			}
		}
	}

	if remain.ProbablyPrime(16) {
		answer = append(answer, BFactor{Factor: big.NewInt(0).Abs(remain), Degree: 1})
		return
	}

	currentNum = cN

	for sqrtRemain.Cmp(currentNum) >= 0 {

		currentNum.Add(currentNum, Int2)

		if mod.Mod(currentNum, Int3).Cmp(Int0) == 0 { currentNum.Add(currentNum, Int2) }
		if mod.Mod(currentNum, Int5).Cmp(Int0) == 0 { continue }
		if mod.Mod(currentNum, Int7).Cmp(Int0) == 0 { continue }
		if mod.Mod(currentNum, Int11).Cmp(Int0) == 0 { continue }
		if mod.Mod(currentNum, Int13).Cmp(Int0) == 0 { continue }
		if mod.Mod(currentNum, Int17).Cmp(Int0) == 0 { continue }
		if mod.Mod(currentNum, Int19).Cmp(Int0) == 0 { continue }

		if mod.Mod(remain, currentNum).Cmp(Int0) != 0 { continue }
		if recInAnswer() { return }
		if remain.ProbablyPrime(16) { break }

	}

	answer = append(answer, BFactor{Factor: big.NewInt(0).Abs(remain), Degree: 1})

	return

}
