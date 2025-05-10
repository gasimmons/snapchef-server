package db

type Recipe struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Ingredients string `json:"ingredients"`
}

func InsertRecipe(title, ingredients string) error {
	_, err := DB.Exec("INSERT INTO recipes (title, ingredients) VALUES (?, ?)", title, ingredients)
	return err
}

func GetAllRecipes() ([]Recipe, error) {
	rows, err := DB.Query("SELECT id, title, ingredients FROM recipes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
