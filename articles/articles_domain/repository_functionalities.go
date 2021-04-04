package articles_domain

type ArticlesRepository interface {
	GetData(af *ArticlesFilter) ([]*Article, error)
	SaveData(amr *Article) (*Article, error)
}
