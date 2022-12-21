package main

import (
	"fmt"
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
			fmt.Printf("FILA: valor de i : %v; COLUMNA: valor de j : %v\n", i, j)

			diff := false
			for k := 0; k < nQuBits; k++ {
				fmt.Printf("valor de k : %v;  valor de t : %v\n", k, appliedQbit)
				auxFila := getNbit(i, nQuBits-k)
				auxColumna := getNbit(j, nQuBits-k)
				fmt.Printf("auxFila: %v\n", auxFila)
				fmt.Printf("auxColumna: %v\n", auxColumna)

				if appliedQbit != k && auxFila != auxColumna {
					diff = true
					break
				}
			}
			if !diff {
				r := getNbit(i, nQuBits-appliedQbit)
				c := getNbit(j, nQuBits-appliedQbit)
				gate[j][i] = m[r][c]
				fmt.Printf("gate[%v][%v] = m[%v][%v] = %v\n", j, i, r, c, m[r][c])
			}
			print("\n\n\n")
		}
	}
	print(matrix(gate).toString())
	return gate
	/*return [][]complex128{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 1, 0},
	}*/
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
