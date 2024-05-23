package order

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/elusiv0/wb_tech_l0/internal/dto/order"
	"github.com/elusiv0/wb_tech_l0/internal/repo"
	"github.com/elusiv0/wb_tech_l0/internal/repo/converter"
	"github.com/elusiv0/wb_tech_l0/internal/repo/model"
	"github.com/elusiv0/wb_tech_l0/pkg/postgres"
	"github.com/jackc/pgx/v4"
)

type OrderRepository struct {
	db     *postgres.Postgres
	logger *slog.Logger
}

func New(
	postgres *postgres.Postgres,
	logger *slog.Logger,
) *OrderRepository {
	repo := &OrderRepository{
		db:     postgres,
		logger: logger,
	}

	return repo
}

var _ repo.OrderRepo = &OrderRepository{}

const (
	orderTable    = "order_"
	paymentTable  = "payment"
	deliveryTable = "delivery"
	itemTable     = "item"
)

func (repo *OrderRepository) Insert(ctx context.Context, orderDto order.Order) (orderResp order.Order, err error) {
	tx, err := repo.db.PgxPool.Begin(ctx)
	var deliveryId int
	var paymentId int
	orderReq := converter.OrderToModelFromDto(orderDto)

	if err != nil {
		return orderResp, fmt.Errorf("Begin Transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
			err = fmt.Errorf("OrderRepository - Insert: %w", err)
			return
		}
		err = tx.Commit(ctx)
	}()

	delivery := orderReq.Delivery
	sql, args, err := repo.db.Builder.
		Insert(deliveryTable).
		Columns("delivery_name", "phone", "zip", "city", "delivery_address", "region", "email").
		Values(
			delivery.DeliveryName, delivery.Phone, delivery.Zip, delivery.City,
			delivery.DeliveryAddress, delivery.Region, delivery.Email,
		).
		Suffix("RETURNING delivery_id").
		ToSql()
	if err != nil {
		return orderResp, fmt.Errorf("Delivery building sql: %w", err)
	}
	row := tx.QueryRow(ctx, sql, args...)
	if err := row.Scan(&deliveryId); err != nil {
		return orderResp, fmt.Errorf("Delivery exec sql: %w", err)
	}
	payment := orderReq.Payment
	sql, args, err = repo.db.Builder.
		Insert(paymentTable).
		Columns(
			"payment_transaction", "request_id", "currency", "payment_provider",
			"amount", "payment_dt", "bank", "delivery_cost",
			"goods_total", "custom_fee",
		).
		Values(
			payment.PaymentTransaction, payment.RequiestId, payment.Currency, payment.PaymentProvider,
			payment.Amount, payment.PaymentDt, payment.Bank, payment.DeliveryCost,
			payment.GoodsTotal, payment.CustomFee,
		).
		Suffix("RETURNING payment_id").
		ToSql()
	if err != nil {
		return orderResp, fmt.Errorf("Payment building sql: %w", err)
	}
	row = tx.QueryRow(ctx, sql, args...)
	if err := row.Scan(&paymentId); err != nil {
		return orderResp, fmt.Errorf("Payment exec sql: %w", err)
	}

	sql, args, err = repo.db.Builder.
		Insert(orderTable).
		Columns(
			"order_uid", "track_number", "order_entry", "delivery_id",
			"payment_id", "locale", "internal_signature", "customer_id",
			"delivery_service", "shard_key", "sm_id", "date_created",
			"oof_shard",
		).
		Values(
			orderReq.OrderUuid, orderReq.TrackNumber, orderReq.OrderEntry, deliveryId,
			paymentId, orderReq.Locale, orderReq.InternalSignature, orderReq.CustomerId,
			orderReq.DeliveryService, orderReq.ShardKey, orderReq.SmId, orderReq.DateCreated,
			orderReq.OofShard,
		).
		Suffix("RETURNING order_uid").
		ToSql()
	if err != nil {
		return orderResp, fmt.Errorf("Order building sql: %w", err)
	}
	row = tx.QueryRow(ctx, sql, args...)
	if err := row.Scan(&orderResp.OrderUid); err != nil {
		return orderResp, fmt.Errorf("Order exec sql: %w", err)
	}

	items := orderReq.Items

	itemsRows := make([][]interface{}, len(items), len(items))
	for i := 0; i < len(items); i++ {
		itemsRows[i] = []interface{}{
			items[i].ChrtId,
			items[i].TrackNumber,
			items[i].Price,
			items[i].Rid,
			items[i].ItemName,
			items[i].Sale,
			items[i].Size,
			items[i].TotalPrice,
			items[i].NmId,
			items[i].Brand,
			items[i].ItemStatus,
		}
	}
	_, err = tx.CopyFrom(ctx, pgx.Identifier{itemTable}, []string{
		"chrt_id", "track_number", "price", "rid",
		"item_name", "sale", "size", "total_price",
		"nm_id", "brand", "item_status",
	}, pgx.CopyFromRows(itemsRows))
	if err != nil {
		return orderResp, fmt.Errorf("Insert items copy from function: %w", err)
	}

	return orderDto, nil
}

