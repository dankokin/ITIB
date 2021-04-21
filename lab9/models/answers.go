package models

import "itib/lab9/cluster"

type MessageAnswer struct {
	Status  int    `json:"status,int"`
	Message string `json:"message,string"`
}

type AddAreaAnswer struct {
	Status  int               `json:"status,int"`
	Message string            `json:"message,string"`
	Data    AddAreaAnswerData `json:"data"`
}

type TrainAnswer struct {
	Status  int             `json:"status,int"`
	Message string          `json:"message,string"`
	Data    TrainAnswerData `json:"data"`
}

type AreaAnswer struct {
	Status  int               `json:"status,int"`
	Message string            `json:"message,string"`
	Data    GetAreaAnswerData `json:"data"`
}

func GetSuccessAnswer(message string) *MessageAnswer {
	return &MessageAnswer{
		Status:  100,
		Message: message,
	}
}

func GetAddAreaAnswer(data *AddAreaAnswerData) *AddAreaAnswer {
	return &AddAreaAnswer{
		Status:  101,
		Message: "ok",
		Data:    *data,
	}
}

func GetTrainAnswer(data *TrainAnswerData) *TrainAnswer {
	return &TrainAnswer{
		Status:  102,
		Message: "ok",
		Data:    *data,
	}
}

func GetAreaAnswer(data *GetAreaAnswerData) *AreaAnswer {
	return &AreaAnswer{
		Status:  103,
		Message: "ok",
		Data:    *data,
	}
}

func GetErrorAnswer(error string) *MessageAnswer {
	return &MessageAnswer{
		Status:  200,
		Message: error,
	}
}

var IncorrectJsonAnswer = MessageAnswer{
	Status:  201,
	Message: "incorrect JSON",
}

var IncorrectRequestAnswer = MessageAnswer{
	Status:  202,
	Message: "incorrect request",
}

type AddPointData struct {
	Id     int             `json:"id"`
	Points []cluster.Point `json:"points"`
}

type AddClusterData struct {
	Id       int             `json:"id"`
	Clusters []cluster.Point `json:"clusters"`
}

type TrainData struct {
	Id                 int    `json:"id"`
	MaxIterations      uint64 `json:"max_age"`
	StepByStep         bool   `json:"by_step"`
	DistanceFunctionId int    `json:"dist_id"`
}

type GetAreaData struct {
	Id         int `json:"id"`
	DistFuncId int `json:"dist_id"`
}

type ClearAreaData struct {
	Id int `json:"id"`
}

type AddAreaAnswerData struct {
	Id int `json:"id"`
}

type GetAreaAnswerData struct {
	Clusters []cluster.Cluster `json:"clusters"`
}

type TrainAnswerData struct {
	Finished bool              `json:"finished"`
	Clusters []cluster.Cluster `json:"clusters"`
}
