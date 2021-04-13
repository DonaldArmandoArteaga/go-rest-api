package tags

import (
	"net/http"

	"github.com/DonaldArmandoArteaga/go-rest-api/tags/tags_domain"
	"github.com/gin-gonic/gin"
)

type TagsController struct {
	g           *gin.Engine
	tagsService *tags_domain.TagsService
}

func CreateTagsController(
	gin *gin.Engine,
	tagsService *tags_domain.TagsService,
) {

	tc := TagsController{
		g:           gin,
		tagsService: tagsService,
	}

	tc.g.GET("/tags", tc.getAllTags)
	tc.g.POST("/tags", tc.saveTags)

}

func (tc *TagsController) getAllTags(c *gin.Context) {

	tags, tagsError := tc.tagsService.GetTags()

	if tagsError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tagsError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tags": TagsToTagsDTO(tags)})

}

func (tc *TagsController) saveTags(c *gin.Context) {

	tagsdto := &TagsDTO{}

	if bindError := c.ShouldBind(tagsdto); bindError == nil {

		tags, err := tc.tagsService.CreateTags(TagsDTOToTags(tagsdto))

		if len(tags) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"tags": TagsToTagsDTO(tags), "error": err.Error()})
		}

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": bindError.Error()})
	}

}
