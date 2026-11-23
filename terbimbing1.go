package main

import "fmt"

func main4() {
	var s, p int

	fmt.Scan(&s, &p)

	if s >= 0 && s <= 1000000 && p >= 0 && p <= 1000000 {
		if s > p {
			fmt.Println("Pembayaran Berhasil")
		} else if s == p {
			fmt.Println("Saldo sangat pas")
		} else {
			fmt.Println("Saldo tidak cukup")
		}
	}
}