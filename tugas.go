package main

import (
	"fmt"
	"os"
	"strconv"
)


type Student struct {
	Nama string
	Alamat string
	Pekerjaan string
	Alasan_memilih_Golang string	
}

func main()  {
	Pelajar := []Student{
		{
			Nama: "Bachtiarudin",
			Alamat: "Kota Serang",
			Pekerjaan: "Mahasiswa",
			Alasan_memilih_Golang: "Golang Bahasa Yang Cepat Dan Ringan",
		},
		{
			Nama: "Abdian Rizki Ramanadhan",
			Alamat: "Kota Tangerang",
			Pekerjaan: "Mahasiswa",
			Alasan_memilih_Golang: "Golang Itu Seru",
		},
		{
			Nama: "Chandra Adhimulya",
			Alamat: "Jakarta",
			Pekerjaan: "Wiarswasta",
			Alasan_memilih_Golang: "Golang itu bahasa yang di kembangkan google",
		},
		{
			Nama: "Indra Wijaya",
			Alamat: "Surabaya",
			Pekerjaan: "Programmer",
			Alasan_memilih_Golang: "Golang bahasa Back End yang enak digunakan",
		},
		{
			Nama: "Muhammad Akbar Hamid",
			Alamat: "Cilegon",
			Pekerjaan: "Mahasiswa",
			Alasan_memilih_Golang: "Golang Sebagai Project Tugas Akhir saya",
		},
		{
			Nama: "Ricky Romansyah",
			Alamat: "Pandeglang",
			Pekerjaan: "Mahasiswa",
			Alasan_memilih_Golang: "Ingin membuat aplikasi dengan golang",
		},
		{
			Nama: "Agus Rudianto",
			Alamat: "Karawang",
			Pekerjaan: "Wiraswasta",
			Alasan_memilih_Golang: "Ingn berkarir sebagai programmer golang",
		},
		{
			Nama: "Aldi Tarigan",
			Alamat: "Jakarat Selatan",
			Pekerjaan: "Mahasiswa",
			Alasan_memilih_Golang: "Mencoba bahasa back end golang",
		},
		{
			Nama: "Dimas Zubaleta",
			Alamat: "Tangrang Selatan",
			Pekerjaan: "Mahasiswa",
			Alasan_memilih_Golang: "Ingin membuat aplikasi microservices dengan golang",
		},
		{
			Nama: "I Made Rismawan Nugraha",
			Alamat: "Denpasar",
			Pekerjaan: "Mahasiswa",
			Alasan_memilih_Golang: "Ingin belajar dan membangun aplikasi dengan beck end golang",
		},
	}

	absen := os.Args[1]

	for i, v := range Pelajar {
		index := strconv.Itoa(i)
		if absen == index {
			fmt.Println("Nama                         :", v.Nama)
			fmt.Println("Alamat                       :", v.Alamat)
			fmt.Println("Pekerjaan                    :", v.Pekerjaan)
			fmt.Println("Alasan Memilih Kelas Golang  :", v.Alasan_memilih_Golang)
		}
	}
}