package dao

import (
	"encoding/json"
	"sync"

	"github.com/YogeLiu/medical/model"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	orderDao  *OrderDao
	onceOrder sync.Once
)

type OrderDao struct {
	db *leveldb.DB
}

func NewOrderDao(path string) *OrderDao {
	var err error
	var db *leveldb.DB
	onceOrder.Do(func() {
		db, err = leveldb.OpenFile(path, nil)
		if err != nil {
			panic(err)
		}
	})
	orderDao = &OrderDao{db: db}
	return orderDao
}

func (dao *OrderDao) Create(order *model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}
	return dao.db.Put([]byte(order.ID), data, nil)
}

func (dao *OrderDao) Get(id string) (*model.Order, error) {
	data, err := dao.db.Get([]byte(id), nil)
	if err != nil {
		return nil, err
	}
	var order model.Order
	err = json.Unmarshal(data, &order)
	return &order, err
}

func (dao *OrderDao) List(page, pageSize int) ([]*model.Order, error) {
	iter := dao.db.NewIterator(nil, nil)
	var orders []*model.Order
	for iter.Next() {
		var order model.Order
		err := json.Unmarshal(iter.Value(), &order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}
