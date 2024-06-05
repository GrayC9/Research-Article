const (
	matrixSize    = 1000
	numGoroutines = 8
)

func multiplyMatrices(A, B, C [][]int, startRow, endRow int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := startRow; i < endRow; i++ {
		for j := 0; j < matrixSize; j++ {
			sum := 0
			for k := 0; k < matrixSize; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}
}

func createMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(100)
		}
	}
	return matrix
}

func main() {
	rand.Seed(time.Now().UnixNano())
	A := createMatrix(matrixSize)
	B := createMatrix(matrixSize)
	C := make([][]int, matrixSize)
	for i := range C {
		C[i] = make([]int, matrixSize)
	}

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	initialAlloc := memStats.TotalAlloc

	start := time.Now()

	var wg sync.WaitGroup
	rowsPerGoroutine := matrixSize / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		startRow := i * rowsPerGoroutine
		endRow := startRow + rowsPerGoroutine
		if i == numGoroutines-1 {
			endRow = matrixSize
		}
		wg.Add(1)
		go multiplyMatrices(A, B, C, startRow, endRow, &wg)
	}

	wg.Wait()

	duration := time.Since(start)
	runtime.ReadMemStats(&memStats)
	finalAlloc := memStats.TotalAlloc

	fmt.Printf("Matrix multiplication with %d goroutines\n", numGoroutines)
	fmt.Printf("Time taken: %d ms\n", duration.Milliseconds())
	fmt.Printf("Memory used: %.2f MB\n", float64(finalAlloc-initialAlloc)/1024/1024)
}