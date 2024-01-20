package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note-app/note"
)


func main(){
	 title , content := getNoteData();
	 
	 userNote, err := note.New(title, content)
	 if(err != nil){
		fmt.Println(err)
		return
	 }

	 userNote.Display()

	 err = userNote.Save()
	
	 if(err != nil){
		print(err)
		return
	 }
	 fmt.Printf("Successfully save json file. \n")
}
func getNoteData()(title string, content string){
	title = getUserInput("Note Title:")
	content = getUserInput("Note Content:")

	return title, content
}

func getUserInput(input string) (value string){
	fmt.Printf("%v ",input)
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	// remove \n from string and \r if windows
	value = strings.TrimSuffix(value,"\n")
	value = strings.TrimSuffix(value,"\r")

	return value
}