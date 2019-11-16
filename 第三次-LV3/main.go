package main

func isprime(n int) {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			break
		}
		if i <= n/2 {
			println(n)
		}
	}
}

func main() {
	for i := 2; i <= 10000; i++ {
		go isprime(i)
	}
}
