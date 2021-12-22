package main

func main() {
	// ch is goroutine-safe
	// stores up to capacity elements,(存储到容量大小)
	// and provides FIFO semantics(提供FIFO语义)
	// sends values between goroutines
	// can cause goroutines block or unblock
	//ch := make(chan Task, 3)
}
