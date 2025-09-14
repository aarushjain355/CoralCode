package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/a9jain/coral-backend/db"
)

type File struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	FileName  string    `json:"file_name" gorm:"type:varchar(255);not null"`
	Content   string    `json:"content" gorm:"type:text;not null"` // Store code as text
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// Input all the database access functions below somehow

func CreateFile(file *File) error {
	query := `
        INSERT INTO files (id, filename, content, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := db.DB.Exec(query,
		file.ID,
		file.FileName,
		file.Content,
		file.CreatedAt,
		file.UpdatedAt,
	)

	if err != nil {
		fmt.Println("Error inserting file:", err)
		return fmt.Errorf("failed to insert file: %v", err)
	} else {
		fmt.Println("File inserted successfully")
	}

	return nil
}

func UpdateFile(fileName string, newContent string) (*File, error) {
	query := `
        UPDATE files
        SET content = $1, updated_at = $2
        WHERE filename = $3
        RETURNING id, filename, content, created_at, updated_at
    `

	var file File
	err := db.DB.QueryRow(query, newContent, time.Now(), fileName).
		Scan(&file.ID, &file.FileName, &file.Content, &file.CreatedAt, &file.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No file found with that filename
		}
		return nil, fmt.Errorf("failed to update file: %v", err)
	}

	return &file, nil
}
