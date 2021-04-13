package tags

import "github.com/DonaldArmandoArteaga/go-rest-api/tags/tags_domain"

func TagsToTagsDTO(tgd []*tags_domain.Tag) *TagsDTO {

	tags := []string{}

	for _, tag := range tgd {
		tags = append(tags, tag.GetTag())
	}

	return &TagsDTO{
		Tags: tags,
	}
}

func TagsDTOToTags(tgdto *TagsDTO) []*tags_domain.Tag {

	tags := []*tags_domain.Tag{}

	for _, tag := range tgdto.Tags {
		tags = append(tags, tags_domain.CreateTags(tag))
	}

	return tags
}

func TagsModelRepositoryToTags(tgmr *TagsModelRepository) *tags_domain.Tag {
	return tags_domain.CreateTags(tgmr.Value)
}

func TagToTagsModelRespository(tag *tags_domain.Tag) *TagsModelRepository {
	return &TagsModelRepository{
		ID:    0,
		Value: tag.GetTag(),
	}
}
