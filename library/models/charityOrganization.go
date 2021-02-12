package models

import (
	"time"
)

type CharityOrganization struct {
	ID                      string                   `json:"id"`
	Name                    string                   `json:"name"`
	CharityOrganizationTags []CharityOrganizationTag `json:"charityOrganizationTags" gorm:"many2many:charity_organization_tags;"`
	CreatedAt               time.Time                `json:"createdAt"`
	UpdatedAt               time.Time                `json:"updatedAt"`
}
