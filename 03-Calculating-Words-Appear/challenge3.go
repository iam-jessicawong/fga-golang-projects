package main

import "fmt"

func calcword(words string) (wordcount map[string]int) {
	// ini bisa juga pake make terus gak perlu dibuat return gitu
	// wordcount = make(map[string]int)
	wordcount = map[string]int{}

	for _, char := range words {
		fmt.Printf("%c\n", char)
		// kalau looping string valuenya harus diconvert ke string dulu, karena yang dilooping valuenya jadi dalam bentuk rune (byte angka)
		wordcount[string(char)]++
	}

	return
}

func main() {
	fmt.Print(calcword("selamat malam"))
}
