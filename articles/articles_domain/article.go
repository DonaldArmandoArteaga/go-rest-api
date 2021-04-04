package articles_domain

import (
	"fmt"
	"strings"
	"time"
)

type (
	CreateArticleImput struct {
		Slug           string
		Title          string
		Description    string
		Body           string
		TagList        []string
		CreatedAt      time.Time
		UpdatedAt      time.Time
		Favorited      bool
		FavoritesCount int
		Username       string
		Bio            string
		Image          string
		Following      bool
	}

	Article struct {
		slug           string
		title          string
		description    string
		body           string
		tagList        []string
		createdAt      time.Time
		updatedAt      time.Time
		favorited      bool
		favoritesCount int
		author         author
	}

	author struct {
		username  string
		bio       string
		image     string
		following bool
	}
)

func CreateArticle(cai *CreateArticleImput) (*Article, error) {

	createArticlesErrors := []string{}

	if cai.Title == "" {
		createArticlesErrors = append(
			createArticlesErrors,
			"El titulo es requerido",
		)
	}

	if cai.Description == "" {
		createArticlesErrors = append(
			createArticlesErrors,
			"La descripcion es requerida",
		)
	}

	if cai.Body == "" {
		createArticlesErrors = append(
			createArticlesErrors,
			"El cuerpo es requerido",
		)
	}

	if len(createArticlesErrors) > 0 {
		return &Article{}, fmt.Errorf(strings.Join(createArticlesErrors, "\n"))
	}

	return &Article{
			slug:           cai.Slug,
			title:          cai.Title,
			description:    cai.Description,
			body:           cai.Body,
			tagList:        cai.TagList,
			createdAt:      cai.CreatedAt,
			updatedAt:      cai.UpdatedAt,
			favorited:      cai.Favorited,
			favoritesCount: cai.FavoritesCount,
			author: author{
				username:  cai.Username,
				bio:       cai.Bio,
				image:     cai.Image,
				following: cai.Following,
			},
		},
		nil
}

func (ar *Article) Slug() string {
	return ar.slug
}

func (ar *Article) Title() string {
	return ar.title
}

func (ar *Article) Description() string {
	return ar.description
}

func (ar *Article) Body() string {
	return ar.body
}

func (ar *Article) TagList() []string {
	return ar.tagList
}

func (ar *Article) CreatedAt() time.Time {
	return ar.createdAt
}

func (ar *Article) UpdatedAt() time.Time {
	return ar.updatedAt
}

func (ar *Article) Favorited() bool {
	return ar.favorited
}

func (ar *Article) FavoritesCount() int {
	return ar.favoritesCount
}

func (ar *Article) Username() string {
	return ar.author.username
}

func (ar *Article) Bio() string {
	return ar.author.bio
}

func (ar *Article) Image() string {
	return ar.author.image
}

func (ar *Article) Following() bool {
	return ar.author.following
}
