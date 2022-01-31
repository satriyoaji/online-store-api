package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"online-store-evermos/src/models"
	"online-store-evermos/src/responses"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (s *Server) AddOrder(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	order := models.Order{}

	order.Prepare()
	err = json.Unmarshal(body, &order)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = order.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	orderCreated, err := order.AddOrder(s.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, order.ID))
	webResponse := responses.GenerateResponse(http.StatusCreated, "New order created !", orderCreated)
	responses.JSON(w, http.StatusCreated, webResponse)
}

func (s *Server) GetOrderByUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	order := models.Order{}
	itemFound, err := order.FindAllOrderByUID(s.DB, pid)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}

	webResponse := responses.GenerateResponse(http.StatusOK, "Order item found !", itemFound)
	responses.JSON(w, http.StatusOK, webResponse)
}
