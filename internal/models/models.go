package models

type CalculateRequest struct {
	Operation string  `json:"operation" binding:"required,oneof=add subtract multiply divide"`
	A         float64 `json:"a" binding:"required"`
	B         float64 `json:"b" binding:"required"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}
