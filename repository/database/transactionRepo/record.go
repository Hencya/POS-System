package transactionRepo

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement:true;unique"`
	CreatedBy string    `gorm:"type:varchar(50);not null"`
	UpdatedBy string    `gorm:"type:varchar(50)"`
	Amount    int64     `gorm:"not null"`
	Notes     string    `gorm:"not null"`
	Date      time.Time `gorm:"not null"`
	Type      string    `gorm:"type:varchar(7);not null"`
}
