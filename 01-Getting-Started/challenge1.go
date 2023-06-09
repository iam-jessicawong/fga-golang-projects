package main

import "fmt"

func main() {
	i, base16 := 21, 15
	var k float64 = 123.456
	var j bool = true

	fmt.Printf("%v\n", i)        //menampilkan nilai i : 21
	fmt.Printf("%T\n", i)        //menampilkan tipe data dari variabel i
	fmt.Printf("%%\n")           //menampilkan tanda %
	fmt.Printf("%t\n", j)        //menampilkan nilai boolean j : true
	fmt.Printf("%t\n", !j)       //menampilkan nilai boolean j : false
	fmt.Printf("%c\n", '\u042F') //menampilkan unicode russia : Я (ya)
	fmt.Printf("%d\n", i)        //menampilkan nilai base 10 : 21
	fmt.Printf("%o\n", i)        //menampilkan nilai base 8 :25
	fmt.Printf("%x\n", base16)   //menampilkan nilai base 16 : f
	fmt.Printf("%X\n", base16)   //menampilkan nilai base 16 : F
	fmt.Printf("%U\n", 'Я')      //menampilkan unicode karakter Я : U+042F
	fmt.Printf("%f\n", k)        //menampilkan float : 123.456000
	fmt.Printf("%E\n", k)        //menampilkan float scientific : 1.234560E+02
}
