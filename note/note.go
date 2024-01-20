package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct{
	Title string  `json:"title"`// struct tag -> it using when save to json file
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note *Note) Display(){
	fmt.Printf("\nYour note title: %v has following content:\n\n%v it created at: %v\n\n", note.Title, note.Content, note.CreatedAt)
}

func (note *Note) Save() (error){
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if(err != nil){
		return err
	}
	
	return os.WriteFile(fileName, json, 0644)
} 

func New(title string, content string)(*Note, error){
	if title == "" || content == ""{
		return &Note{}, errors.New("Invalid input")
	}

	return &Note{Title: title, Content: content, CreatedAt: time.Now()}, nil
}