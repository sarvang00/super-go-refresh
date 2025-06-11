package main

import "fmt"

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
	return i.Calories > 70
}

type Recipie struct {
	Name                string
	IngridientsQuantity map[Ingridient]int
	totalCalories       float64
}

func (r Recipie) isHighCalorie() bool {
	return r.totalCalories > 680
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
func (r *Recipie) TalculatetotalCalories() {
	var totalCals float64 = 0

	for ing, ingQuant := range r.IngridientsQuantity {
		totalCals += ing.Calories * float64(ingQuant)
	}

	r.totalCalories = totalCals
}

func judge(f food) {
	// fmt.Println(f)
	if _, ok := f.(Ingridient); ok {
		fmt.Println("Judging food ingridient")
	} else if _, ok := f.(Recipie); ok {
		fmt.Println("Judging food recipie")
	}
	fmt.Println("Is it high calorie?", f.isHighCalorie())
}

func main() {
	ingridient1 := Ingridient{
		Name:     "Pepper",
		Calories: 8,
	}

	ingridient2 := Ingridient{
		Name:     "Pepper Mozza Sauce",
		Calories: 85,
	}

	// fmt.Println(ingridient1)
	// fmt.Println(ingridient1.isHighCalorie())

	// fmt.Println(ingridient2)
	// fmt.Println(ingridient2.isHighCalorie())

	recpie1 := Recipie{
		IngridientsQuantity: make(map[Ingridient]int),
		totalCalories:       0,
	}

	recpie1.IngridientsQuantity[ingridient1] = 2
	recpie1.IngridientsQuantity[ingridient2] = 1
	recpie1.TalculatetotalCalories()

	recpie1.displayRecipie()
	judge(recpie1)
	judge(ingridient1)
	judge(ingridient2)
}
