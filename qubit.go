package main

import (
	"math"
	"math/cmplx"
	"math/rand"
)

/*type qubit struct {
	val0 complex128
	val1 complex128
}

func (q *qubit) measure() {
	prob := cmplx.Abs(q.val0)
	if rand.Float64() <= prob {
		q.val0 = 1
		q.val1 = 0
	} else {
		q.val0 = 0
		q.val1 = 1
	}

}

func (q *qubit) normalize() {
	mod := q.val0*q.val0 + q.val1*q.val1
	mod = cmplx.Sqrt(mod)
	q.val0 /= mod
	q.val1 /= mod
}

func (q *qubit) getValue() string {
	if q.val0 == 1 {
		return "0"
	} else if q.val1 == 1 {
		return "1"
	}
	return fmt.Sprintf("%f, %f", q.val0, q.val1)
}

/*func newQubit() qubit {
	return qubit{
		val0: 1,
		val1: 0,
	}
}*/

type entanglement struct {
	entangled qubit
	valIf0    complex128
	valIf1    complex128
}

type qubit struct {
	vector matrix
}

func newQubit(n int) qubit {
	if n <= 1 {
		return qubit{
			vector: [][]complex128{{1}, {0}},
		}
	}
	totalStates := int(math.Pow(2, float64(n)))
	v := newVectorVertSize(totalStates)
	v[0][0] = 1
	return qubit{v}
}

func (q *qubit) nQubit() int {
	return int(math.Log2(float64(len(q.vector))))
}

/*
	func (q *qubit) measure(index int) *qubit {
		prob := float64(0)
		rand.Seed(time.Now().UnixMilli())
		dice := rand.Float64()
		aux := float64(0)
		fmt.Printf("Dice: %f\n", dice)
		for i := range q.vector {
			prob = q.prob(i) + aux
			fmt.Printf("Probability: %f\n", prob)
			if dice <= prob {
				q.vector[i][0] = 1
				for i += 1; i < len(q.vector); i++ {
					q.vector[i][0] = 0
				}
				return &qubit{
					vector: q.vector,
				}
			}
			q.vector[i][0] = 0
			aux += prob
		}
		return &qubit{
			vector: [][]complex128{{0}, {1}},
		}
	}
*/

func (q *qubit) measure(index int) *qubit {
	/*n := q.nQubit()
	valueReference := int(math.Pow(2, float64(n-index)))
	rest := valueReference / 2

	zeroIndex, oneIndex := make([]int, 0), make([]int, 0)
	zeroProbability := float64(0)
	for i, p := range q.allProbs() {
		//fmt.Sprintf("Probability for %v: %v\n", i, p)
		if i%valueReference < rest {
			zeroIndex = append(zeroIndex, i)
			zeroProbability += p
			continue
		}
		oneIndex = append(oneIndex, i)
	}*/
	zeroIndex, oneIndex, zeroProbability := q.indexes(index)
	dice := rand.Float64()
	//fmt.Sprintf("ZeroProb: %v\n\n\n\n", zeroProbability)
	res := newQubit(1)
	if dice > zeroProbability { //Sale 1, los valores en los que el qubit es 0 se ponen a probabilidad 0
		for _, i := range zeroIndex {
			q.vector[i][0] = 0i
		}
		res.vector[0][0] = 0
		res.vector[1][0] = 1
	} else { //Sale 0, los valores en los que el qubit es 1 se ponen a probabilidad 0
		for _, i := range oneIndex {
			q.vector[i][0] = 0i
		}
	}
	q.normalize()
	return &res
}

func (q *qubit) prob(index int) float64 {
	aux := cmplx.Abs(q.vector[index][0])
	return aux * aux
}

func (q *qubit) allProbs() []float64 {
	res := make([]float64, len(q.vector))
	for i := range q.vector {
		res[i] = q.prob(i)
	}
	return res
}

/*
	func (q *qubit) getState(index int) qState {
		if index < 0 || index >= len(q.vector) {
			panic("Valor no existente")
		}
		return qState{
			val:  q.vector[index][0],
			prob: q.prob(index),
		}
	}

	func (q *qubit) getAllStates() []qState {
		var res []qState
		for i := range q.vector {
			aux := q.getState(i)
			if aux.prob >= 0.00001 { //Estado imposible
				res = append(res, aux)
			}
		}
		return res
	}
*/
func (q *qubit) normalize() {
	list := q.allProbs()
	sum := 0.0
	for i := range list {
		sum += list[i]
	}
	sum = math.Sqrt(sum)
	for _, elem := range q.vector {
		elem[0] /= complex(sum, 0)
	}
}

func (q *qubit) clone() *qubit {
	var res qubit
	res.vector = make(matrix, len(q.vector))
	for i := 0; i < len(q.vector); i++ {
		res.vector[i] = make([]complex128, 1)
		res.vector[i][0] = q.vector[i][0]
	}
	return &res
}

func (q *qubit) indexes(index int) ([]int, []int, float64) {
	n := q.nQubit()
	valueReference := int(math.Pow(2, float64(n-index)))
	rest := valueReference / 2

	zeroIndex, oneIndex := make([]int, 0), make([]int, 0)
	zeroProbability := float64(0)
	for i, p := range q.allProbs() {
		//fmt.Sprintf("Probability for %v: %v\n", i, p)
		if i%valueReference < rest {
			zeroIndex = append(zeroIndex, i)
			zeroProbability += p
			continue
		}
		oneIndex = append(oneIndex, i)
	}

	return zeroIndex, oneIndex, zeroProbability
}
