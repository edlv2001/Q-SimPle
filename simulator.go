package main

import (
	"fmt"
	"math"
)

type qSim struct {
	qubits  *qubit
	actions []action
}

func (c *qSim) simulate() {
	for i := range c.actions {
		c.executeAction(&c.actions[i])
	}
}

func new(n int) qSim {
	aux := newQubit(n)
	return qSim{
		qubits:  &aux,
		actions: make([]action, 0),
	}
}

type qGateSim struct {
	gate          qGate
	target        []int
	controlQubits []int
}

func (q *qSim) measure(qb ...int) *qubit {
	if len(qb) == 0 {
		qb = make([]int, q.nQubit())
		for i := 0; i < len(qb); i++ {
			qb[i] = i + 1
		}
	}
	res := q.qubits.measure(qb[0] - 1)
	for i := 1; i < len(qb); i++ {
		res = &qubit{
			res.vector.tensorVectorProduct(q.qubits.measure(qb[i] - 1).vector),
		}
	}

	return res

}

func (q *qSim) nQubit() int {
	return int(math.Log2(float64(len(q.qubits.vector))))
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

	//Aplicar estado
	aux := &qubit{
		vector: multiply(matrix(g), q.qubits.vector),
	}
	q.qubits = aux
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
	str := "\n\n\nState: \n\n"
	for i := 0; i < len(q.qubits.vector); i++ {
		str += fmt.Sprint("\t", i, ":  ", q.qubits.vector[i][0], "\n")
	}
	str += "\n\n"
	print(str)
}

func (q *qSim) count(times int) []int {
	res := make([]int, len(q.qubits.vector))
	for i := 0; i < times; i++ {
		qbitClone := q.qubits.clone()

		for j := 0; j < qbitClone.nQubit(); j++ {
			qbitClone.measure(j)
		}
		for j := 0; j < len(q.qubits.vector); j++ {
			if qbitClone.vector[j][0] == 1 {
				res[j]++
			}
		}

	}

	return res

}

/*
	func (q *qSim) addGate(gate qGate, qubits ...int) {
		aux := qGateSim{
			gate:   gate,
			target: qubits,
		}
		q.gates = append(q.gates, &aux)
	}

func (q *qSim) addGateFunction(gateF *func(m qGate, q ...int), qubit int) {

}

func (q *qSim) addControlledGate(gate qGate) {

}
*/
type gateSequence []*qGate
type controlledGate qGate
type measureGate []int

func newControlledGate(m qGate, nQuBits int, control int, appliedQbit int) controlledGate {
	aux := ControlledGate(m, nQuBits, control, appliedQbit)
	return controlledGate(aux)
}

func newGateSequence(size int) gateSequence {
	res := make(gateSequence, size)
	for i := 0; i < size; i++ {
		res[i] = nil
	}
	return res
}

const (
	optimized = iota
	debug     = iota
)

func (c controlledGate) obtainMatrix() *qGate {
	aux := qGate(c)
	aux2 := &aux
	return aux2
}

func (gSequence gateSequence) obtainMatrix() *qGate {
	if len(gSequence) == 0 {
		return nil
	}
	identidad := I()
	var res *qGate
	if gSequence[0] == nil {
		res = &identidad
	} else {
		res = gSequence[0]
	}
	var aux qGate
	for i := 1; i < len(gSequence); i++ {

		if gSequence[i] == nil {
			aux = res.tensorProduct(&identidad)
		} else {
			aux = res.tensorProduct(gSequence[i])
		}
		res = &aux
	}

	return res
}

type action struct {
	gateSequence   *gateSequence
	controlledGate *controlledGate
	measureGate    *[]int
	show           bool
}

const (
	unitaryGate = iota
	cGate       = iota
	measurement = iota
	showState   = iota
)

func newActionUnitary(size int) action {
	aux := newGateSequence(size)
	return action{
		gateSequence:   &aux,
		controlledGate: nil,
		measureGate:    nil,
		show:           false,
	}
}

func newActionControlled(m qGate, nQuBits int, control int, appliedQbit int) action {
	aux := newControlledGate(m, nQuBits, control, appliedQbit)
	return action{
		gateSequence:   nil,
		controlledGate: &aux,
		measureGate:    nil,
		show:           false,
	}
}

func newActionMeasure(qubits ...int) action {
	aux := qubits
	return action{
		gateSequence:   nil,
		controlledGate: nil,
		measureGate:    &aux,
		show:           false,
	}
}

func newActionShow() action {
	return action{
		gateSequence:   nil,
		controlledGate: nil,
		measureGate:    nil,
		show:           true,
	}
}

func (q *qSim) executeAction(a *action) {
	if a == nil {
		return
	}
	if a.controlledGate != nil {
		aux := *a.controlledGate.obtainMatrix()
		q.qubits.vector = multiply(matrix(aux), q.qubits.vector)
	} else if a.gateSequence != nil {
		aux := *a.gateSequence.obtainMatrix()
		q.qubits.vector = multiply(matrix(aux), q.qubits.vector)
	} else if a.measureGate != nil {
		q.measureN(*a.measureGate)
	} else if a.show {
		q.printStateVector()
	}
}

func (q *qSim) measureN(index []int) {
	if len(index) == 0 {
		for i := 0; i < q.nQubit(); i++ {
			q.qubits.measure(i)
		}
	}
	for _, i := range index {
		q.qubits.measure(i)
	}
}

func (gSequence gateSequence) addGate(gate *qGate, qbits ...int) {
	for _, i := range qbits {
		if gSequence[i] == nil {
			gSequence[i] = gate
		} else {
			aux := qGate(multiply(matrix(*gSequence[i]), matrix(*gate)))
			gSequence[i] = &aux
		}
	}
}

func (q *qSim) addUnitaryGate(gate qGate, index ...int) {
	if len(q.actions) == 0 || q.actions[len(q.actions)-1].gateSequence == nil {
		q.actions = append(q.actions, newActionUnitary(q.nQubit()))
	}
	q.actions[len(q.actions)-1].gateSequence.addGate(&gate, index...)
}

func (q *qSim) addControlledGate(gate qGate, control int, target int) {
	q.actions = append(q.actions, newActionControlled(gate, q.nQubit(), control, target))
}

func (q *qSim) MEASURE(index ...int) {
	q.actions = append(q.actions, newActionMeasure(index...))
}

func (q *qSim) SHOW() {
	q.actions = append(q.actions, newActionShow())
}

func (q *qSim) NOT(index ...int) {
	q.addUnitaryGate(X(), index...)
}

func (q *qSim) H(index ...int) {
	q.addUnitaryGate(H(), index...)
}

func (q *qSim) CNOT(control int, target int) {
	q.addControlledGate(X(), control, target)
}

func (q *qSim) CH(control int, target int) {
	q.addControlledGate(H(), control, target)
}

func (q *qSim) SWAP(qubit1 int, qubit2 int) {
	q.addControlledGate(X(), qubit1, qubit2)
	q.addControlledGate(X(), qubit2, qubit1)
	q.addControlledGate(X(), qubit1, qubit2)
}
