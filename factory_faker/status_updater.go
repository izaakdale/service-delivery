package factoryfaker

import (
	"math/rand"
	"time"

	"github.com/izaakdale/service-delivery/dao"
	"github.com/izaakdale/service-delivery/model/delivery"
	"github.com/izaakdale/utils-go/logger"
)

func UpdateStatuses(in *delivery.DeliveryRequest) {
	// defer wg.Done()
	logger.Info("Updating status in factory")
	rand.Seed(time.Now().UnixNano())

	interval := rand.Intn(30)
	time.Sleep(time.Second * time.Duration(interval))

	logger.Info("Updating order status to out for delivery")
	dao.UpdateStatus(&dao.DeliveryRecord{
		OrderId:    dao.QuotePrefixPK + in.OrderId,
		RecordType: dao.StatusSK,
		Status:     delivery.ORDER_STATUS_name[int32(delivery.ORDER_STATUS_OUT_FOR_DELIVERY)],
		Address:    in.Address,
	})

	interval2 := rand.Intn(30)
	time.Sleep(time.Second * time.Duration(interval2))

	logger.Info("Updating order status to delivered")
	dao.UpdateStatus(&dao.DeliveryRecord{
		OrderId:    "ORDER#" + in.OrderId,
		RecordType: dao.StatusSK,
		Status:     delivery.ORDER_STATUS_name[int32(delivery.ORDER_STATUS_DELIVERED)],
		Address:    in.Address,
	})
}
