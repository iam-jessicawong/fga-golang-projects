package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func (b Biodata) Introduce() {
	fmt.Printf("Nama: %v\nAlamat: %v\nPekerjaan: %v\nAlasan memilih kelas Golang: %v", b.nama, b.alamat, b.pekerjaan, b.alasan)
}

var people = []Biodata{
	{
		nama:      "jessica",
		alamat:    "tapaktuan",
		pekerjaan: "web developer",
		alasan:    "menarik dan banyak dicari",
	},
	{
		nama:      "ayu",
		alamat:    "jakarta",
		pekerjaan: "it support",
		alasan:    "ingin belajar hal baru",
	},
}

func main() {
	absen_no := os.Args[1:] //take arguments

	// validate argument
	if len(absen_no) == 0 {
		fmt.Printf("please give one argument from 1 - %d\n", len(people))
	} else {
		val, err := strconv.Atoi(absen_no[0])            //convert to number
		if err != nil || val > len(people) || val <= 0 { //check if the argument is valid number
			fmt.Printf("please input an argument from 1 - %d\n", len(people))
		} else {
			people[val-1].Introduce() //access and print data
		}
	}
}
