package piscine

import "ft"

const N = 8

var position = [N]int{}

func PrintString(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func PrintInt(n int) {
	if n < 0 {
		ft.PrintRune('-')
		n *= -1
	}
	if n > 9 {
		PrintInt(n / 10)
	}
	ft.PrintRune(rune(n%10) + '0')
}

func isSafe(queen_number, row_position int) bool {
	for i := 0; i < queen_number; i++ {
		other_row_pos := position[i]

		if other_row_pos == row_position || other_row_pos == row_position-(queen_number-i) || other_row_pos == row_position+(queen_number-i) {
			return false
		}
	}
	return true
}

func solve(k int) {
	if k == N {
		for i := 0; i < N; i++ {
			PrintInt(position[i] + 1)
		}
		PrintString("\n")
	} else {
		for i := 0; i < N; i++ {
			if isSafe(k, i) {
				position[k] = i
				solve(k + 1)
			}
		}
	}
}

func EightQueens() {
	solve(0)
}
