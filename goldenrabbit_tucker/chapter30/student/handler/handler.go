package student_handler

import (
	"encoding/json"
	"fmt"
	"gt/chap30/ex/student"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func List(response http.ResponseWriter, request *http.Request) {
	buff, err := json.Marshal(student.Manager.FindAll())
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("student list fail."))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(buff)
}

func Read(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	student, err := student.Manager.Find(id)

	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(err.Error()))
		return
	}

	buff, err := json.Marshal(student)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("student list fail."))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write(buff)
}

func Regist(response http.ResponseWriter, request *http.Request) {

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("body read fail!!!"))
		return
	}

	studentInfo := new(student.Student)

	if err := json.Unmarshal(body, studentInfo); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(fmt.Sprintf("body bind fail. body: %s", body)))
		return
	}

	if err := student.Manager.Insert(studentInfo); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(fmt.Sprintf("student insert fail. error: %s", err.Error())))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("ok"))
}

func Modify(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("body read fail!!!"))
		return
	}

	studentInfo := new(student.Student)

	if err := json.Unmarshal(body, studentInfo); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(fmt.Sprintf("body bind fail. body: %s", body)))
		return
	}

	studentInfo.Id = id

	if err := student.Manager.Update(studentInfo); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(fmt.Sprintf("student update fail. error: %s", err.Error())))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("ok"))
}

func Destroy(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	err := student.Manager.Delete(id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(err.Error()))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("ok"))
}
