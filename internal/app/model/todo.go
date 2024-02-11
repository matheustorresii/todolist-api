package model

type Todo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	DeviceId    string `json:"device_id"`
}
