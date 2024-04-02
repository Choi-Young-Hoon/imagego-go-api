package database

import "gorm.io/gorm"

func NewImage() Image {
	return Image{
		dbConnection: GeteDefaultDBConnection(),
	}
}

type Image struct {
	gorm.Model
	UserID      string `gorm:"references:UserID" `
	Title       string `gorm:"not null"`
	Description string `gorm:""`
	ImageName   string `gorm:"not null"`

	dbConnection *DBConnection `gorm:"-"`
}

func (image *Image) TableName() string {
	return "image"
}

func (image *Image) Create() error {
	result := image.dbConnection.db.Create(image)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (image *Image) FindById(id string) error {
	result := image.dbConnection.db.Where("id = ?", id).First(image)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (image *Image) FindByUserId(userId string) ([]Image, error) {
	var images []Image
	result := image.dbConnection.db.Where("user_id = ?", userId).Find(&images)
	if result.Error != nil {
		return nil, result.Error
	}
	return images, nil
}
