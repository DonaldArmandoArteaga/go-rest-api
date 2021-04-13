package articles

import (
	"fmt"

	articlesDomain "github.com/DonaldArmandoArteaga/go-rest-api/articles/articles_domain"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func CreateArticleRepository(db *gorm.DB) (*ArticleRepository, error) {

	err := db.AutoMigrate(&ArticleModelRepository{})

	if err != nil {
		return &ArticleRepository{}, err
	}

	return &ArticleRepository{db: db}, nil
}

func (ar *ArticleRepository) GetData(af *articlesDomain.ArticlesFilter) ([]*articlesDomain.Article, error) {

	amr := []*ArticleModelRepository{}
	/*query := ar.db.
		Limit(int(af.Limit)).
		Offset(int(af.Offset)) //.
		//Joins("JOIN tags ON tags.article_model_repository = articles.id")

	if af.Author != "" {
		query.Where("role = ?", "admin")
	}

	if af.Favorited != "" {
		query.Where("role = ?", "admin")
	}

	if af.Tag != "" {
		query.Where("role = ?", "admin")
	}

	query.Find(&amr)
	if query.Error != nil {
		return []*articlesDomain.Article{}, query.Error
	}
	*/

	a := &ArticleModelRepository{}
	err := ar.db.Model(a).Association("TagList").Count() //.Find(&TagListRepository{})

	fmt.Println(err)
	fmt.Println(amr)

	articles := []*articlesDomain.Article{}
	for _, article := range amr {
		articleTemp, articleError := ArticleModelRepositoryToArticle(article)
		if articleError != nil {
			return []*articlesDomain.Article{}, articleError
		}
		articles = append(articles, articleTemp)
	}

	return articles, nil
}

func (ar *ArticleRepository) SaveData(
	article *articlesDomain.Article,
) (*articlesDomain.Article, error) {

	amr := ArticleToArticleModelRepository(article)
	tx := ar.db.Create(amr)

	if tx.Error != nil {
		return &articlesDomain.Article{}, tx.Error
	}

	article, err := ArticleModelRepositoryToArticle(amr)

	if err != nil {
		return &articlesDomain.Article{}, err
	}

	return article, nil

}
