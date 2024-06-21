package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

// FindNextPrime is a function that returns the next prime number after the int passed as parameter.
func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	for i := nb; ; i++ {
		if IsPrime(i) {
			return i
		}
	}
}
