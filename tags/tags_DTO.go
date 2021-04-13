package tags

type TagsDTO struct {
	Tags []string `json:"tags" binding:"required"`
}
