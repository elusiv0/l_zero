package converter

import (
	deliveryDto "github.com/elusiv0/wb_tech_l0/internal/dto/delivery"
	itemDto "github.com/elusiv0/wb_tech_l0/internal/dto/item"
	orderDto "github.com/elusiv0/wb_tech_l0/internal/dto/order"
	paymentDto "github.com/elusiv0/wb_tech_l0/internal/dto/payment"
	"github.com/elusiv0/wb_tech_l0/internal/repo/model"
)

func OrderToDtoFromModel(orderModel model.Order) orderDto.Order {
	var itemsDto []itemDto.Item
	for _, val := range orderModel.Items {
		itemDto := ItemToDtoFromModel(val)
		itemsDto = append(itemsDto, itemDto)
	}
	deliveryDto := DeliveryToDtoFromModel(orderModel.Delivery)
	paymentDto := PaymentToDtoFromModel(orderModel.Payment)
	orderDto := orderDto.Order{
		OrderUid:          orderModel.OrderUuid,
		TrackNumber:       orderModel.TrackNumber,
		Entry:             orderModel.OrderEntry,
		Delivery:          deliveryDto,
		Payment:           paymentDto,
		Items:             itemsDto,
		Locale:            orderModel.Locale,
		InternalSignature: orderModel.InternalSignature,
		CustomerId:        orderModel.CustomerId,
		DeliveryService:   orderModel.DeliveryService,
		ShardKey:          orderModel.ShardKey,
		SmId:              orderModel.SmId,
		DateCreated:       orderModel.DateCreated,
		OofShard:          orderModel.OofShard,
	}

	return orderDto
}

func OrderToModelFromDto(orderDto orderDto.Order) model.Order {
	var itemsModel []model.Item
	for _, val := range orderDto.Items {
		itemModel := ItemToModelFromDto(val)
		itemsModel = append(itemsModel, itemModel)
	}
	deliveryModel := DeliveryToModelFromDto(orderDto.Delivery)
	paymentModel := PaymentToModelFromDto(orderDto.Payment)
	orderModel := model.Order{
		OrderUuid:         orderDto.OrderUid,
		TrackNumber:       orderDto.TrackNumber,
		OrderEntry:        orderDto.Entry,
		Delivery:          deliveryModel,
		Payment:           paymentModel,
		Items:             itemsModel,
		Locale:            orderDto.Locale,
		InternalSignature: orderDto.InternalSignature,
		CustomerId:        orderDto.CustomerId,
		DeliveryService:   orderDto.DeliveryService,
		ShardKey:          orderDto.ShardKey,
		SmId:              orderDto.SmId,
		DateCreated:       orderDto.DateCreated,
		OofShard:          orderDto.OofShard,
	}

	return orderModel
}

func PaymentToDtoFromModel(paymentModel model.Payment) paymentDto.Payment {
	paymentDto := paymentDto.Payment{
		Transaction:  paymentModel.PaymentTransaction,
		RequestId:    paymentModel.RequiestId,
		Currency:     paymentModel.Currency,
		Provider:     paymentModel.PaymentProvider,
		Amount:       paymentModel.Amount,
		PaymentDt:    paymentModel.PaymentDt,
		Bank:         paymentModel.Bank,
		DeliveryCost: paymentModel.DeliveryCost,
		GoodsTotal:   paymentModel.GoodsTotal,
		CustomFee:    paymentModel.CustomFee,
	}

	return paymentDto
}

func PaymentToModelFromDto(paymentDto paymentDto.Payment) model.Payment {
	paymentModel := model.Payment{
		PaymentTransaction: paymentDto.Transaction,
		RequiestId:         paymentDto.RequestId,
		Currency:           paymentDto.Currency,
		PaymentProvider:    paymentDto.Provider,
		Amount:             paymentDto.Amount,
		PaymentDt:          paymentDto.PaymentDt,
		Bank:               paymentDto.Bank,
		DeliveryCost:       paymentDto.DeliveryCost,
		GoodsTotal:         paymentDto.GoodsTotal,
		CustomFee:          paymentDto.CustomFee,
	}

	return paymentModel
}

func DeliveryToDtoFromModel(deliveryModel model.Delivery) deliveryDto.Delivery {
	deliveryDto := deliveryDto.Delivery{
		Name:    deliveryModel.DeliveryName,
		Phone:   deliveryModel.Phone,
		Zip:     deliveryModel.Zip,
		City:    deliveryModel.City,
		Address: deliveryModel.DeliveryAddress,
		Region:  deliveryModel.Region,
		Email:   deliveryModel.Email,
	}

	return deliveryDto
}

func DeliveryToModelFromDto(deliveryDto deliveryDto.Delivery) model.Delivery {
	deliveryModel := model.Delivery{
		DeliveryName:    deliveryDto.Name,
		Phone:           deliveryDto.Phone,
		Zip:             deliveryDto.Zip,
		City:            deliveryDto.City,
		DeliveryAddress: deliveryDto.Address,
		Region:          deliveryDto.Region,
		Email:           deliveryDto.Email,
	}

	return deliveryModel
}

func ItemToDtoFromModel(itemModel model.Item) itemDto.Item {
	itemDto := itemDto.Item{
		ChrtId:      itemModel.ChrtId,
		TrackNumber: itemModel.TrackNumber,
		Price:       itemModel.Price,
		Rid:         itemModel.Rid,
		Name:        itemModel.ItemName,
		Sale:        itemModel.Sale,
		Size:        itemModel.Size,
		TotalPrice:  itemModel.TotalPrice,
		NmId:        itemModel.NmId,
		Brand:       itemModel.Brand,
		Status:      itemModel.ItemStatus,
	}

	return itemDto
}

func ItemToModelFromDto(itemDto itemDto.Item) model.Item {
	itemModel := model.Item{
		ChrtId:      itemDto.ChrtId,
		TrackNumber: itemDto.TrackNumber,
		Price:       itemDto.Price,
		Rid:         itemDto.Rid,
		ItemName:    itemDto.Name,
		Sale:        itemDto.Sale,
		Size:        itemDto.Size,
		TotalPrice:  itemDto.TotalPrice,
		NmId:        itemDto.NmId,
		Brand:       itemDto.Brand,
		ItemStatus:  itemDto.Status,
	}

	return itemModel
}
