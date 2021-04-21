package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"itib/lab9/models"
	"itib/lab9/store"
	"itib/lab9/utils"
)

const (
	defaultMaxIterations    = 100
	defaultDistanceFunction = 1
)

func AddArea(w http.ResponseWriter, r *http.Request) {
	id, err := store.AddArea()
	if err != nil {
		utils.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
		return
	}

	utils.SendJson(w, http.StatusOK, models.GetAddAreaAnswer(&models.AddAreaAnswerData{Id: id}))
}

func AddPoint(w http.ResponseWriter, r *http.Request) {
	var target models.AddPointData
	err := json.NewDecoder(r.Body).Decode(&target)
	defer r.Body.Close()
	if err != nil {
		utils.SendJson(w, http.StatusInternalServerError, models.IncorrectJsonAnswer)
		return
	}

	area, err := store.GetArea(target.Id)
	if err != nil {
		utils.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
		return
	}

	for i := range target.Points {
		area.AddPoint(target.Points[i])
	}

	utils.SendJson(w, http.StatusOK, models.GetSuccessAnswer("ok"))
}

func AddCluster(w http.ResponseWriter, r *http.Request) {
	var target models.AddClusterData
	err := json.NewDecoder(r.Body).Decode(&target)
	defer r.Body.Close()
	if err != nil {
		utils.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
		return
	}

	area, err := store.GetArea(target.Id)
	if err != nil {
		utils.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
		return
	}

	for i := range target.Clusters {
		area.AddCluster(target.Clusters[i])
	}

	utils.SendJson(w, http.StatusOK, models.GetSuccessAnswer("ok"))
}

func Learn(w http.ResponseWriter, r *http.Request) {
	var target models.TrainData
	err := json.NewDecoder(r.Body).Decode(&target)
	defer r.Body.Close()
	if err != nil {
		utils.SendJson(w, http.StatusInternalServerError, models.IncorrectJsonAnswer)
		return
	}

	if target.MaxIterations == 0 {
		target.MaxIterations = defaultMaxIterations
	}

	if target.DistanceFunctionId == 0 {
		target.DistanceFunctionId = defaultDistanceFunction
	}

	area, err := store.GetArea(target.Id)

	if err != nil {
		utils.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
		return
	}

	distanceFunction, err := store.GetDistanceFunction(target.DistanceFunctionId)
	if err != nil {
		utils.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
		return
	}

	var isFinished bool
	if target.StepByStep {
		isFinished = area.DoStep(distanceFunction)
	} else {
		isFinished = area.Learn(distanceFunction, target.MaxIterations)
	}

	clusters := area.GetClusters(distanceFunction)
	utils.SendJson(w, http.StatusOK, models.GetTrainAnswer(&models.TrainAnswerData{
		Finished: isFinished,
		Clusters: clusters,
	}))
}

func JsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(res, req)
	})
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	homePage, err := ioutil.ReadFile("static/main_page.html")
	if err != nil {
		utils.SendText(w, http.StatusInternalServerError, "Home page is not opened")
	}
	utils.SendText(w, http.StatusOK, string(homePage))
}

func ApiGetArea(w http.ResponseWriter, r *http.Request) {
	strAreaId, exists := mux.Vars(r)["id"]
	if !exists {
		utils.SendJson(w, http.StatusBadRequest, models.IncorrectRequestAnswer)
		return
	}
	areaId, err := strconv.Atoi(strAreaId)
	if err != nil {
		utils.SendJson(w, http.StatusBadRequest, models.IncorrectRequestAnswer)
		return
	}

	distanceFunctionId := defaultDistanceFunction
	strFunctionId, exists := mux.Vars(r)["dist_id"]
	if exists {
		distanceFunctionId, err = strconv.Atoi(strFunctionId)
		if err != nil {
			utils.SendJson(w, http.StatusBadRequest, models.IncorrectRequestAnswer)
			return
		}
	}

	area, err := store.GetArea(areaId)
	if err != nil {
		utils.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
		return
	}

	distanceFunction, err := store.GetDistanceFunction(distanceFunctionId)
	if err != nil {
		utils.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
		return
	}

	utils.SendJson(w, http.StatusOK, models.GetAreaAnswer(&models.GetAreaAnswerData{
		Clusters: area.GetClusters(distanceFunction),
	}))
}

func ApiClearArea(w http.ResponseWriter, r *http.Request) {
	var target models.ClearAreaData
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &target)
	if err != nil {
		utils.SendJson(w, http.StatusInternalServerError, models.IncorrectJsonAnswer)
		return
	}
	defer r.Body.Close()

	area, err := store.GetArea(target.Id)
	if err != nil {
		fmt.Println(err)
		utils.SendJson(w, http.StatusInternalServerError, models.IncorrectJsonAnswer)
		return
	}
	area.Clear()

	utils.SendJson(w, http.StatusOK, models.GetSuccessAnswer("ok"))
}

func ApiHomeHandler(w http.ResponseWriter, _ *http.Request) {
	utils.SendJson(w, http.StatusOK, models.GetSuccessAnswer("Main page!"))
}
