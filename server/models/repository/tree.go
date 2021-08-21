package repository

import (
	"../domain"
	"gorm.io/gorm"
)

type TreeRepository struct {
	Conn *gorm.DB
}

func NewTreeRepository(conn *gorm.DB) *TreeRepository {
	return &TreeRepository{Conn: conn}
}

func (repo *TreeRepository) Create(tree domain.Tree) error {
	return repo.Conn.Create(&tree).Error
}

func (repo *TreeRepository) FindParent(name string) (domain.Version, error) {
	var parent domain.Version

	subQuery := repo.Conn.Model(&domain.Tree{}).Where("name = ?", name).Select("parent_name")
	err := repo.Conn.Where("name = (?)", subQuery).First(&parent).Error
	return parent, err
}

func (repo *TreeRepository) ListChildren(name string) ([]domain.Version, error) {
	var children []domain.Version

	subQuery := repo.Conn.Model(&domain.Tree{}).Where("parent_name = ?", name).Select("name")
	err := repo.Conn.Where("name IN (?)", subQuery).Find(&children).Error
	return children, err
}

func (repo *TreeRepository) ListTreeBFS(name string) ([]domain.TreeMinimized, error) {
	var tree []domain.TreeMinimized
	err := repo.Conn.Raw("SELECT * FROM list_tree_bfs(?)", name).Scan(&tree).Error
	return tree, err

}

func (repo *TreeRepository) Delete(name string) error {
	return repo.Conn.Where("name = ?", name).Delete(&domain.Tree{}).Error
}
