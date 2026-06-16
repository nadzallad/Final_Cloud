package entity

import "time"

type WarehouseLog struct {
	WarehouseID    int       `gorm:"column:warehouse_id;primaryKey;autoIncrement" json:"warehouse_id"`
	UserID         int       `gorm:"column:user_id" json:"user_id"`
	TrackingNumber string    `gorm:"column:tracking_number" json:"tracking_number"`
	ItemName       string    `gorm:"column:item_name" json:"item_name"`
	Stock          int       `gorm:"column:stock" json:"stock"`
	Status         string    `gorm:"column:status;default:IN_WAREHOUSE" json:"status"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (WarehouseLog) TableName() string {
	return "warehouse_logs"
}
