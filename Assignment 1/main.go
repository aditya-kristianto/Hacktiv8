package main

import (
	"fmt"
	"os"
	"strconv"
)

type friend struct {
	name    string
	address string
	job     string
	reason  string
}

func main() {
	friends := []friend{
		{
			name:    "clara",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "fiqri",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "medy",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "lutfi",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "raiza",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "ian",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "umar",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "fazry",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "pandi",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "tantut",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		},
	}

	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	index = index - 1

	if index < 0 || index > len(friends)-1 {
		fmt.Println("Tidak berhasil menemukan data teman.")
	} else {
		fmt.Println("Nama : " + friends[index].name)
		fmt.Println("Alamat : " + friends[index].address)
		fmt.Println("Pekerjaan : " + friends[index].job)
		fmt.Println("Alasan memilih kelas Golang : " + friends[index].reason)
	}
}
