package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name`
	Email string `json:"email`
	Phone string `json:"phone`
}

func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var contacts []Contact
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error loading contacts:", err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("=== CONTACTS MANAGER ===\n",
			"1. Add contact\n",
			"2. Show contacts\n",
			"3. Exit\n",
			"Please, choose an option: ")
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error reading option:", err)
			return
		}
		switch option {
		case 1:
			var newContact Contact
			fmt.Print("Name: ")
			newContact.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			newContact.Email, _ = reader.ReadString('\n')
			fmt.Print("Phone: ")
			newContact.Phone, _ = reader.ReadString('\n')

			contacts = append(contacts, newContact)

			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error saving contact", err)
			}
		case 2:
			fmt.Println("===============================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Name: %s Email: %s Phone: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("===============================================")
		case 3:
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
