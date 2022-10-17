package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        string
	Brand     string
	Name      string
	Price     int
	Quantity  int
	Total     int
	Title     string
	CreatedAt time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.NewString()

	if len(p.Name) < 5 {
		return fmt.Errorf("nama terlalu pendek")
	}

	p.Total = p.Quantity * p.Price

	return nil
}

func (p *Product) AfterCreate(tx *gorm.DB) error {
	fmt.Println("created success with id", p.ID)
	return nil
}
