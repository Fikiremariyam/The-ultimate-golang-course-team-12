/*
	ultimate golang course  week one assignment

Title :- todo list
approach :- we create a note object that hold id and body of our to do list

we implemnted an interactive cli to implement crud
*/package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Note struct {
	content string
}

var notes = make(map[int]Note)
var idCounter = 1

func displayNotes() {
	if len(notes) == 0 {
		fmt.Println("No notes available.")
		return
	}

	fmt.Println("Displaying all notes:")
	for id, note := range notes {
		fmt.Printf("ID: %d, Content: %s\n", id, note.content)
	}
}

func createNote(content string) int {
	note := Note{content: content}
	notes[idCounter] = note
	idCounter++
	return idCounter - 1
}

func deleteNote(id int) bool {
	_, exists := notes[id]
	if exists {
		delete(notes, id)
	} else {
		fmt.Println("Invalid input")
	}
	return exists
}

func readNote() {
	displayNotes()
	fmt.Println("Press Enter to go back to the main page")
	var ent string
	fmt.Scanln(&ent)
	manageState()
}

func creatingPage() {
	fmt.Print("You can enter your note below:\n")
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

	message = strings.TrimSpace(message)

	createNote(message)
	manageState()
}

func deletingPage() {
	fmt.Println("Please enter the index of the message you want to delete:\n")
	displayNotes()
	var index int
	fmt.Scan(&index)
	deleteNote(index)
	manageState()
}

func manageState() {
	fmt.Println("Please enter one of the following options:\n  C for create\n  D for delete\n  R for read\n  X to exit")

	var request string
	fmt.Scanln(&request)
	request = strings.ToLower(request)

	switch request {
	case "c":
		creatingPage()
	case "d":
		deletingPage()
	case "r":
		readNote()
	case "x":
		fmt.Println("Closed")
		fmt.Println("Thanks for using our app!")
		return
	default:
		fmt.Println("Invallid option, please try again.")
		manageState()
	}
}

func main() {
	manageState()
}
