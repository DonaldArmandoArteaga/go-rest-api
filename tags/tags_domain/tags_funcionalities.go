package tags_domain

type TagsFunctionalities interface {
	GetTags() []*Tag
	CreateTags(tags []*Tag) ([]*Tag, error)
}
