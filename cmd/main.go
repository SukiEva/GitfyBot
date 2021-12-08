package main

import (
	"GitfyBot/internal/db"
	"fmt"
)

func main() {
	db.AddUser("12356")
	user := db.GetUser("12356")
	fmt.Println(user.Id)
}
