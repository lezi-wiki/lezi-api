package model

import (
	"encoding/gob"
	"github.com/lezi-wiki/lezi-api/pkg/cache"
	"gorm.io/gorm"
	"strings"
)

type Namespace struct {
	gorm.Model
	Name string `gorm:"not null;unique_index"`

	Texts []Text `gorm:"foreignKey:Namespace;references:Name"`
}

func init() {
	gob.Register(&Namespace{})
}

type NamespaceService interface {
	Get(namespace Namespace) (*Namespace, error)
	List(namespace Namespace) ([]Namespace, error)
	ListAll() ([]Namespace, error)
	CreateNamespace(namespace Namespace) (*Namespace, error)
	UpdateNamespace(id uint, namespace Namespace) (*Namespace, error)
	DeleteNamespace(id uint) error
	Count() int64
	Exists(namespace Namespace) bool
}

type NamespaceServiceImpl struct {
	db *gorm.DB
}

func NewNamespaceService(db *gorm.DB) NamespaceService {
	return &NamespaceServiceImpl{db: db}
}

func (n *NamespaceServiceImpl) Get(namespace Namespace) (*Namespace, error) {
	key := "namespace_name_" + strings.ToLower(namespace.Name)
	c, exist := cache.Get(key)
	if exist {
		if namespace, ok := c.(Namespace); ok {
			return &namespace, nil
		}
	}

	err := n.db.Model(&Namespace{}).Preload("Texts").Where(&namespace).First(&namespace).Error
	if err != nil {
		return nil, err
	}

	_ = cache.Set(key, namespace, 0)

	return &namespace, nil
}

func (n *NamespaceServiceImpl) List(namespace Namespace) ([]Namespace, error) {
	var rows []Namespace
	err := n.db.Model(&Namespace{}).Where(&namespace).Find(&rows).Error
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (n *NamespaceServiceImpl) ListAll() ([]Namespace, error) {
	return n.List(Namespace{})
}

func (n *NamespaceServiceImpl) CreateNamespace(namespace Namespace) (*Namespace, error) {
	err := n.db.Model(&Namespace{}).Create(&namespace).Error
	if err != nil {
		return nil, err
	}

	return &namespace, nil
}

func (n *NamespaceServiceImpl) UpdateNamespace(id uint, namespace Namespace) (*Namespace, error) {
	err := n.db.Model(&Namespace{}).Where("id = ?", id).Updates(&namespace).Error
	if err != nil {
		return nil, err
	}

	return &namespace, nil
}

func (n *NamespaceServiceImpl) DeleteNamespace(id uint) error {
	var ns Namespace
	err := n.db.Model(&Namespace{}).Where("id = ?", id).Delete(&ns).Error
	if err != nil {
		return err
	}

	_ = cache.Delete("namespace_name_" + ns.Name)

	return nil
}

func (n *NamespaceServiceImpl) Count() int64 {
	var count int64
	n.db.Model(&Namespace{}).Count(&count)
	return count
}

func (n *NamespaceServiceImpl) Exists(namespace Namespace) bool {
	var count int64
	n.db.Model(&Namespace{}).Where(&namespace).Count(&count)
	return count > 0
}
