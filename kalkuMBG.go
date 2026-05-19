package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nilai int64
	var konversi string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("SELAMAT DATANG DI KALKULATOR MBG")
	fmt.Println("Setiap harinya MBG menghabiskan 1,2 Triliyun (Rp1.200.000.000.000)")

	fmt.Print("APA YANG INGIN ANDA KONVERSI\t\t: ")
	fmt.Fscanln(reader, &konversi)

	fmt.Printf("Berapa nilainya %s Anda ?\t\t: ", konversi)
	fmt.Scanln(&nilai)

	hMBG := float64(nilai) / 1200000000000.0

	if hMBG/7 >= 1 {
		fmt.Printf("%s Anda setara dengan %.2f minggu MBG\n", konversi, hMBG/7)
	} else if hMBG >= 1 {
		fmt.Printf("%s Anda setara dengan %.2f hari MBG\n", konversi, hMBG)
	} else if hMBG*24 >= 1 {
		fmt.Printf("%s Anda setara dengan %.2f jam MBG\n", konversi, hMBG*24)
	} else if hMBG*1440 >= 1 {
		fmt.Printf("%s Anda setara dengan %.2f menit MBG\n", konversi, hMBG*1440)
	} else {
		fmt.Printf("%s Anda setara dengan %.2f detik MBG\n", konversi, hMBG*86400)
	}
}
