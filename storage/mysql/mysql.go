package mysql

import (
	"errors"
	"extend-lib/storage"
	"fmt"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"

	// gorm framework use it to init mysql connection
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GormStore implement a store by mysql
type GormStore struct {
	conn *gorm.DB
}

// Client return a mysql db connection
func (s *GormStore) Client() interface{} {
	return s.conn
}

// IsExist whether entry is exist
func (s *GormStore) IsExist(id int64, name string, v interface{}) bool {
	if id != 0 {
		return !s.conn.Where("id = ?", id).First(v).RecordNotFound()
	}
	if len(name) == 0 {
		return false
	} else {
		return !s.conn.Where("name = ?", name).First(v).RecordNotFound()
	}
}

// Get get an entry object
func (s *GormStore) Get(id int64, name string, v interface{}) error {
	if id > 0 {
		return s.conn.Where("id = ?", id).First(v).Error
	}
	if name != "" {
		return s.conn.Where("name = ?", name).First(v).Error
	}

	return errors.New("parameter error")
}

// Put create an new entry object
func (s *GormStore) Put(name string, v interface{}) error {
	isExist := s.IsExist(0, name, v)
	if !isExist {
		return s.conn.Create(v).Error
	}
	return errors.New("object exist")
}

// Update update an entry object
func (s *GormStore) Update(id int64, name string, has interface{}, v interface{}) error {
	if id > 0 {
		return s.conn.Model(has).Where("id = ?", id).Update(v).Error
	}
	return s.conn.Model(has).Where("name = ?", name).Update(v).Error
}

// Remove remove an entry object
func (s *GormStore) Remove(id int64, name string, v interface{}) error {
	if id <= 0 && name == "" {
		return errors.New("conditions for error")
	}
	if !s.IsExist(id, name, v) {
		return errors.New("record not found")
	}
	tx := s.conn.Begin()
	if id > 0 {
		if err := tx.Where("id = ?", id).Delete(v).Error; err != nil {
			return err
		}
	}
	if name != "" {
		if err := tx.Where("name = ?", name).Delete(v).Error; err != nil {
			return err
		}
	}
	return tx.Commit().Error
}

// List get a list of given object
func (s *GormStore) List(size, current int, v interface{}) (int, error) {
	var count int
	s.conn.Find(v).Count(&count)
	if s.conn.Error != nil {
		return 0, s.conn.Error
	}
	//if current page is 0, return all entry
	if current == 0 {
		result := s.conn.Find(v).Order("id desc")
		return count, result.Error
	}
	if count == 0 {
		result := s.conn.Offset((current - 1) * size).Limit(size).Find(v).Order("id desc")
		return count, result.Error
	}
	var realPage int
	if count%size > 0 {
		realPage = count/size + 1
	} else {
		realPage = count / size
	}
	if current > realPage {
		return count, errors.New("current page out of total page")
	}
	result := s.conn.Offset((current - 1) * size).Limit(size).Find(v).Order("id desc")
	return count, result.Error
}

// GetAssociation for relational database it will return an entry by 'join' keyword
func (s *GormStore) GetAssociation(association string, v ...interface{}) error {
	if len(v) != 2 {
		return fmt.Errorf("error relate, must give object and relate")
	}
	return s.conn.Model(v[0]).Association(association).Find(v[1]).Error
}

// NewStore return a new Store implement by mysql
func NewStore(uri string) (storage.Store, error) {
	// uri := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Endpoint, cfg.DB)
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		glog.V(0).Infoln(err)
		return nil, err
	}
	return &GormStore{conn: db}, nil
}
