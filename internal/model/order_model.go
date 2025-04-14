package model

// Order 定义订单模型
type Order struct {
	BaseModel
	OrderId    string `json:"order_id" gorm:"column:order_id;type:varchar(255);not null;index"`
	OrderName  string `json:"order_name" gorm:"column:order_name;type:varchar(255);not null"`
	OrderPrice string `json:"order_price" gorm:"column:order_price;type:varchar(255);not null"`
}

// TableComment 设置表 comment
func (m *Order) TableComment() string {
	return "订单表"
}
