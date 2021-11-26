package book

import (
	"encoding/json"
	"errors"
)

//get json string and parse it to the book
func (book *Book) UnarshalString(input string) error {
	data := []byte(input)
	json.Unmarshal(data, book)
	if book.Name == "" || book.Author == "" || book.Inventory == 0 || book.Pagecount == 0 {
		return errors.New("unvalid input os schema")
	}
	return nil
}

//get json string and parse it to the books
func (books *Books) UnarshalString(input string) error {
	data := []byte(input)
	json.Unmarshal(data, books)
	for _, book := range books.Books {
		if book.Name == "" || book.Author == "" || book.Inventory == 0 || book.Pagecount == 0 {
			return errors.New("unvalid input os schema")
		}
	}
	return nil
}

//convert book th the json string
func (book *Book) ToString() (string, error) {
	result, err := json.Marshal(book)
	if err != nil {
		return "", err
	}
	sResult := string(result)
	return sResult, nil

}
