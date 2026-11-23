package main

import "fmt"

func main3() {
var j, tarif int
fmt.Scan(&j)

tarif = 10000

if j > 5 {
tarif += 4000
}

fmt.Println(tarif)
}