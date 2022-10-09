package factoryfaker

// UpdateStatuses creates a random interval of time up to 30 seconds
// and updates the status to OUT_FOR_DELIVERY, and then repeats for DELIVERED
// func UpdateStatuses(in *delivery.DeliveryRequest) {
// 	logger.Info("Updating status in factory")
// 	rand.Seed(time.Now().UnixNano())

// 	interval := rand.Intn(30)
// 	time.Sleep(time.Second * time.Duration(interval))

// 	logger.Info("Updating order status to out for delivery")
// 	dao.UpdateStatus(&dao.DeliveryRecord{
// 		OrderId:    dao.QuotePrefixPK + in.OrderId,
// 		RecordType: dao.StatusSK,
// 		Status:     delivery.ORDER_STATUS_name[int32(delivery.ORDER_STATUS_OUT_FOR_DELIVERY)],
// 		Address:    in.Address,
// 	})

// 	interval2 := rand.Intn(30)
// 	time.Sleep(time.Second * time.Duration(interval2))

// 	logger.Info("Updating order status to delivered")
// 	dao.UpdateStatus(&dao.DeliveryRecord{
// 		OrderId:    "ORDER#" + in.OrderId,
// 		RecordType: dao.StatusSK,
// 		Status:     delivery.ORDER_STATUS_name[int32(delivery.ORDER_STATUS_DELIVERED)],
// 		Address:    in.Address,
// 	})
// }
