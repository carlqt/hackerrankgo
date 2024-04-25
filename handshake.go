package hackerrankgo

// geeksforgeeks has a sample solution using recursion
// https://www.geeksforgeeks.org/number-of-handshakes-such-that-a-person-shakes-hands-only-once/
func Handshake(n int32) int32 {
	if n <= 1 {
		return 0
	}

	return n * (n - 1) / 2
}
