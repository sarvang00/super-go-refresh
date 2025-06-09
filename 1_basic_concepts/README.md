# Project Idea: "Go Recipe Book & Calorie Counter"

### Concept
You'll build a simple command-line application that allows a user to define a few recipes, calculate their total calories (based on ingredient calories), and potentially flag recipes that exceed a certain calorie limit.

### Why it fits the criteria
- Basic Go Syntax: You'll be writing pure Go code.
- Sub-2-hr
- Variables:
    - Declare variables for recipe names, ingredient names, calorie counts, total calories, user input, etc. 
    - Example: `var recipeName string`, `var totalCalories float64`
- Data Types:
    - string for names.
    - float64 for calorie counts (allows for decimals).
    - int for loop counters or simple quantities.
    - bool for conditional flags (e.g., isHighCalorie).
- Loops:
    - for loop: Iterate through ingredients in a recipe to calculate total calories.
    - for loop (infinite with break): For a main menu loop where the user can choose to add recipes, view them, or exit.
- Conditionals:
    - if/else if/else:
        - Check user input for menu choices.
        - Determine if a recipe is "high calorie" (e.g., if totalCalories > 500).
        - Handle potential errors during user input (e.g., non-numeric calorie input).
- Error, Panic, Recover:
    - Error: Implement a simple error return from a function that parses calorie input (e.g., if strconv.ParseFloat fails). Handle this error gracefully.
    - Panic: Demonstrate a panic in a truly unrecoverable scenario (e.g., if a critical ingredient is missing and you must have it, though for this simple app, try to avoid genuine panics in core logic and use them more as a demonstration). A better way to demonstrate panic in a small app would be to simulate an extreme, unexpected error, such as a function that should always return a value but for a specific, edge case input, you decide to panic to indicate a serious logical flaw.
    - Recover: Pair recover with panic within a defer statement to catch the panic and print a message instead of crashing the program. You could wrap the calorie calculation in a defer that calls recover to illustrate this.
- Functions:
    Create functions like:
    - addRecipe(): Prompts user for recipe details.
    - calculateTotalCalories(ingredients []Ingredient): Takes a slice of ingredients and returns the total calorie count.
    - displayRecipe(recipe Recipe): Prints recipe details.
    - mainMenu(): Displays menu options and handles user input.
- Packages, Imports, Exports:
    - main package: Your primary application logic.
    - fmt: For input/output.
    - strconv: To convert string input to numbers.
    - errors: To create custom error types (optional, but good for demonstrating).
    - Custom package (optional but good for demonstration): Create a small sub-directory models/ with models.go containing your Ingredient and Recipe structs. This demonstrates how to import and use types from other packages. Export these structs and any related helper functions (by capitalizing their first letter).
- Type Casting, Interface:
    - Type Casting: When converting user input (string) to float64 using strconv.ParseFloat.
    - Interface (conceptual): While a full-blown interface implementation might be overkill for a sub-2-hour project, you can introduce the idea of an interface. For instance, imagine if you had different types of "food items" (e.g., Ingredient, PreparedMeal) and wanted a function PrintNutritionalInfo(item Nutritious) where Nutritious is an interface with a GetCalories() method. You could simply define the Nutritious interface and have your Ingredient struct implement it (by having a GetCalories() method). You wouldn't need multiple concrete types to demonstrate the concept.
- Arrays, Slices, Maps:
    - Slices: Most useful for storing a dynamic list of Ingredient structs within a Recipe struct (`[]Ingredient`).
    - Maps: Use a map to store recipes by name for easy lookup (e.g., `recipes map[string]Recipe`). This allows you to quickly retrieve a recipe by its name.
    - Arrays: Less likely to be the primary choice due to fixed size, but you could briefly mention their existence or use one for a very small, fixed set of options (e.g., a simple menu display).
- Make, Structs:
    - Structs: Define Ingredient and Recipe structs.
        ```Go
        type Ingredient struct {
            Name    string
            Calories float64
        }

        type Recipe struct {
            Name       string
            Ingredients []Ingredient
            TotalCalories float64
        }
        ```
    - make: Use make to initialize your slice of Ingredients within a Recipe or to initialize your recipes map:
        ```Go
        recipe.Ingredients = make([]Ingredient, 0)
        recipes = make(map[string]Recipe)
        ```