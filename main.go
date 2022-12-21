package main

import (
	"math/rand"
	"time"
)

func main() {
	printInit()

	rand.Seed(time.Now().UnixNano())
	//v2 := newVectorVert(4, 5, 6, 1i)
	//m := newMatrixValue([]complex128{1, 0, 0, 0}, []complex128{0, 1, 0, 0}, []complex128{0, 0, 0, 1}, []complex128{0, 0, 1, 0})
	//v3 := multiply(m, v2)
	q := new(1)
	qbit := newQubit(3)
	q.qbit[0] = &qbit

	//q.execGate(H(), 1, 2, 3)

	//q.printStateVector()
	q.execGate(H(), 1)
	//fmt.Println(q.qbit[0].vector.toString())
	q.printStateVector()
	q.qbit[0].vector = multiply(matrix(ControlledGate(H(), 3, 0, 1)), q.qbit[0].vector)

	//fmt.Printf("%v\n", q.count(500))
	//q.measure(1, 2)

	//q.measure(2)
	//fmt.Println(q.qbit[0].vector.toString())
	q.printStateVector()
}

func printInit() {
	print("░██████╗░  ░██████╗██╗███╗░░░███╗░░░░░░██████╗░██╗░░░░░███████╗\n██╔═══██╗  ██╔════╝██║████╗░████║░░░░░░██╔══██╗██║░░░░░██╔════╝\n██║██╗██║  ╚█████╗░██║██╔████╔██║█████╗██████╔╝██║░░░░░█████╗░░\n╚██████╔╝  ░╚═══██╗██║██║╚██╔╝██║╚════╝██╔═══╝░██║░░░░░██╔══╝░░\n░╚═██╔═╝░  ██████╔╝██║██║░╚═╝░██║░░░░░░██║░░░░░███████╗███████╗\n░░░╚═╝░░░  ╚═════╝░╚═╝╚═╝░░░░░╚═╝░░░░░░╚═╝░░░░░╚══════╝╚══════╝")
	print("\n                               by Eduardo de la Vega Fernández\n")
}
