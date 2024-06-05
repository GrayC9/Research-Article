public class ThreadCreationTimeTest {

    public static void main(String[] args) {
        int numThreads = 10000;
        List<Thread> threads = new ArrayList<>(numThreads);

        long start = System.currentTimeMillis();

        for (int i = 0; i < numThreads; i++) {
            Thread thread = new Thread(() -> {
                // Empty task
            });
            threads.add(thread);
            thread.start();
        }

        for (Thread thread : threads) {
            try {
                thread.join();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        long end = System.currentTimeMillis();
        long duration = end - start;

        System.out.println("Time taken to create " + numThreads + " threads: " + duration + " milliseconds");
    }
}