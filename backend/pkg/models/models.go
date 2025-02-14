package models

import (
	"time"
	"fmt"
	"gorm.io/gorm"
)

type Coauthor struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	PaperID  int    `json:"paper_id" gorm:"index"`
	Coauthor string `json:"coauthor"`
	Name     string `json:"name"`
}

type Conference struct {
	ID       int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Date     time.Time `json:"date"`
	Platform string    `json:"platform"`
	Title    string    `json:"title"`
}

type Feedback struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Feedback string `json:"feedback"`
}

type People struct {
	Name       string  `json:"name" gorm:"not null"`
	Email      string  `json:"email" gorm:"primaryKey;unique"`
	Password   string  `json:"password" gorm:"not null"` // Store **hashed** passwords
	Department string  `json:"department" gorm:"not null"`
	User_Type  string  `json:"user_type" gorm:"not null"`
	Domain     *string `json:"domain,omitempty"`
}

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

type Research struct {
	ResearchID int       `json:"research_id" gorm:"primaryKey;autoIncrement"`
	Title      string    `json:"title"`
	Domain     string    `json:"domain"`
	Lead       string    `json:"lead"`
	Year 	   int 	  	 `json:"year,omitempty"`
	Citations  int       `json:"citations"`
	Status     string    `json:"status"`
}

type BlacklistedToken struct {
	ID    uint   `gorm:"primaryKey"`
	Token string `gorm:"unique"`
}

type ResearchMember struct {
	ID		   int	  `json:"id" gorm:"primaryKey;autoIncrement"`
	ResearchID int    `json:"research_id" gorm:"index;foreignKey:ResearchID;references:ResearchID"`
	MemberName string `json:"member_name"`
	Email      string `json:"email"`
}

type Student struct {
	Email       string `json:"email" gorm:"primaryKey"`
	Name        string `json:"name"`
	RollNumber  string `json:"roll_number" gorm:"uniqueIndex"`
	Department  string `json:"department"`
	Batch       int    `json:"batch"`
	Supervisor  string `json:"supervisor"`
	ResearchID  *int   `json:"research_id,omitempty" gorm:"foreignKey:ResearchID;references:ResearchID"`
	ProjectID   *int   `json:"project_id,omitempty" gorm:"foreignKey:ProjectID;references:ProjectID"`
	Conference  *int   `json:"conference,omitempty" gorm:"foreignKey:Conference;references:ID"`
	Publication *int   `json:"publication,omitempty" gorm:"foreignKey:Publication;references:PaperID"`
}

type Supervisor struct {
	Email      string `json:"email" gorm:"primaryKey"`
	Name       string `json:"name"`
	Department string `json:"department"`
	ResearchID *int   `json:"research_id,omitempty" gorm:"foreignKey:ResearchID;references:ResearchID"`
	ProjectID  *int   `json:"project_id,omitempty" gorm:"foreignKey:ProjectID;references:ProjectID"`
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&Coauthor{},
		&Conference{},
		&Feedback{},
		&People{},
		&Project{},
		&Publication{},
		&Research{},
		&BlacklistedToken{},
		&ResearchMember{},
		&Student{},
		&Supervisor{},
	)
	fmt.Println("✅ Database migration successful!")
}