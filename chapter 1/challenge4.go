package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	p := []Person{
		{
			Nama:      "Adi",
			Alamat:    "Jl. Telekomunikasi",
			Pekerjaan: "IT",
			Alasan:    "Suka kode",
		}, {
			Nama:      "Whydi",
			Alamat:    "Jl. Ciganitri",
			Pekerjaan: "Finance",
			Alasan:    "Suka angka",
		}, {
			Nama:      "Wyn",
			Alamat:    "Jl. Sukapura",
			Pekerjaan: "Guru",
			Alasan:    "Suka mengajar",
		}, {
			Nama:      "Ikr",
			Alamat:    "Jl. Ciganitri II",
			Pekerjaan: "Atlet",
			Alasan:    "Suka bola",
		}, {
			Nama:      "Mir",
			Alamat:    "Jl. Cipagalo",
			Pekerjaan: "Tukang",
			Alasan:    "Suka ngoprek",
		}, {
			Nama:      "Ghif",
			Alamat:    "Jl. Mangga Dua",
			Pekerjaan: "Game",
			Alasan:    "Suka main",
		}, {
			Nama:      "Hazm",
			Alamat:    "Jl. Sukabirus",
			Pekerjaan: "Kepala Sekolah",
			Alasan:    "Suka menasihat",
		}, {
			Nama:      "Jek",
			Alamat:    "Jl. Terusan",
			Pekerjaan: "Pedagang",
			Alasan:    "Suka dimsum",
		}, {
			Nama:      "Ivn",
			Alamat:    "Jl. Sadink",
			Pekerjaan: "FE",
			Alasan:    "Suka slicing",
		}, {
			Nama:      "Azhr",
			Alamat:    "Jl. Batununggal",
			Pekerjaan: "FE",
			Alasan:    "Suka chill",
		},
	}
	absen, err := strconv.Atoi(os.Args[1])
	cariBiodata(p, absen, err)

}

func cariBiodata(person []Person, absen int, err error) {
	if err == nil {
		if absen > 0 && absen <= len(person) {
			fmt.Printf("Absen:\t\t%d\n", absen)
			absen = absen - 1
			fmt.Printf("Nama:\t\t%s\n", person[absen].Nama)
			fmt.Printf("Alamat:\t\t%s\n", person[absen].Alamat)
			fmt.Printf("Pekerjaan:\t%s\n", person[absen].Pekerjaan)
			fmt.Printf("Alasan:\t\t%s\n", person[absen].Alasan)
		} else {
			fmt.Printf("Tidak ada data dengan absen %d.", absen)
		}
	} else {
		fmt.Println("Hanya menerima arguments berupa angka.")
	}
}
