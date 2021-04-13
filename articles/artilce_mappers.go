package articles

import (
	"time"

	articlesDomain "github.com/DonaldArmandoArteaga/go-rest-api/articles/articles_domain"
)

func ArticleDTOToArticle(a *ArticleCommandDTO) (*articlesDomain.Article, error) {
	article, err := articlesDomain.CreateArticle(
		&articlesDomain.CreateArticleImput{
			Slug:           "",
			Title:          a.Title,
			Description:    a.Description,
			Body:           a.Body,
			TagList:        a.TagList,
			CreatedAt:      time.Time{},
			UpdatedAt:      time.Time{},
			Favorited:      false,
			FavoritesCount: 0,
			Username:       "",
			Bio:            "",
			Image:          "",
			Following:      false,
		},
	)

	if err != nil {
		return &articlesDomain.Article{}, err
	}

	return article, nil
}

func ArticleToArticleQueryDTO(ar *articlesDomain.Article) *ArticleQueryDTO {
	return &ArticleQueryDTO{
		Slug:           ar.Slug(),
		Title:          ar.Title(),
		Description:    ar.Description(),
		Body:           ar.Body(),
		TagList:        ar.TagList(),
		CreatedAt:      ar.CreatedAt(),
		UpdatedAt:      ar.UpdatedAt(),
		Favorited:      ar.Favorited(),
		FavoritesCount: ar.FavoritesCount(),
		Author: AuthorQueryDTO{
			Username:  ar.Username(),
			Bio:       ar.Bio(),
			Image:     ar.Image(),
			Following: ar.Following(),
		},
	}
}

func ArticleModelRepositoryToArticle(ar *ArticleModelRepository) (*articlesDomain.Article, error) {

	tagList := []string{}
	for _, tag := range ar.TagList {
		tagList = append(tagList, tag.Value)
	}

	article, err := articlesDomain.CreateArticle(
		&articlesDomain.CreateArticleImput{
			Slug:           ar.Slug,
			Title:          ar.Title,
			Description:    ar.Description,
			Body:           ar.Body,
			TagList:        tagList,
			CreatedAt:      ar.CreatedAt,
			UpdatedAt:      ar.UpdatedAt,
			Favorited:      ar.Favorited,
			FavoritesCount: ar.FavoritesCount,
			Username:       ar.Username,
			Bio:            ar.Bio,
			Image:          ar.Image,
			Following:      ar.Following,
		},
	)

	if err != nil {
		return &articlesDomain.Article{}, err
	}

	return article, nil
}

func ArticleToArticleModelRepository(ar *articlesDomain.Article) *ArticleModelRepository {

	return &ArticleModelRepository{
		ID:             0,
		Slug:           ar.Slug(),
		Title:          ar.Title(),
		Description:    ar.Description(),
		Body:           ar.Body(),
		CreatedAt:      ar.CreatedAt(),
		UpdatedAt:      ar.UpdatedAt(),
		Favorited:      ar.Favorited(),
		FavoritesCount: ar.FavoritesCount(),
		Username:       ar.Username(),
		Bio:            ar.Bio(),
		Image:          ar.Image(),
		Following:      ar.Following(),
	}
}
