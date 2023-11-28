package service

import (
	"fmt"
	"strconv"

	"github.com/YogeLiu/medical/dao"
	"github.com/YogeLiu/medical/log"
	"github.com/YogeLiu/medical/model"
)

type MedicineService struct {
	dao *dao.MedicineDAO
}

func NewDetailService() *MedicineService {
	return &MedicineService{dao: dao.NewMedicineDAO()}
}

func (s *MedicineService) AddDetail(detail *model.Medicine) error {
	medical, err := s.dao.Get(detail.Name)
	if err != nil {
		log.Log.Errorf("get medicine error: %+v", err.Error())
		return err
	}
	if medical != nil || medical.Name != "" {
		log.Log.Error("medicine already exist")
		return fmt.Errorf("medicine already exist")
	}
	err = s.dao.Create(detail)
	if err != nil {
		log.Log.Errorf("create medicine error: %+v", err.Error())
	}
	return err
}

func (s *MedicineService) ListDetail(page, pageSize int) (list *model.MedicineList, err error) {
	medicines, err := s.dao.List(page, pageSize)
	if err != nil {
		log.Log.Errorf("list medicine error: %+v", err.Error())
		return nil, err
	}
	start := (page - 1) * pageSize
	for idx, medicine := range medicines {
		medicine.ID = strconv.Itoa(start + idx + 1)
	}
	list = &model.MedicineList{
		Medicines: medicines,
		HasMore:   len(medicines) == pageSize,
	}
	return list, nil
}

func (s *MedicineService) GetDetailByName(name string) (*model.Medicine, error) {
	return s.dao.Get(name)
}

func (s *MedicineService) UpdateDetail(new *model.Medicine, oldName string) error {
	oldMedi, err := s.dao.Get(oldName)
	if err != nil {
		log.Log.Errorf("get medicine error: %+v", err.Error())
		return err
	}
	if new.Amount == 0 {
		new.Amount = oldMedi.Amount
	}
	if new.Name == "" {
		new.Name = oldName
	}
	if new.Price == 0.0 {
		new.Price = oldMedi.Price
	}
	err = s.dao.Update(new, oldName)
	if err != nil {
		log.Log.Errorf("update medicine error: %+v", err.Error())
	}
	return err
}

func (s *MedicineService) DeleteDetail(name string) error {
	err := s.dao.Delete(name)
	if err != nil {
		log.Log.Errorf("delete medicine error: %+v", err.Error())
	}
	return err

}
