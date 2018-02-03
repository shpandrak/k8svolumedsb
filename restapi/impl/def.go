package impl

import (
	"ocopea/k8svolumedsb/models"
	"net/http"
	"github.com/go-openapi/runtime"
)

type ErrorResponse interface {
	SetPayload(payload *models.Error)
	WriteResponse(http.ResponseWriter, runtime.Producer)
}

