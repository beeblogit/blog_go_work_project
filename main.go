package main

import (
	"encoding/json"
	"fmt"
	"github.com/beeblogit/blog_go_work_pkg/domain"
)

func main() {
	u := domain.User{
		ID:        "1231",
		FirstName: "Nahuel",
		LastName:  "Costamagna",
		Country:   "Argentina",
	}

	value, _ := json.Marshal(u)
	fmt.Println(string(value))
}
