package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type food interface {
	isHighCalorie() bool
}

type Ingridient struct {
	Name     string
	Calories float64
}

// Checks if an ingridient is High Calorie.
// For each ingridient, the bar is at 70, to check if it is High Calorie.
func (i Ingridient) isHighCalorie() bool {
	return i.Calories < 40
}

type Recipie struct {
	Name                string
	IngridientsQuantity map[Ingridient]int
	totalCalories       float64
}

func (r Recipie) isHighCalorie() bool {
	return r.totalCalories < 500
}

func (r Recipie) displayRecipie() {
	fmt.Println("Ingridient\tQuantity")
	fmt.Println("------------\t------")
	for ing, quant := range r.IngridientsQuantity {
		fmt.Printf("%v\t%v\n", ing.Name, quant)
	}
	fmt.Println("--------------------")
	fmt.Println("Total Calories: ", r.totalCalories)
}

// Calculates total calories for a Recipie
func culatetotalCalories(IngridientsQuantity map[Ingridient]int) float64 {
	var totalCals float64 = 0

	for ing, ingQuant := range IngridientsQuantity {
		totalCals += ing.Calories * float64(ingQuant)
	}

	return totalCals
}

func judge(f food) bool {
	// fmt.Println(f)
	// if _, ok := f.(Ingridient); ok {
	// 	fmt.Println("Judging food ingridient")
	// } else if _, ok := f.(Recipie); ok {
	// 	fmt.Println("Judging food recipie")
	// }
	// fmt.Println("Is it high calorie?", f.isHighCalorie())
	return f.isHighCalorie()
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		log.Fatal("Error while reading input!")
	}
	return strings.TrimSpace(input), err
}

func getUserInput(r *bufio.Reader) int {
	fmt.Println("Select options:")
	fmt.Println("1. Add ingridient")
	fmt.Println("2. Print ingridients")
	fmt.Println("3. Add recipies")
	fmt.Println("4. Print recipies")
	fmt.Println("5. Get Health Report")
	fmt.Println("6. Exit")

	userInput, err1 := getInput("Enter your choice: ", r)

	if err1 != nil {
		fmt.Println("Error: ", err1)
	}

	userIntInput, err2 := strconv.Atoi(userInput)

	if err2 != nil {
		fmt.Println("Error: ", err2)
	}

	return userIntInput
}

func addIngridient(ingridientName string, calorieCount float64) Ingridient {
	return Ingridient{
		Name:     ingridientName,
		Calories: calorieCount,
	}
}

func addRecipie(recipieName string, ingridientsMap map[Ingridient]int) Recipie {
	return Recipie{
		Name:                recipieName,
		IngridientsQuantity: ingridientsMap,
		totalCalories:       culatetotalCalories(ingridientsMap),
	}
}

func printIngridients(ingridients []Ingridient) {
	fmt.Println("Ingridient Name \t Calories")
	fmt.Println("---------------------------")
	for _, ing := range ingridients {
		fmt.Printf("%s\t%v\n", ing.Name, ing.Calories)
	}
	fmt.Println("---------------------------")
}

func printRecipies(recipies []Recipie) {
	fmt.Println("Recipie Name \t Total Calories")
	fmt.Println("------------------------------")
	for _, recipie := range recipies {
		fmt.Printf("%s \t %f\n", recipie.Name, recipie.totalCalories)
		fmt.Println("--")
		fmt.Println("Ingridient Name \t Quantity")
		for item, quant := range recipie.IngridientsQuantity {
			fmt.Printf("%s \t %d\n", item.Name, quant)
		}
		fmt.Println("----")
	}
	fmt.Println("------------------------------")
}

func getHealthReport(ingridients []Ingridient, recipies []Recipie) {
	fmt.Println("Recipie Name \t Healthy?")
	fmt.Println("------------------------------")
	for _, recipie := range recipies {
		fmt.Println(recipie.Name, "\t", judge(recipie))
	}
	fmt.Println("------------------------------")
	fmt.Println("Ingridient Name \t Healthy?")
	fmt.Println("------------------------------")
	for _, ingridient := range ingridients {
		fmt.Println(ingridient.Name, "\t", judge(ingridient))
	}
	fmt.Println("------------------------------")
}

func actionLogic(ingridients []Ingridient, recipies []Recipie, userInput int, reader *bufio.Reader) {
	switch userInput {
	case 1:
		ingridientName, err1 := getInput("Enter ingridient name: ", reader)
		if err1 != nil {
			fmt.Println(err1)
		}
		calorieCountString, err2 := getInput("Enter calorie count value: ", reader)
		if err2 != nil {
			fmt.Println(err2)
		}
		calorieCount, err3 := strconv.ParseFloat(calorieCountString, 64)
		if err3 != nil {
			fmt.Println(err3)
		}
		ingridients = append(ingridients, addIngridient(ingridientName, calorieCount))
		userInput = getUserInput(reader)
		actionLogic(ingridients, recipies, userInput, reader)
	case 2:
		printIngridients(ingridients)
		userInput = getUserInput(reader)
		actionLogic(ingridients, recipies, userInput, reader)
	case 3:
		recipieName, err1 := getInput("Enter ingridient name: ", reader)
		if err1 != nil {
			fmt.Println(err1)
		}
		ingridientsMap := make(map[Ingridient]int)
		for _, ingridient := range ingridients {
			ingridientsMap[ingridient] = 1
		}
		recipies = append(recipies, addRecipie(recipieName, ingridientsMap))
		userInput = getUserInput(reader)
		actionLogic(ingridients, recipies, userInput, reader)
	case 4:
		printRecipies(recipies)
		userInput = getUserInput(reader)
		actionLogic(ingridients, recipies, userInput, reader)
	case 5:
		getHealthReport(ingridients, recipies)
		userInput = getUserInput(reader)
		actionLogic(ingridients, recipies, userInput, reader)
	case 6:
		os.Exit(0)
	default:
		userInput = getUserInput(reader)
		actionLogic(ingridients, recipies, userInput, reader)
	}
}

func main() {
	// ingridient1 := Ingridient{
	// 	Name:     "Pepper",
	// 	Calories: 8,
	// }

	// ingridient2 := Ingridient{
	// 	Name:     "Pepper Mozza Sauce",
	// 	Calories: 85,
	// }

	var ingridients []Ingridient
	var recipies []Recipie

	// fmt.Println(ingridient1)
	// fmt.Println(ingridient1.isHighCalorie())

	// fmt.Println(ingridient2)
	// fmt.Println(ingridient2.isHighCalorie())

	// recpie1 := Recipie{
	// IngridientsQuantity: make(map[Ingridient]int),
	// 	totalCalories:       0,
	// }

	// recpie1.IngridientsQuantity[ingridient1] = 2
	// recpie1.IngridientsQuantity[ingridient2] = 1
	// recpie1.TalculatetotalCalories()

	// recpie1.displayRecipie()
	// judge(recpie1)
	// judge(ingridient1)
	// judge(ingridient2)

	reader := bufio.NewReader(os.Stdin)

	userInput := getUserInput(reader)
	actionLogic(ingridients, recipies, userInput, reader)

}
