package impl

import (
	"github.com/go-openapi/runtime/middleware"
	"ocopea/k8svolumedsb/models"
	"ocopea/k8svolumedsb/restapi/operations/dsb_web"
	k8sClient "ocopea/kubernetes/client"
	"log"
)

func CopyServiceInstance(
k8s k8sClient.ClientInterface,
params dsb_web.CopyServiceInstanceParams) middleware.Responder {

	log.Println("faking copy yey")
	return dsb_web.NewCopyServiceInstanceOK().WithPayload(&models.CopyServiceInstanceResponse{
		CopyID:        *params.CopyDetails.CopyID,
		Status:        0,
		StatusMessage: "yey",
	})
}

