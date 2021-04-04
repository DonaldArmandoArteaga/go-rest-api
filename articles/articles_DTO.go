package articles

import "time"

type (
	ArticleCommandDTO struct {
		Title       string   `json:"title" binding:"required"`
		Description string   `json:"description" binding:"required"`
		Body        string   `json:"body" binding:"required"`
		TagList     []string `json:"tagList"`
	}

	ArticleQueryDTO struct {
		Slug           string         `json:"slug"`
		Title          string         `json:"title"`
		Description    string         `json:"description"`
		Body           string         `json:"body"`
		TagList        []string       `json:"tagList"`
		CreatedAt      time.Time      `json:"createdAt"`
		UpdatedAt      time.Time      `json:"updatedAt"`
		Favorited      bool           `json:"favorited"`
		FavoritesCount int            `json:"favoritesCount"`
		Author         AuthorQueryDTO `json:"author"`
	}

	AuthorQueryDTO struct {
		Username  string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	}
)
