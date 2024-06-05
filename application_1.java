public class ParallelMergeSort {

    private static final int MAX_THREADS = 8;
    private static final ForkJoinPool pool = new ForkJoinPool(MAX_THREADS);

    public static void main(String[] args) {
        int[] arr = new int[1000000];
        Random random = new Random();
        for (int i = 0; i < arr.length; i++) {
            arr[i] = random.nextInt(1000000);
        }

        long startTime = System.currentTimeMillis();

        int[] sortedArr = parallelMergeSort(arr);
        long endTime = System.currentTimeMillis();

        Runtime runtime = Runtime.getRuntime();
        long memoryUsed = (runtime.totalMemory() - runtime.freeMemory()) / (1024 * 1024);

        System.out.println("Time taken: " + (endTime - startTime) + " ms");
        System.out.println("Memory used: " + memoryUsed + " MB");

        if (isSorted(sortedArr)) {
            System.out.println("Array is sorted!");
        } else {
            System.out.println("Array is not sorted!");
        }
    }

    public static int[] parallelMergeSort(int[] arr) {
        return pool.invoke(new MergeSortTask(arr));
    }

    private static class MergeSortTask extends RecursiveTask<int[]> {
        private final int[] arr;

        public MergeSortTask(int[] arr) {
            this.arr = arr;
        }

        @Override
        protected int[] compute() {
            if (arr.length <= 1) {
                return arr;
            }

            int mid = arr.length / 2;

            MergeSortTask leftTask = new MergeSortTask(Arrays.copyOfRange(arr, 0, mid));
            MergeSortTask rightTask = new MergeSortTask(Arrays.copyOfRange(arr, mid, arr.length));

            invokeAll(leftTask, rightTask);

            int[] left = leftTask.join();
            int[] right = rightTask.join();

            return merge(left, right);
        }

        private int[] merge(int[] left, int[] right) {
            int[] result = new int[left.length + right.length];
            int i = 0, j = 0, k = 0;

            while (i < left.length && j < right.length) {
                if (left[i] <= right[j]) {
                    result[k++] = left[i++];
                } else {
                    result[k++] = right[j++];
                }
            }

            while (i < left.length) {
                result[k++] = left[i++];
            }

            while (j < right.length) {
                result[k++] = right[j++];
            }

            return result;
        }
    }

    private static boolean isSorted(int[] arr) {
        for (int i = 1; i < arr.length; i++) {
            if (arr[i - 1] > arr[i]) {
                return false;
            }
        }
        return true;
    }
}