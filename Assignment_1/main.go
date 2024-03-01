package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var daftarTeman = []Teman{
	{1, "TIARA RAHMANIA HADININGRUM", "Jakarta", "Software Engineer", "Mengagumi fitur concurrency di Go."},
	{2, "CHRISJHON", "Bandung", "Data Scientist", "Ingin mempelajari kemampuan pemrograman backend."},
	{3, "VIKY LORENT SEA PUTRA", "Surabaya", "UI/UX Designer", "Terpesona dengan performa tinggi Go dalam web development."},
	{4, "REFNI PASHA OLINA", "Yogyakarta", "Mobile App Developer", "Mendengar tentang kecepatan kompilasi Go."},
	{5, "Jeremia Letare Pane", "Semarang", "DevOps Engineer", "Menaruh minat pada keamanan built-in di Go."},
}

func main() {
	args := os.Args


	absen, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	teman := GetTemanByAbsen(absen)
	if teman != nil {
		fmt.Println("Nama:", teman.Nama)
		fmt.Println("Alamat:", teman.Alamat)
		fmt.Println("Pekerjaan:", teman.Pekerjaan)
		fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
	} else {
		fmt.Println("Data teman dengan absen", absen, "tidak ditemukan.")
	}
}

func GetTemanByAbsen(absen int) *Teman {
	for _, teman := range daftarTeman {
		if teman.Absen == absen {
			return &teman
		}
	}
	return nil
}
