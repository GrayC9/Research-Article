std::queue<int> q;
std::mutex mtx;
std::condition_variable cv;

void worker() {
    for (int i = 0; i < 5; ++i) {
        std::this_thread::sleep_for(std::chrono::seconds(1));
        std::unique_lock<std::mutex> lock(mtx);
        q.push(i);
        cv.notify_one();
    }
}

int main() {
    std::thread t(worker);

    for (int i = 0; i < 5; ++i) {
        std::unique_lock<std::mutex> lock(mtx);
        cv.wait(lock, []{ return !q.empty(); });
        std::cout << q.front() << std::endl;
        q.pop();
    }

    t.join();
    return 0;
}