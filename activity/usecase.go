package activity

import (
	"context"
	"time"
)

type ActivityUsecase interface{
	GetAllActivity(c context.Context) (activities []Activity, err error)
	GetActivityById(c context.Context, id int) (activity Activity, err error)
	CreateActivity(c context.Context, payload Payload) (activity Activity, err error)
	UpdateActivity(c context.Context, id int, payload Payload) (activity Activity, err error)
	DeleteActivity(c context.Context, id int) (status string, err error)
}

type activityUsecase struct{
	activityRepository ActivityRepository
}

func NewActivityUsecase(activityRepository ActivityRepository) ActivityUsecase{
	return &activityUsecase{activityRepository}
}

func (uc *activityUsecase) GetAllActivity(c context.Context) (activities []Activity, err error){
	activities, err = uc.activityRepository.GetAll(c)

	return activities, err
}

func (uc *activityUsecase) GetActivityById(c context.Context, id int) (activity Activity, err error){
	activity, err = uc.activityRepository.GetById(c, id)

	return activity, err
}

func (uc *activityUsecase) CreateActivity(c context.Context, payload Payload) (activity Activity, err error){
	activity, err = uc.activityRepository.Create(c, activity)

	return activity, err
}

func (uc *activityUsecase) UpdateActivity(c context.Context, id int, payload Payload) (activity Activity, err error){
	activity.Title = payload.Title
	activity.Email = payload.Email
	activity.UpdatedAt = time.Now()

	activity, err = uc.activityRepository.Update(c, id, activity)

	return activity, err	
}

func (uc *activityUsecase) DeleteActivity(c context.Context, id int) (status string, err error){
	status, err = uc.activityRepository.Delete(c, id)

	return status, err
}