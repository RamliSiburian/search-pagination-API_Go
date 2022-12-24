package repositories

import (
	"MisterAladin/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(article models.Article) (models.Article, error)
	FindArticle() ([]models.Article, error)
	GetArticleById(ID int) (models.Article, error)
	GetArticleByUser(UserID string) ([]models.Article, error)
	SearchArticle(param string) ([]models.Article, error)
	UpdateArticle(article models.Article) (models.Article, error)
	DeleteArticle(article models.Article) (models.Article, error)
}

type articles struct {
	db *gorm.DB
}

func RepositoryArticle(db *gorm.DB) *articles {
	return &articles{db}
}

func (r *articles) CreateArticle(article models.Article) (models.Article, error) {
	err := r.db.Preload("User").Create(&article).Error

	return article, err
}

func (r *articles) FindArticle() ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Order("id desc").Preload("User").Find(&articles).Error

	return articles, err
}

func (r *articles) GetArticleById(ID int) (models.Article, error) {
	var article models.Article

	err := r.db.First(&article, ID).Error

	return article, err
}
func (r *articles) GetArticleByUser(UserID string) ([]models.Article, error) {
	var article []models.Article
	err := r.db.Where("user_id=?", UserID).Order("id desc").Preload("User").Find(&article).Error

	return article, err
}

func (r *articles) SearchArticle(param string) ([]models.Article, error) {
	var article []models.Article
	err := r.db.Where("title LIKE ? ", "%"+param+"%").Or("body LIKE ? ", "%"+param+"%").Find(&article).Error

	return article, err
}

func (r *articles) UpdateArticle(article models.Article) (models.Article, error) {
	err := r.db.Save(&article).Error

	return article, err
}

func (r *articles) DeleteArticle(article models.Article) (models.Article, error) {
	err := r.db.Delete(&article).Error

	return article, err
}
