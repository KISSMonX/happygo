package main

func main() {
	s := []string{"1", "2", "4", "6"}
	c := make(chan string, 10)

	c <- s[0]
}
