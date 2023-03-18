package main

import (
	"encoding/json"
	"fmt"
	"time"

	"herculesgabriel/golang-playground/src/jobs"
	"herculesgabriel/golang-playground/src/schemas"
	mail "herculesgabriel/golang-playground/src/services/mail/implementations"
	"herculesgabriel/golang-playground/src/usecases"
	"herculesgabriel/golang-playground/src/utils"
)

func basicLearnings() {
	divisionResult, err := utils.Divide(0, 2)
	if err != nil {
		utils.Logger("Error while dividing values")
	}

	sumResult := utils.Sum(1, 2)
	sumAllResult := utils.SumAll(1, 2, 3)

	fmt.Println("0 / 2 =", divisionResult)
	fmt.Println("1 + 2 =", sumResult)
	fmt.Println("1 + 2 + 3 =", sumAllResult)
}

func threads() {
	go jobs.Countdown("second countdown", 5)
	jobs.Countdown("first countdown", 10)
}

func channels() {
	stringChannel := make(chan string)

	go func() {
		stringChannel <- "Hello,"
		stringChannel <- "World!"
		stringChannel <- "-"
	}()

	var endSignal string
	for endSignal != "-" {
		msg := <-stringChannel

		if msg == "-" {
			endSignal = msg
			continue
		}

		time.Sleep(time.Second)
		fmt.Println(msg)
	}
}

func arraysAndSlices() {
	var array [2]int = [2]int{1, 2}
	fmt.Println(array)

	var slice []int = []int{1, 2}
	fmt.Println(slice)

	var newSlice []int = make([]int, 2)
	fmt.Println(newSlice)
}

func maps() {
	var activeModules map[string]bool = make(map[string]bool)
	activeModules["finances"] = true
	activeModules["health"] = false

	emails := map[string]string{
		"bob":  "bob@mail.com",
		"kate": "kate@mail.com",
	}

	fmt.Println(activeModules)
	fmt.Println(emails["bob"])

	delete(emails, "kate")
	_, exists := activeModules["kate"]
	fmt.Println(exists)
}

func structs() {
	someone := schemas.Person{
		Name: "Someone",
		Age:  27,
		Address: schemas.Address{
			City:    "Curitiba",
			Country: "Brasil",
		},
	}

	city := schemas.Address{
		City:    "Colombo",
		Country: "Brasil",
	}

	someoneElse := schemas.Person{
		Name:    "Someone Else",
		Age:     26,
		Address: city,
	}

	supermarket := schemas.Place{
		Name:    "Menudo",
		Code:    1,
		Address: city,
	}

	fmt.Println(someone)
	fmt.Println(someoneElse)
	fmt.Println(supermarket)

	fmt.Println(supermarket.Address.City)
	fmt.Println(supermarket.City)

	fmt.Println(someone.Greeting())
}

func serializer() {
	team := schemas.NewTeam()

	team.AddPlayer("Bob")
	team.AddPlayer("Kate")

	result, _ := json.Marshal(team)

	data := []byte(result)
	var parsedTeam schemas.Team
	json.Unmarshal(data, &parsedTeam)

	fmt.Println("JSON -> " + string(result))
	fmt.Println()
	fmt.Print("Team -> ")
	fmt.Println(parsedTeam)
}

func interfaces() {
	goodMail := mail.GoodMail{}
	sendPromoEmail := usecases.NewSendPromoEmail(goodMail)

	sendPromoEmail.Execute("bob@mail.com")
}

func Runner(funcs ...func()) {
	for _, fn := range funcs {
		fn()
		fmt.Println()
	}
}

func main() {
	funcs := []func(){
		basicLearnings,
		threads,
		channels,
		arraysAndSlices,
		maps,
		structs,
		serializer,
		interfaces,
	}

	Runner(funcs...)


	
}
