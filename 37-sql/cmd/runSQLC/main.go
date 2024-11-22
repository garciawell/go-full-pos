package main

import (
	"context"
	"database/sql"

	"github.com/garciawell/go-full-pos/37-sql/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")

	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// 	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 		ID:          uuid.New().String(),
	// 		Name:        "Backend",
	// 		Description: sql.NullString{String: "Backend development", Valid: true},
	// 	})

	// 	if err != nil {
	// 		panic(err)
	// 	}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "636f62c5-cd1f-40f1-a5d4-681b7eae0a78",
		Name:        "Backend UPDATED",
		Description: sql.NullString{String: "Backend development UPDATED", Valid: true},
	})

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	err = queries.DeleteCategory(ctx, "636f62c5-cd1f-40f1-a5d4-681b7eae0a78")
}
