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

func (s *Server) CreateItem(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	item := models.Item{}

	item.Prepare()
	err = json.Unmarshal(body, &item)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = item.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	itemCreated, err := item.CreateItem(s.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, item.ID))
	webResponse := responses.GenerateResponse(http.StatusCreated, "New item created !", itemCreated)
	responses.JSON(w, http.StatusCreated, webResponse)
}

func (s *Server) GetItems(w http.ResponseWriter, _ *http.Request) {
	item := models.Item{}

	items, err := item.FindAllItem(s.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	webResponse := responses.GenerateResponse(http.StatusOK, "Item(s) found !", items)
	responses.JSON(w, http.StatusOK, webResponse)
}

func (s *Server) GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	item := models.Item{}
	itemFound, err := item.FindItemByID(s.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	webResponse := responses.GenerateResponse(http.StatusOK, "Item found !", itemFound)
	responses.JSON(w, http.StatusOK, webResponse)
}

func (s *Server) UpdateItem(w http.ResponseWriter, r *http.Request) {
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

	item := models.Item{}
	err = json.Unmarshal(body, &item)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = item.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	itemUpdate, err := item.UpdateItem(s.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	webResponse := responses.GenerateResponse(http.StatusOK, "Item updated !", itemUpdate)
	responses.JSON(w, http.StatusOK, webResponse)
}