func (repo *OrderRepository) GetAll(ctx context.Context) ([]order.Order, error) {
	var orderResp []order.Order
	repo.logger.Info("building get all orders sql stmt")
	sql, _, err := repo.db.Builder.
		Select(
			"o.order_uid", "o.track_number", "o.order_entry", "d.delivery_name",
			"d.phone", "d.zip", "d.city", "d.delivery_address",
			"d.region", "d.email", "p.payment_transaction", "p.request_id",
			"p.currency", "p.payment_provider", "p.amount", "p.payment_dt",
			"p.bank", "p.delivery_cost", "p.goods_total", "p.custom_fee",
			"o.locale", "o.internal_signature", "o.customer_id", "o.delivery_service",
			"o.shard_key", "o.sm_id", "o.date_created", "o.oof_shard",
			"i.item_id", "i.chrt_id", "i.track_number", "i.price",
			"i.rid", "i.item_name", "i.sale", "i.size",
			"i.total_price", "i.nm_id", "i.brand", "i.item_status",
		).
		From(orderTable + " o").
		LeftJoin(deliveryTable + " d ON o.delivery_id = d.delivery_id").
		LeftJoin(paymentTable + " p ON o.payment_id = p.payment_id").
		LeftJoin(itemTable + " i ON o.track_number = i.track_number").
		ToSql()
	if err != nil {
		return orderResp, fmt.Errorf("Order sql builder: %w", err)
	}

	repo.logger.Info("execute get all orders sql stmt")
	rows, err := repo.db.PgxPool.Query(ctx, sql)
	if err != nil {
		return orderResp, fmt.Errorf("Order exec sql: %w", err)
	}
	defer rows.Close()
	orders := make(map[string]*model.Order)
	repo.logger.Info("extract orders from result rows")
	for rows.Next() {
		curOrder := &model.Order{}
		item := model.Item{}

		err = rows.Scan(
			&curOrder.OrderUuid, &curOrder.TrackNumber, &curOrder.OrderEntry,
			&curOrder.Delivery.DeliveryName, &curOrder.Delivery.Phone, &curOrder.Delivery.Zip,
			&curOrder.Delivery.City, &curOrder.Delivery.DeliveryAddress, &curOrder.Delivery.Region,
			&curOrder.Delivery.Email, &curOrder.Payment.PaymentTransaction, &curOrder.Payment.RequiestId,
			&curOrder.Payment.Currency, &curOrder.Payment.PaymentProvider, &curOrder.Payment.Amount,
			&curOrder.Payment.PaymentDt, &curOrder.Payment.Bank, &curOrder.Payment.DeliveryCost,
			&curOrder.Payment.GoodsTotal, &curOrder.Payment.CustomFee, &curOrder.Locale,
			&curOrder.InternalSignature, &curOrder.CustomerId, &curOrder.DeliveryService,
			&curOrder.ShardKey, &curOrder.SmId, &curOrder.DateCreated,
			&curOrder.OofShard, &item.ItemId, &item.ChrtId,
			&item.TrackNumber, &item.Price, &item.Rid,
			&item.ItemName, &item.Sale, &item.Size,
			&item.TotalPrice, &item.NmId, &item.Brand,
			&item.ItemStatus,
		)
		if err != nil {
			return orderResp, fmt.Errorf("Order scan query: %w", err)
		}
		if _, ok := orders[curOrder.OrderUuid]; !ok {
			orders[curOrder.OrderUuid] = curOrder
		}
		curOrder = orders[curOrder.OrderUuid]
		curOrder.Items = append(curOrder.Items, item)
	}
	repo.logger.Info("convert model orders to dto")
	for _, val := range orders {
		orderDto := converter.OrderToDtoFromModel(*val)
		orderResp = append(orderResp, orderDto)
	}

	return orderResp, nil
}
