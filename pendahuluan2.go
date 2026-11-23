package main

import "fmt"

func main1() {
var t, h, k int
var hasil string
fmt.Scan(&t, &h, &k)

hasil = "TIDAK BOLEH UAS"

if h*100 >= 75*t && k >= 8 {
hasil = "BOLEH UAS"
}

fmt.Println(hasil)
}