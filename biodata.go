package main

import (
	"fmt"
	"os"
	"strconv"
)

type friend struct {
	Id        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var friends []friend

func init() {
	friends = append(friends, friend{1, "Kevin", "makassar", "programmer", "Belajar IT mengasikkan"})
	friends = append(friends, friend{2, "Dylan", "Gowa", "System Administrator", "Membangun Infrastructure yang durable challenging"})
	friends = append(friends, friend{3, "Prayudi", "Pare-pare", "Devops", "Tidak harus selalu berhubungan dengan mesin dalam pekerjaan"})
}

func validator(s []string) (a int, e error) {
	if len(s) != 2 {
		return 0, fmt.Errorf("Please enter the number, or don't use a space between two numbers.")
	}
	a, e = strconv.Atoi(s[1])
	if e == nil {
		return
	}
	e = fmt.Errorf("%s isn't number", s[1])
	return
}

func findFriend(f []friend, i int, callback func(friend)) {
	var friend friend
	for _, value := range f {
		if int(value.Id) == i {
			friend = value
		}
	}

	callback(friend)
}

func print(f friend) {
	if (f == friend{}) {
		fmt.Println("Friend not found")
		os.Exit(0)
	}

	string := `
	Nama : %s
	Pekerjaan: %s
	Alamat: %s
	Alasan: %s
	`
	fmt.Printf(string, f.Nama, f.Pekerjaan, f.Alamat, f.Alasan)
}

func main() {
	arg, err := validator(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	findFriend(friends, arg, print)
}
