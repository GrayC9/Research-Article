const int MAX_THREADS = 8;

std::vector<int> merge(const std::vector<int>& left, const std::vector<int>& right) {
    std::vector<int> result(left.size() + right.size());
    auto it_left = left.begin(), it_right = right.begin();
    auto it_result = result.begin();

    while (it_left != left.end() && it_right != right.end()) {
        if (*it_left < *it_right) {
            *it_result++ = *it_left++;
        } else {
            *it_result++ = *it_right++;
        }
    }
    std::copy(it_left, left.end(), it_result);
    std::copy(it_right, right.end(), it_result);

    return result;
}

std::vector<int> parallelMergeSort(std::vector<int>& arr, int depth) {
    if (arr.size() <= 1) {
        return arr;
    }

    if (depth <= 0) {
        std::sort(arr.begin(), arr.end());
        return arr;
    }

    size_t mid = arr.size() / 2;
    std::vector<int> left(arr.begin(), arr.begin() + mid);
    std::vector<int> right(arr.begin() + mid, arr.end());

    auto handle_left = std::async(std::launch::async, parallelMergeSort, std::ref(left), depth - 1);
    auto handle_right = std::async(std::launch::async, parallelMergeSort, std::ref(right), depth - 1);

    left = handle_left.get();
    right = handle_right.get();

    return merge(left, right);
}

int main() {
    std::vector<int> arr(1000000);
    std::generate(arr.begin(), arr.end(), rand);

    auto start_time = std::chrono::high_resolution_clock::now();

    std::vector<int> sorted_arr = parallelMergeSort(arr, MAX_THREADS);

    auto end_time = std::chrono::high_resolution_clock::now();
    std::chrono::duration<double, std::milli> duration = end_time - start_time;
    struct rusage usage;
    getrusage(RUSAGE_SELF, &usage);
    long memory_used = usage.ru_maxrss; // В килобайтах
    std::cout << "Time taken: " << duration.count() << " ms" << std::endl;
    std::cout << "Memory used: " << memory_used / 1024 << " MB" << std::endl;
    if (std::is_sorted(sorted_arr.begin(), sorted_arr.end())) {
        std::cout << "Array is sorted!" << std::endl;
    } else {
        std::cout << "Array is not sorted!" << std::endl;
    }

    return 0;
}