package main

func main() {
	q := new(2)

	q.CH(0, 1)
	q.CH(1, 0)

	q.simulate()

}

func printInit() {
	print("░██████╗░  ░██████╗██╗███╗░░░███╗░░░░░░██████╗░██╗░░░░░███████╗\n██╔═══██╗  ██╔════╝██║████╗░████║░░░░░░██╔══██╗██║░░░░░██╔════╝\n██║██╗██║  ╚█████╗░██║██╔████╔██║█████╗██████╔╝██║░░░░░█████╗░░\n╚██████╔╝  ░╚═══██╗██║██║╚██╔╝██║╚════╝██╔═══╝░██║░░░░░██╔══╝░░\n░╚═██╔═╝░  ██████╔╝██║██║░╚═╝░██║░░░░░░██║░░░░░███████╗███████╗\n░░░╚═╝░░░  ╚═════╝░╚═╝╚═╝░░░░░╚═╝░░░░░░╚═╝░░░░░╚══════╝╚══════╝")
	print("\n                               by Eduardo de la Vega Fernández\n")
}
