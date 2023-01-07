package main

import (
	"math"
)

type qGate matrix

// const sqrt float64 = math.Sqrt(2)
const sqrt2 complex128 = 1.414213562373095 //raiz de 2

func H() qGate {
	return [][]complex128{
		{1 / sqrt2, 1 / sqrt2},
		{1 / sqrt2, -1 / sqrt2},
	}
}

func I() qGate {
	return [][]complex128{
		{1, 0},
		{0, 1},
	}
}

func X() qGate {
	return [][]complex128{
		{0, 1},
		{1, 0},
	}
}

func Y() qGate {
	return [][]complex128{
		{0, -1i},
		{1i, 0},
	}
}

func Z() qGate {
	return [][]complex128{
		{1, 0},
		{0, -1},
	}
}

func ControlledGate(m qGate, nQuBits int, control int, appliedQbit int) qGate {
	gate := In(int(math.Pow(2, float64(nQuBits))))
	//valueReference := int(math.Pow(2, float64(nQuBits-control)))
	dim := len(gate)
	q := newQubit(nQuBits)
	zeroIndexControl, _, _ := q.indexes(control)

	for i := 0; i < dim; i++ { //recorrido de fila
		//Si el qubit control es cero, no se hace nada
		if contains(zeroIndexControl, i) {
			continue
		}

		for j := 0; j < dim; j++ { //recorrido de la columna
			if contains(zeroIndexControl, j) {
				continue
			}

			diff := false
			for k := 0; k < nQuBits; k++ {
				auxFila := getNbit(i, nQuBits-k)
				auxColumna := getNbit(j, nQuBits-k)

				if appliedQbit != k && auxFila != auxColumna {
					diff = true
					break
				}
			}
			if !diff {
				r := getNbit(i, nQuBits-appliedQbit)
				c := getNbit(j, nQuBits-appliedQbit)
				gate[j][i] = m[r][c]
			}
		}
	}
	return gate
}

func In(n int) qGate {
	res := make([][]complex128, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]complex128, len(res))
		res[i][i] = 1
	}
	return res
}

func (q1 *qGate) tensorProduct(q2 *qGate) qGate {
	return qGate(matrix(*q1).tensorProduct(matrix(*q2)))
}

func getNbit(number int, bit int) int {
	for i := 0; i < bit-1; i++ {
		number /= 2
	}
	return number % 2
}
