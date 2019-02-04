package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("You typed: %s\n", line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}
}

type Zone struct {
	ID    int
	Name  string
	Rooms []*Room
}
type Room struct {
	ID          int
	Zone        *Zone
	Name        string
	Description string
	Exits       [6]Exit
}

type Exit struct {
	To          *Room
	Description string
}

func loadRoom() {
	db, err := sql.Open("sqlite3", "./world.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
