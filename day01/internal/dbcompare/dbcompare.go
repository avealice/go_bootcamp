package dbcompare

import (
	"fmt"
	"reflect"
)

func CompareDataBases(originalDB, stolenDB DataBase) {
	CheckAdded(originalDB.Recipes, stolenDB.Recipes)
	CheckRemoved(originalDB.Recipes, stolenDB.Recipes)
	CheckTime(originalDB.Recipes, stolenDB.Recipes)
	for _, originalRecipe := range originalDB.Recipes {
		stolenRecipe, found := FindRecipeByName(stolenDB.Recipes, originalRecipe.Name)
		if !found {
			continue
		}
		CompareIngredients(originalRecipe.Name, originalRecipe.Ingredients, stolenRecipe.Ingredients)
	}
}

func FindRecipeByName(recipes []Recipe, name string) (Recipe, bool) {
	for _, recipe := range recipes {
		if recipe.Name == name {
			return recipe, true
		}
	}
	return Recipe{}, false
}

func CheckAdded(originalRecipes, stolenRecipes []Recipe) {
	for _, stolenRecipe := range stolenRecipes {
		_, found := FindRecipeByName(originalRecipes, stolenRecipe.Name)
		if !found {
			fmt.Printf("ADDED cake \"%s\"\n", stolenRecipe.Name)
		}
	}
}

func CheckRemoved(originalRecipes, stolenRecipes []Recipe) {
	for _, originalRecipe := range originalRecipes {
		_, found := FindRecipeByName(stolenRecipes, originalRecipe.Name)
		if !found {
			fmt.Printf("REMOVED cake \"%s\"\n", originalRecipe.Name)
		}
	}
}

func CheckTime(originalRecipes, stolenRecipes []Recipe) {
	for _, originalRecipe := range originalRecipes {
		stolenRecipe, found := FindRecipeByName(stolenRecipes, originalRecipe.Name)
		if originalRecipe.StoveTime != stolenRecipe.StoveTime && found {
			fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", originalRecipe.Name, stolenRecipe.StoveTime, originalRecipe.StoveTime)
		}
	}
}

func findIngredientByName(ingredients []Ingredient, name string) (Ingredient, bool) {
	for _, ingredient := range ingredients {
		if ingredient.ItemName == name {
			return ingredient, true
		}
	}
	return Ingredient{}, false
}

func CompareIngredients(recipeName string, originalIngredients, stolenIngredients []Ingredient) {
	for _, stolenIngredient := range stolenIngredients {
		_, found := findIngredientByName(originalIngredients, stolenIngredient.ItemName)
		if !found {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", stolenIngredient.ItemName, recipeName)
		}
	}
	for _, originalIngredient := range originalIngredients {
		_, found := findIngredientByName(stolenIngredients, originalIngredient.ItemName)
		if !found {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", originalIngredient.ItemName, recipeName)
			continue
		}
	}

	for _, originalIngredient := range originalIngredients {
		stolenIngredient, found := findIngredientByName(stolenIngredients, originalIngredient.ItemName)
		if !found {
			continue
		}
		if !reflect.DeepEqual(originalIngredient, stolenIngredient) {
			printIngredientChanges(originalIngredient, stolenIngredient, recipeName)
		}
	}
}

func printIngredientChanges(original, stolen Ingredient, recipeName string) {
	if original.ItemCount != stolen.ItemCount && original.ItemUnit == stolen.ItemUnit {
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
			original.ItemName, recipeName, stolen.ItemCount, original.ItemCount)
	}

	if original.ItemUnit != stolen.ItemUnit {
		if stolen.ItemUnit == "" {
			fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", original.ItemUnit, original.ItemName, recipeName)
		} else if original.ItemUnit == "" {
			fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", stolen.ItemUnit, original.ItemName, recipeName)
		} else {
			fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", original.ItemName, recipeName, stolen.ItemUnit, original.ItemUnit)
		}
	}
}
