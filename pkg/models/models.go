package models

import (
	"time"
	"fmt"
	"gorm.io/gorm"
)

// Coauthor Model
type Coauthor struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	PaperID  int    `json:"paper_id" gorm:"index"`
	Coauthor string `json:"coauthor"`
	Name     string `json:"name"`
}

// Conference Model
type Conference struct {
	ID       int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Date     time.Time `json:"date"`
	Platform string    `json:"platform"`
	Title    string    `json:"title"`
}

// Feedback Model
type Feedback struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Feedback string `json:"feedback"`
}

// Login Model
type Login struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

// People Model
type People struct {
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"`
	Department string `json:"department"`
	Type       string `json:"type"`
	Domain     string `json:"domain,omitempty"`
}

// Project Model
type Project struct {
	ProjectID       int       `json:"project_id" gorm:"primaryKey;autoIncrement"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	InstructorEmail string    `json:"instructor_email,omitempty"`
	ProjectLead     string    `json:"project_lead"`
	Domain          string    `json:"domain"`
	Status          string    `json:"status"`
	Citations       int       `json:"citations"`
	Year 			int 	  `json:"year,omitempty"`

}

// Publication Model
type Publication struct {
	PaperID    int       `json:"paper_id" gorm:"primaryKey;autoIncrement"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Publisher  string    `json:"publisher"`
	Year 	   int 	  	 `json:"year,omitempty"`
	Conference string    `json:"conference"`
	Citations  int       `json:"citations"`
	Link       string    `json:"link"`
}

// Research Model
type Research struct {
	ResearchID int       `json:"research_id" gorm:"primaryKey;autoIncrement"`
	Title      string    `json:"title"`
	Domain     string    `json:"domain"`
	Lead       string    `json:"lead"`
	Year 	   int 	  	 `json:"year,omitempty"`
	Citations  int       `json:"citations"`
	Status     string    `json:"status"`
}

// ResearchMember Model
type ResearchMember struct {
	ResearchID int    `json:"research_id"`
	MemberName string `json:"member_name"`
	Email      string `json:"email"`
}

// Student Model
type Student struct {
	Email       string `json:"email" gorm:"primaryKey"`
	Name        string `json:"name"`
	RollNumber  string `json:"roll_number" gorm:"unique"`
	Department  string `json:"department"`
	Batch       int    `json:"batch"`
	Supervisor  string `json:"supervisor"`
	ResearchID  int    `json:"research_id,omitempty"`
	ProjectID   int    `json:"project_id,omitempty"`
	Conference  int    `json:"conference,omitempty"`
	Publication int    `json:"publication,omitempty"`
}

// Supervisor Model
type Supervisor struct {
	Email      string `json:"email" gorm:"primaryKey"`
	Name       string `json:"name"`
	Department string `json:"department"`
	ResearchID int    `json:"research_id,omitempty"`
	ProjectID  int    `json:"project_id,omitempty"`
}

// AutoMigrate function for database migrations
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&Coauthor{},
		&Conference{},
		&Feedback{},
		&Login{},
		&People{},
		&Project{},
		&Publication{},
		&Research{},
		&ResearchMember{},
		&Student{},
		&Supervisor{},
	)
	if err != nil {
		fmt.Println("Database migration failed!")
	} else {
		fmt.Println("Database migration successful!")
	}
}