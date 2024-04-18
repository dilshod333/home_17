package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Employee struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

func main() {
	
	file, err := os.Open("employees.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	
	var employees []Employee
	err = json.NewDecoder(file).Decode(&employees) 
	if err != nil {
		log.Fatal(err)
	}

	for i, emp := range employees {
		if emp.Id == 3 {
			employees[i].Age = 50
			break
		}
	}

	
	newEmployee := Employee{
		Id:       6,
		Name:     "Asilbek",
		Age:      29,
		Position: "programmer",
	}


	employees = append(employees, newEmployee)

	
	newData, err := json.MarshalIndent(employees, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	
	outputFile, err := os.Create("employees_new.json")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	_, err = outputFile.Write(newData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully saved to employees_new.json.")
}
