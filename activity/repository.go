package activity

import (
	"context"
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type ActivityRepository interface{
	GetAll(c context.Context) (activities []Activity, err error)
	GetById(c context.Context, id int) (activity Activity, err error)
	Create(c context.Context, payload Activity) (activity Activity, err error)
	Update(c context.Context, id int, payload Activity) (activty Activity, err error)
	Delete(c context.Context, id int) (status string, err error)
}

type activityRepository struct{
	Conn *gorm.DB
}

func NewActivityRepository(Conn *gorm.DB) ActivityRepository{
	return &activityRepository{Conn}
}

func (repo *activityRepository) GetAll(c context.Context) (activities []Activity, err error){
	err = repo.Conn.Find(&activities).Error

	if (err != nil){
		return []Activity{}, err
	}
	
	return activities, nil
}

func (repo *activityRepository) GetById(c context.Context, id int) (activity Activity, err error){
	err = repo.Conn.Where("id = ?", id).First(&activity).Error

	if (err != nil){
		return Activity{}, err
	}

	return activity, err
}

func (repo *activityRepository) Create(c context.Context, payload Activity) (activty Activity, err error){
	activty = payload

	err = repo.Conn.Create(&activty).Error

	if err != nil{
		return Activity{}, err		
	}

	return activty, nil
}

func (repo *activityRepository) Update(c context.Context, id int, payload Activity) (activity Activity, err error){
	err = repo.Conn.First(&activity, id).Error

	if err != nil{
		return Activity{}, err
	}

	activity.Title = payload.Title
	activity.UpdatedAt = time.Now()

	err = repo.Conn.Save(&activity).Error

	if err != nil{
		return Activity{}, err
	}

	return activity, nil
}

func (repo *activityRepository) Delete(c context.Context, id int) (status string, err error){
	var activity Activity
	err = repo.Conn.First(&activity, id).Error

	idString := strconv.Itoa(id)

	if err != nil{
		return "Not Found", errors.New("Activity with ID "+idString+" Not Found")
	}

	err = repo.Conn.Delete(&activity, id).Error

	if err != nil{
		return "", err
	}

	return "", nil
}