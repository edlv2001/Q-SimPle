package main

func main() {
	q := new(2)

	q.H(0, 1)
	q.NOT(1)
	q.SHOW()
	q.SWAP(0, 1)
	q.SHOW()
	q.simulate()

}

func printInit() {
	print("░██████╗░  ░██████╗██╗███╗░░░███╗░░░░░░██████╗░██╗░░░░░███████╗\n██╔═══██╗  ██╔════╝██║████╗░████║░░░░░░██╔══██╗██║░░░░░██╔════╝\n██║██╗██║  ╚█████╗░██║██╔████╔██║█████╗██████╔╝██║░░░░░█████╗░░\n╚██████╔╝  ░╚═══██╗██║██║╚██╔╝██║╚════╝██╔═══╝░██║░░░░░██╔══╝░░\n░╚═██╔═╝░  ██████╔╝██║██║░╚═╝░██║░░░░░░██║░░░░░███████╗███████╗\n░░░╚═╝░░░  ╚═════╝░╚═╝╚═╝░░░░░╚═╝░░░░░░╚═╝░░░░░╚══════╝╚══════╝")
	print("\n                               by Eduardo de la Vega Fernández\n")
}
