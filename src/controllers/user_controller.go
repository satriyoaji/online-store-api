package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"online-store-evermos/src/formaterror"
	"online-store-evermos/src/models"
	"online-store-evermos/src/responses"
	"strconv"
)

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}

	user.Prepare()
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userCreated, err := user.SaveUser(s.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	webResponse := responses.GenerateResponse(http.StatusCreated, "New user created !", userCreated)
	responses.JSON(w, http.StatusCreated, webResponse)
}

func (s *Server) GetUsers(w http.ResponseWriter, _ *http.Request) {
	user := models.User{}

	users, err := user.FindAllUser(s.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	webResponse := responses.GenerateResponse(http.StatusOK, "User(s) found !", users)
	responses.JSON(w, http.StatusOK, webResponse)
}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userFound, err := user.FindUserById(s.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	webResponse := responses.GenerateResponse(http.StatusOK, "User found !", userFound)
	responses.JSON(w, http.StatusOK, webResponse)
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedUser, err := user.UpdateUser(s.DB, uid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	webResponse := responses.GenerateResponse(http.StatusOK, "User updated !", updatedUser)
	responses.JSON(w, http.StatusOK, webResponse)
}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user := models.User{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = user.DeleteUser(s.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	webResponse := responses.GenerateResponse(http.StatusOK, "User deleted !", "")
	responses.JSON(w, http.StatusNoContent, webResponse)
}
