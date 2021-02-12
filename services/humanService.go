package services

import (
	"log"

	"github.com/jrmbrgs/goodkarma-api/library/models"
	"github.com/jrmbrgs/goodkarma-api/repositories"
)

type HumanService struct {
	Repo *repositories.HumanRepository
}

func (hs *HumanService) BecomeAGoodKarmaMember(h models.NewHuman) *models.Human {
	log.Println(">HumanService.BecomeAGoodKarmaMember")
	var human, _ = hs.Repo.Save(h)
	return human
}

func (hs *HumanService) CommitAnOffense() *models.Offense {
	var o models.Offense = models.Offense{}
	return &o
}
