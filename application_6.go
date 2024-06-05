func worker(ch chan int) {
	for {
		data := <-ch
		fmt.Println("Received:", data)
	}
}

func main() {
	ch := make(chan int)
	go worker(ch)

	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Sent:", i)
		time.Sleep(time.Second)
	}
}