package dao

import (
	"encoding/json"
	"sync"

	"github.com/YogeLiu/medical/model"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	medicineDao  *MedicineDAO
	onceMedicine sync.Once
)

type MedicineDAO struct {
	db *leveldb.DB
}

func NewMedicineDAO() *MedicineDAO {
	var err error
	var db *leveldb.DB
	path := "./leveldb/medicine.db"
	onceMedicine.Do(func() {
		db, err = leveldb.OpenFile(path, nil)
		if err != nil {
			panic(err)
		}
	})
	medicineDao = &MedicineDAO{db: db}
	return medicineDao
}

func (dao *MedicineDAO) Create(medicine *model.Medicine) error {
	data, err := json.Marshal(medicine)
	if err != nil {
		return err
	}
	return dao.db.Put([]byte(medicine.Name), data, nil)
}

func (dao *MedicineDAO) Get(name string) (*model.Medicine, error) {
	data, err := dao.db.Get([]byte(name), nil)
	if err != nil {
		return nil, err
	}
	var medicine model.Medicine
	err = json.Unmarshal(data, &medicine)
	return &medicine, err
}

func (dao *MedicineDAO) Update(newMed *model.Medicine, oldName string) error {
	data, err := json.Marshal(newMed)
	if err != nil {
		return err
	}
	batch := &leveldb.Batch{}
	batch.Delete([]byte(oldName))
	batch.Put([]byte(newMed.Name), data)
	return dao.db.Write(batch, nil)
}

func (dao *MedicineDAO) Delete(name string) error {
	return dao.db.Delete([]byte(name), nil)
}

func (dao *MedicineDAO) List(page, pageSize int) ([]*model.Medicine, error) {
	iter := dao.db.NewIterator(nil, nil)
	var medicines []*model.Medicine
	for iter.Next() {
		var medicine model.Medicine
		err := json.Unmarshal(iter.Value(), &medicine)
		if err != nil {
			return nil, err
		}
		medicines = append(medicines, &medicine)
	}
	return medicines, nil
}
