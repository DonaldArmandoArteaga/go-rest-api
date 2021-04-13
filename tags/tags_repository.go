package tags

import (
	"github.com/DonaldArmandoArteaga/go-rest-api/tags/tags_domain"
	"gorm.io/gorm"
)

type TagsRepository struct {
	db *gorm.DB
}

func CreateTagsRepository(db *gorm.DB) *TagsRepository {
	return &TagsRepository{
		db: db,
	}
}

func (tgr *TagsRepository) GetTags() []*tags_domain.Tag {

	tgmr := []*TagsModelRepository{}
	tags := []*tags_domain.Tag{}
	tgr.db.Find(&tgmr)

	for _, tag := range tgmr {
		tags = append(tags, TagsModelRepositoryToTags(tag))
	}

	return tags
}

func (tgr *TagsRepository) CreateTags(tags []*tags_domain.Tag) ([]*tags_domain.Tag, error) {

	tgmr := []*TagsModelRepository{}

	for _, tag := range tags {
		tgmr = append(tgmr, TagToTagsModelRespository(tag))
	}

	tx := tgr.db.Create(tgmr)

	if tx.Error != nil {
		return []*tags_domain.Tag{}, tx.Error
	}

	return tags, nil

}
