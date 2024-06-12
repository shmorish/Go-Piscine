/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printcomb2.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: morishi <morishi@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 17:53:13 by morishi           #+#    #+#             */
/*   Updated: 2024/06/11 18:03:21 by morishi          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func printNumbers(n0, n1, n2, n3 rune) {
	ft.PrintRune(n0)
	ft.PrintRune(n1)
	ft.PrintRune(' ')
	ft.PrintRune(n2)
	ft.PrintRune(n3)
	if n0 == '9' && n1 == '8' && n2 == '9' && n3 == '9' {
		ft.PrintRune('\n')
	} else {
		ft.PrintRune(',')
		ft.PrintRune(' ')
	}
}

func PrintComb2() {
	for n0 := '0'; n0 <= '9'; n0++ {
		for n1 := '0'; n1 <= '9'; n1++ {
			for n2 := n0; n2 <= '9'; n2++ {
				for n3 := '0'; n3 <= '9'; n3++ {
					if (n0 == n2) && (n1 >= n3) {
						continue
					}
					printNumbers(n0, n1, n2, n3)
				}
			}
		}
	}
}
