package service

import (
	//"mysql/entity"
	"mysql/model"
	"mysql/resource"
)

func GetReviewer(id uint) (qu model.Reviewer, err error) {
	err = resource.GetDB().Where("id = ?", id).Find(&qu).Error
	return
}

func SelectNameReviewer(id uint) (qu model.Reviewer, err error) {
	err = resource.GetDB().Select("name,staff_id").Where("id =?", id).Find(&qu).Error
	return
}

func UpdateReviewer(id uint) (err error) {
	err = resource.GetDB().Model(&model.Reviewer{}).Where("id = ?", id).Update("cellphone", "12").Error
	return
}

func MapUpdateReviewer(id uint, m map[string]interface{})(err error){
	err = resource.GetDB().Model(&model.Reviewer{}).Where("id = ?",id).Update(m).Error
	return
}
