package models

// Car represents the model for a car
type Car struct {
	ID    uint   `gorm:"primaryKey"`
	Brand string `gorm:"not null; type: varchar(191)"`
	Model string `gorm:"not null; type: varchar(191)"`
	Price int    `gorm:"not null; type: int"`
}
