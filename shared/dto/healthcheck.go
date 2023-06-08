package dto

const (
	OK    = "OK"
	Error = "Error"

	HTTP = "Http"
	DB   = "Database"
)

type (
	// Status Status
	Status struct {
		Name   string      `json:"name"`
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	}

	HCStatus struct {
		Status []Status `json:"status"`
	}

	HCData struct {
		HandlerCount uint32 `json:"handlerCount"`
	}
)