package articles_domain

import "time"

type ArticleServices interface {
	GetAll(af *ArticlesFilter) ([]*Article, error)
	Save(ar *Article) (*Article, error)
}

type ArticlesFilter struct {
	Tag       string
	Author    string
	Favorited string
	Limit     uint64
	Offset    uint64
}

type ArticleService struct {
	ArticleRepository ArticlesRepository
}

func (as *ArticleService) GetAll(af *ArticlesFilter) ([]*Article, error) {

	if af.Limit == 0 || af.Limit > 100 {
		af.Limit = 20
	}

	articles, err := as.ArticleRepository.GetData(af)

	if err != nil {
		return []*Article{}, err
	}

	return articles, nil
}

func (as *ArticleService) Save(ar *Article) (*Article, error) {
	ar.createdAt = time.Now()
	return as.ArticleRepository.SaveData(ar)
}
