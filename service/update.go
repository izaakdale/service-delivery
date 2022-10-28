package service

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/izaakdale/service-delivery/dao"
	"github.com/izaakdale/service-delivery/model/delivery"
	"github.com/izaakdale/utils-go/response"
)

func UpdateDeliveryStatusHandler(w http.ResponseWriter, r *http.Request) {
	var req delivery.UpdateDeliveryStatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.NewBadRequestError("Could not decode the inputted values"))
		return
	}

	status, err := sanitizeStatus(req.Status)
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.NewBadRequestError(err.Error()))
		return
	}

	dao.UpdateStatus(&dao.StatusRecord{
		OrderId:    dao.OrderPrefixPK + req.OrderId,
		RecordType: dao.StatusSK,
		Status:     status,
	})

	response.WriteJson(w, http.StatusAccepted, response.Response{
		Code:    http.StatusAccepted,
		Message: "Updated!",
	})
}

func sanitizeStatus(status string) (string, error) {
	_, ok := delivery.ORDER_STATUS_value[status]
	if !ok {
		return "", errors.New("Incorrect value for status")
	}
	return status, nil
}
