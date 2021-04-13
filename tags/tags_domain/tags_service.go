package tags_domain

import (
	"fmt"
)

type TagsService struct {
	tagsFunctionalities TagsFunctionalities
}

func CreateTagsService(tagsFunctionalities TagsFunctionalities) *TagsService {
	return &TagsService{
		tagsFunctionalities: tagsFunctionalities,
	}
}

func (ts *TagsService) GetTags() ([]*Tag, error) {
	return []*Tag{CreateTags("Hola"), CreateTags("Mundo")}, nil
}

func (ts *TagsService) CreateTags(tags []*Tag) ([]*Tag, error) {

	duplicateTags := []string{}
	unduplicatedTags := []*Tag{}

	flag := false
	for i := 0; i < len(tags); i++ {
		for j := 0; j < len(unduplicatedTags); j++ {

			if tags[i].GetTag() == tags[j].GetTag() {
				flag = true
			}

		}

		if flag {
			duplicateTags = append(duplicateTags, tags[i].GetTag())
		} else {
			unduplicatedTags = append(unduplicatedTags, tags[i])
		}

		flag = false
	}

	saveTags := ts.tagsFunctionalities.GetTags()
	newTags := []*Tag{}
	errorTags := []string{}

	flag = false
	for i := 0; i < len(saveTags); i++ {
		for jj := 0; jj < len(unduplicatedTags); jj++ {

			if saveTags[i].GetTag() == unduplicatedTags[jj].GetTag() {
				flag = true
			}
		}

		if flag {
			errorTags = append(errorTags, saveTags[i].GetTag())
		} else {
			newTags = append(newTags, saveTags[i])
		}

		flag = false

	}

	if len(saveTags) == 0 {
		newTags = unduplicatedTags
	}

	if len(newTags) > 0 {

		createTags, createTagsError := ts.tagsFunctionalities.CreateTags(newTags)

		if createTagsError != nil {
			return []*Tag{}, createTagsError
		}

		return createTags, fmt.Errorf(
			"%s %v %s %v",
			"Tags ya existentes: ", errorTags,
			"Tags duplicados en la solicitud: ", duplicateTags,
		)

	}

	return []*Tag{}, fmt.Errorf(
		"%s %v %s %v",
		"Tags ya existentes: ", errorTags,
		"Tags duplicados en la solicitud: ", duplicateTags,
	)

}
