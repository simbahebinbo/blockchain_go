package main

func main() {
	bc := NewBlockchain()
	defer bc.Close()

	cli := CLI{bc}
	cli.Run()
}
