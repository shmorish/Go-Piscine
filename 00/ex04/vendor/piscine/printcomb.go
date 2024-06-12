/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printcomb.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: morishi <morishi@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/11 17:51:31 by morishi           #+#    #+#             */
/*   Updated: 2024/06/11 17:51:37 by morishi          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func PrintComb() {
	for num1 := '0'; num1 <= '7'; num1++ {
		for num2 := num1 + 1; num2 <= '8'; num2++ {
			for num3 := num2 + 1; num3 <= '9'; num3++ {
				ft.PrintRune(num1)
				ft.PrintRune(num2)
				ft.PrintRune(num3)
				if num1 != '7' {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
			}
		}
	}
	ft.PrintRune('\n')
}
