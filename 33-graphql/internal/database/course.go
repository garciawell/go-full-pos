package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Couse struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Couse {
	return &Couse{db: db}
}

func (c *Couse) Create(name string, description string, categoryID string) (*Couse, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES (?, ?, ?, ?)", id, name, description, categoryID)
	if err != nil {
		return nil, err
	}
	return &Couse{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

func (c *Couse) FindAll() ([]*Couse, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*Couse
	for rows.Next() {
		var course Couse
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}
