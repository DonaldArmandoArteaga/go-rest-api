package tags

func (TagsModelRepository) TableName() string {
	return "tags"
}

type TagsModelRepository struct {
	ID    uint64 `gorm:"primaryKey"`
	Value string
}
