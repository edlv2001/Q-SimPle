package main

import (
	"fmt"
	"math"
)

type qSim struct {
	qbit  []*qubit
	gates []*qGateSim
}

func (c *qSim) simulate() {
	for i := range c.gates {
		(*c.gates[i].gate)(c.gates[i].args...)
	}
}

func new(n int) qSim {
	return qSim{
		qbit:  make([]*qubit, n),
		gates: make([]*qGateSim, n),
	}
}

type qGateSim struct {
	gate *func(q ...int)
	args []int
}

func (q *qSim) measure(qb ...int) *qubit {
	if len(qb) == 0 {
		qb = make([]int, q.nQubit())
		for i := 0; i < len(qb); i++ {
			qb[i] = i + 1
		}
		/*
			n := q.nQubit()

			//m := make([]*qubit, 0, n)
			res := q.qbit[0].measure(0)
			for i := 1; i < n; i++ {
				//m = append(m, q.qbit[0].measure(i))
				res = &qubit{
					res.vector.tensorVectorProduct(q.qbit[0].measure(i).vector),
				}
			}
			//q.qbit[0] = q.qbit[0].vector.tensorVectorProduct(res)
			return res
		*/
	}
	//m := make([]*qubit, 0, len(qb))
	res := q.qbit[0].measure(qb[0] - 1)
	for i := 1; i < len(qb); i++ {
		//m = append(m, q.qbit)
		//q.qbit[i].measure()
		res = &qubit{
			res.vector.tensorVectorProduct(q.qbit[0].measure(qb[i] - 1).vector),
		}
	}
	//print(len(res.vector))
	//print("Res: \n" + res.vector.toString())

	return res

}

func (q *qSim) nQubit() int {
	return int(math.Log2(float64(len(q.qbit[0].vector))))
}

func (q *qSim) execGate(m qGate, index ...int) *qSim {
	var g qGate
	id := I()

	//Construir matriz a aplicar al estado
	if contains(index, 1) {
		g = m
	} else {
		g = I()
	}

	for i := 2; i < q.nQubit()+1; i++ {
		if contains(index, i) {
			g = g.tensorProduct(&m)
		} else {
			g = g.tensorProduct(&id)
		}
	}
	//print(matrix(g).toString())
	//print(q.qbit[0].vector.toString())
	//fmt.Println(multiply(matrix(g), q.qbit[0].vector).toString())

	//Aplicar estado
	aux := &qubit{
		vector: multiply(matrix(g), q.qbit[0].vector),
	}
	q.qbit[0] = aux
	//print("\n\n\n")
	//print(q.qbit[0].vector.toString())
	return q
}

func contains(list []int, n int) bool {
	for _, val := range list {
		if val == n {
			return true
		}
	}
	return false
}

func (q *qSim) printStateVector() {
	str := "\n\n\nQ State: \n\n"
	for i := 0; i < len(q.qbit[0].vector); i++ {
		str += fmt.Sprint("\t", i, ":  ", q.qbit[0].vector[i][0], "\n")
	}
	str += "\n\n"
	print(str)
}

func (q *qSim) count(times int) []int {
	res := make([]int, len(q.qbit[0].vector))
	for i := 0; i < times; i++ {
		qbitClone := q.qbit[0].clone()

		for j := 0; j < qbitClone.nQubit(); j++ {
			qbitClone.measure(j)
		}
		for j := 0; j < len(q.qbit[0].vector); j++ {
			if qbitClone.vector[j][0] == 1 {
				res[j]++
			}
		}

	}

	return res

}
