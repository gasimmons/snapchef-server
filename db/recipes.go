package db

import "database/sql"

func InsertRecipeWithUser(title, ingredients string, userID int) error {
	res, err := DB.Exec("INSERT INTO recipes (title, ingredients) VALUES (?, ?)", title, ingredients)
	if err != nil {
		return err
	}
	recipeID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	_, err = DB.Exec("INSERT INTO user_recipes (userId, recipeId) VALUES (?, ?)", userID, recipeID)
	return err
}

func GetRecipesByUser(userID string) ([]Recipe, error) {
	rows, err := DB.Query(`
		SELECT recipes.recipeId, recipes.title, recipes.ingredients
		FROM recipes
		JOIN user_recipes ON recipes.recipeId = user_recipes.recipeId
		WHERE user_recipes.userId = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var recipes []Recipe
	for rows.Next() {
		var r Recipe
		if err := rows.Scan(&r.ID, &r.Title, &r.Ingredients); err != nil {
			return nil, err
		}
		recipes = append(recipes, r)
	}
	return recipes, nil
}
