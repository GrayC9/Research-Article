const (
	maxGoroutines = 16
  )
   
  func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j := 0, 0
	for k := 0; k < len(result); k++ {
	  if i >= len(left) {
		result[k] = right[j]
		j++
	  } else if j >= len(right) {
		result[k] = left[i]
		i++
	  } else if left[i] < right[j] {
		result[k] = left[i]
		i++
	  } else {
		result[k] = right[j]
		j++
	  }
	}
	return result
  }
   
  func parallelMergeSort(arr []int, maxDepth int, sem chan struct{}, wg *sync.WaitGroup) []int {
	defer wg.Done()
   
	if len(arr) <= 1 {
	  return arr
	}
   
	if maxDepth <= 0 {
	  return mergeSort(arr)
	}
   
	mid := len(arr) / 2
	var left, right []int
	var leftWg, rightWg sync.WaitGroup
   
	sem <- struct{}{}
	leftWg.Add(1)
	go func() {
	  defer leftWg.Done()
	  left = parallelMergeSort(arr[:mid], maxDepth-1, sem, &leftWg)
	  <-sem 
	}()
   
	sem <- struct{}{}
	rightWg.Add(1)
	go func() {
	  defer rightWg.Done()
	  right = parallelMergeSort(arr[mid:], maxDepth-1, sem, &rightWg)
	  <-sem 
	}()
   
	leftWg.Wait()
	rightWg.Wait()
   
	return merge(left, right)
  }
   
  func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
	  return arr
	}
   
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
   
	return merge(left, right)
  }
   
  func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
   
	arr := make([]int, 1000000)
	for i := range arr {
	  arr[i] = rand.Intn(1000000)
	}
   
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	startMem := memStats.Alloc
	startTime := time.Now()
   
	sem := make(chan struct{}, maxGoroutines)
	var wg sync.WaitGroup
	wg.Add(1)
	sortedArr := parallelMergeSort(arr, runtime.NumCPU(), sem, &wg)
	wg.Wait()
   
	duration := time.Since(startTime)
	runtime.ReadMemStats(&memStats)
	endMem := memStats.Alloc
   
	fmt.Printf("Time потрачено: %v ms\n", duration.Milliseconds())
	fmt.Printf("Memory used: %v MB\n", (endMem-startMem)/1024/1024)
   
	for i := 1; i < len(sortedArr); i++ {
	  if sortedArr[i-1] > sortedArr[i]) {
		fmt.Println("Массив не отсортирован!")
		return
	  }
	}
	fmt.Println("Отсортирован!")
  }