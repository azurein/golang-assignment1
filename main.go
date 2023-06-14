package main

import (
	"fmt"
	"os"
)

type User struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	users := []User{
		{nama: "robin", alamat: "indonesia", pekerjaan: "backend", alasan: "kritis"},
		{nama: "merry", alamat: "korea", pekerjaan: "tester", alasan: "teliti"},
		{nama: "jacky", alamat: "china", pekerjaan: "frontend", alasan: "estetik"},
	}
	args := os.Args[1]
	for _, user := range users {
		if user.nama == args {
			fmt.Println(user)
			return
		}
	}
	fmt.Println("not found")
}
