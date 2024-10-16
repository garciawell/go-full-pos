package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name string, description string, categoryID string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO Courses (id, name, description, category_id) VALUES (?, ?, ?, ?)", id, name, description, categoryID)
	if err != nil {
		return nil, err
	}
	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

func (c *Course) FindAll() ([]*Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM Courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]*Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM Courses WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}
