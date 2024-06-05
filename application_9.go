func main() {
	var wg sync.WaitGroup
	numGoroutines := 10000
	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Time taken to create %d goroutines: %v\n", numGoroutines, duration)
}