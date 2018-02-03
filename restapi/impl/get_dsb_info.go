package impl

import (
	"ocopea/k8svolumedsb/models"
	"ocopea/k8svolumedsb/restapi/operations/dsb_web"
	"github.com/go-openapi/runtime/middleware"
)

func DsbInfoResponse() middleware.Responder {

	return dsb_web.NewGetDSBInfoOK().WithPayload(
		&models.DsbInfo{
			Name: "k8s-volume-dsb",
			Description: "Docker volume and DSB for kubernetes",
			Type:  "volume",
			Plans: []*models.DsbPlan{
				{
					ID: "persistent-volume-claim",
					Name: "Persistent volume claim",
					Description: "Kubernetes persistent volume claim",
					DsbSettings: nil,
					CopyProtocols: []*models.DsbSupportedCopyProtocol{
						{
							CopyProtocol: "ShpanRest",
							CopyProtocolVersion: "1.0",
						},
					},
					Protocols: []*models.DsbSupportedProtocol{
						{
							Protocol: "docker-volume",
						},
					},
				},
			},
		})

}
