package articles

import (
	"time"

	"github.com/DonaldArmandoArteaga/go-rest-api/tags"
)

func (*ArticleModelRepository) TableName() string {
	return "articles"
}

type ArticleModelRepository struct {
	ID             uint64 `gorm:"primaryKey"`
	Slug           string
	Title          string                      `gorm:"not null"`
	Description    string                      `gorm:"not null"`
	Body           string                      `gorm:"not null"`
	TagList        []*tags.TagsModelRepository `gorm:"many2many:articles_tags;"`
	CreatedAt      time.Time                   `gorm:"not null"`
	UpdatedAt      time.Time
	Favorited      bool
	FavoritesCount int
	Username       string
	Bio            string
	Image          string
	Following      bool
}
