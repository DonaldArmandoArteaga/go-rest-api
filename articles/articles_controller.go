package articles

import (
	"bytes"
	"net/http"
	"strconv"

	articledomain "github.com/DonaldArmandoArteaga/go-rest-api/articles/articles_domain"
	"github.com/gin-gonic/gin"
)

type ArticlesController struct {
	Gin             *gin.Engine
	ArticleServices articledomain.ArticleServices
}

func CreateArticlesController(ac *ArticlesController) {
	ac.Gin.GET("/articles", ac.getAllArticles)
	ac.Gin.POST("/articles", ac.saveArticle)
}

func (ac *ArticlesController) getAllArticles(c *gin.Context) {

	var buffer bytes.Buffer

	limit, parseLimitErr := strconv.ParseUint(c.DefaultQuery("limit", "0"), 10, 64)

	if parseLimitErr != nil {
		buffer.WriteString("El limit debe ser un numero entero positivo \n")
	}

	offset, parseOffsetErr := strconv.ParseUint(c.DefaultQuery("offset", "0"), 10, 64)

	if parseOffsetErr != nil {
		buffer.WriteString("El offset debe ser un numero entero positivo \n")
	}

	articles, err := ac.ArticleServices.GetAll(
		&articledomain.ArticlesFilter{
			Tag:       c.Query("tag"),
			Author:    c.Query("author"),
			Favorited: c.Query("favorited"),
			Limit:     limit,
			Offset:    offset,
		},
	)

	if buffer.Len() > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": buffer.String()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	articlesDTO := []*ArticleQueryDTO{}

	for _, articleRepository := range articles {
		articlesDTO = append(articlesDTO, ArticleToArticleQueryDTO(articleRepository))
	}

	c.JSON(http.StatusOK, gin.H{"articles": articlesDTO})

}

func (ac *ArticlesController) saveArticle(c *gin.Context) {

	articleCommandDTO := &ArticleCommandDTO{}

	if bindError := c.ShouldBind(articleCommandDTO); bindError == nil {

		articleConvert, errConvert := ArticleDTOToArticle(articleCommandDTO)

		if errConvert != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errConvert.Error()})
		}

		article, err := ac.ArticleServices.Save(articleConvert)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"article": ArticleToArticleQueryDTO(article)})
		}

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": bindError.Error()})
	}

}
