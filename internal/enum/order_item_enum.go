package enum

type OrderItemEnum string

const (
	Pending    OrderItemEnum = "Pending"
	OnProgress OrderItemEnum = "OnProgress"
	Done       OrderItemEnum = "Done"
	Cancel     OrderItemEnum = "Cancel"
)
