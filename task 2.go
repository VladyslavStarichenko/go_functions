package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Stadium struct {
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Capacity int      `json:"capacity"`
	Sports   []string `json:"sports"`
}

func main() {
	// Create a new file and write structured data to it
	stadiums := []Stadium{
		{Name: "Wembley Stadium", Address: "Wembley, London", Capacity: 90000, Sports: []string{"Football"}},
		{Name: "Old Trafford", Address: "Sir Matt Busby Way, Old Trafford, Manchester", Capacity: 75000, Sports: []string{"Football"}},
	}
	if err := WriteStadiumsToFile("stadiums.json", stadiums); err != nil {
		log.Fatal(err)
	}

	// Print the contents of the JSON file to the console
	if err := PrintJsonFileContents("stadiums.json"); err != nil {
		log.Fatal(err)
	}

	// Remove a stadium from the JSON file
	if err := RemoveStadiumFromFile("stadiums.json", "Old Trafford"); err != nil {
		log.Fatal(err)
	}

	// Add new stadiums to the JSON file
	newStadiums := []Stadium{
		{Name: "Emirates Stadium", Address: "Hornsey Rd, London", Capacity: 60000, Sports: []string{"Football"}},
		{Name: "Etihad Stadium", Address: "Ashton New Rd, Manchester", Capacity: 55000, Sports: []string{"Football"}},
	}
	if err := AddStadiumsToFile("stadiums.json", newStadiums); err != nil {
		log.Fatal(err)
	}

	// Display the updated contents of the JSON file on the console
	if err := displayFileContent("stadiums.json"); err != nil {
		log.Fatal(err)
	}
}

// 1. Create a new JSON file and write structured data to it

func WriteStadiumsToFile(filename string, stadiums []Stadium) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(stadiums, "", "  ")
	if err != nil {
		return err
	}

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}

// 2. Print the contents of a JSON file to the console

func PrintJsonFileContents(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}

// 3. Remove a stadium with a given name from a JSON file

func RemoveStadiumFromFile(filename string, name string) error {
	// Read the existing JSON file into memory
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// Parse the JSON data into a slice of Stadium structs
	stadiums := []Stadium{}
	if err := json.Unmarshal(data, &stadiums); err != nil {
		return err
	}

	// Remove the stadium with the given name
	for i, stadium := range stadiums {
		if stadium.Name == name {
			stadiums = append(stadiums[:i], stadiums[i+1:]...)
			break
		}
	}

	// Marshal the updated slice back into JSON format
	updatedData, err := json.MarshalIndent(stadiums, "", "  ")
	if err != nil {
		return err
	}

	// Write the updated JSON data back to the file
	if err := ioutil.WriteFile(filename, updatedData, 0644); err != nil {
		return err
	}

	return nil
}

// 4. Add K new stadiums to the end of a JSON file

func AddStadiumsToFile(filename string, newStadiums []Stadium) error {
	// Read the existing JSON file into memory
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// Parse the JSON data into a slice of Stadium structs
	stadiums := []Stadium{}
	if err := json.Unmarshal(data, &stadiums); err != nil {
		return err
	}

	// Add the new stadiums to the slice
	stadiums = append(stadiums, newStadiums...)

	// Marshal the updated slice back into JSON format
	updatedData, err := json.MarshalIndent(stadiums, "", "  ")
	if err != nil {
		return err
	}

	// Write the updated JSON data back to the file
	if err := ioutil.WriteFile(filename, updatedData, 0644); err != nil {
		return err
	}

	return nil
}

// Функція для виведення зміненого файлу на екран
func displayFileContent(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
