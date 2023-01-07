package main

import (
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

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

func (q *qubit) measure(index int) *qubit {
	zeroIndex, oneIndex, zeroProbability := q.indexes(index)
	rand.Seed(time.Now().UnixMicro())
	dice := rand.Float64()
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
		if i%valueReference < rest {
			zeroIndex = append(zeroIndex, i)
			zeroProbability += p
			continue
		}
		oneIndex = append(oneIndex, i)
	}

	return zeroIndex, oneIndex, zeroProbability
}
