package main

import (
	"github.com/DonaldArmandoArteaga/go-rest-api/articles"
	articledomain "github.com/DonaldArmandoArteaga/go-rest-api/articles/articles_domain"
	"github.com/DonaldArmandoArteaga/go-rest-api/tags"
	"github.com/DonaldArmandoArteaga/go-rest-api/tags/tags_domain"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=America/Bogota"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//TODO handle error when initializate database connection
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	articleRepositoy, articlesError := articles.CreateArticleRepository(db)

	//TODO handle error when initializate app
	if articlesError != nil {
		panic(articlesError.Error())
	}

	articles.CreateArticlesController(
		&articles.ArticlesController{
			Gin: r,
			ArticleServices: &articledomain.ArticleService{
				ArticleRepository: articleRepositoy,
			},
		},
	)

	tags.CreateTagsController(
		r,
		tags_domain.CreateTagsService(tags.CreateTagsRepository(db)),
	)

	r.Run()

}
