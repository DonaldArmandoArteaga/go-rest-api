package tags_domain

type Tag struct {
	value string
}

func CreateTags(tag string) *Tag {
	return &Tag{
		value: tag,
	}
}

func (t *Tag) GetTag() string {
	return t.value
}
