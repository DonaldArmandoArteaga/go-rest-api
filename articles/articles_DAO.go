package articles

import (
	"time"

	"gorm.io/gorm"
)

func (ArticleModelRepository) TableName() string {
	return "articles"
}

func (TagListRepository) TableName() string {
	return "tags"
}

type ArticleModelRepository struct {
	gorm.Model

	ID             uint64 `gorm:"primaryKey"`
	Slug           string
	Title          string              `gorm:"not null"`
	Description    string              `gorm:"not null"`
	Body           string              `gorm:"not null"`
	TagList        []TagListRepository `gorm:"foreignKey:ArticleModelRepository"`
	CreatedAt      time.Time           `gorm:"not null"`
	UpdatedAt      time.Time
	Favorited      bool
	FavoritesCount int
	Username       string
	Bio            string
	Image          string
	Following      bool
}

type TagListRepository struct {
	gorm.Model

	ID                     uint64 `gorm:"primaryKey"`
	Value                  string
	ArticleModelRepository uint64
}
