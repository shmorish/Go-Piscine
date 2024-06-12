/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printcombn.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: morishi <morishi@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 19:17:37 by morishi           #+#    #+#             */
/*   Updated: 2024/06/11 19:22:37 by morishi          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func PrintNumber(n int) {
	ft.PrintRune(rune(n) + '0')
}

func printCombination(digits [10]int, n int) {
	for i := 0; i < n; i++ {
		PrintNumber(digits[i])
	}
	if digits[0] != 9-(n-1) {
		ft.PrintRune(',')
		ft.PrintRune(' ')
	}
}

func backtrack(digits [10]int, index, start, n int) {
	if index == n {
		printCombination(digits, n)
		return
	}
	for i := start; i <= 9-(n-index-1); i++ {
		digits[index] = i
		backtrack(digits, index+1, i+1, n)
	}
}

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	var digits [10]int
	backtrack(digits, 0, 0, n)
	ft.PrintRune('\n')
}
