package controllers

import (
	"medpoinraiden/internal/models"

	"github.com/sev-2/raiden"
)

type DokterResponse struct {
    Message string `json:"message"`
}

type DokterController struct {
	raiden.ControllerBase
	Http string `path:"/dokter" type:"rest"`
	Model models.Dokter
	Result DokterResponse
}



