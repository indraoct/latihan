package response

import "latihan/routing/models"

type OutputJson struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type OutputDataJson struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    []models.Bahasa `json:"data"`
}
