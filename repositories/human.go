package repositories

import (
	"log"

	"github.com/jrmbrgs/goodkarma-api/library/models"
	"gorm.io/gorm"
)

type HumanRepository struct {
	Database *gorm.DB
}

func (r HumanRepository) Save(nh models.NewHuman) (*models.Human, error) {
	log.Println(">HumanRepository.save", nh.Name)

	human := models.Human{}
	human.Name = nh.Name
	r.Database.Create(&human)
	return &human, nil
}
