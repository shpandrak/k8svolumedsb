package impl

import (
	"ocopea/k8svolumedsb/models"
	"ocopea/k8svolumedsb/restapi/operations/dsb_web"
	"github.com/go-openapi/runtime/middleware"
	k8sClient "ocopea/kubernetes/client"
	"log"
)

type K8sDsbBindings struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     string `json:"port"`
}

func DsbGetServiceInstancesResponse(
k8s *k8sClient.Client,
params dsb_web.GetServiceInstanceParams) middleware.Responder {

	d, err := getBindingInfoForInstance(k8s, params.InstanceID)
	if (err != nil) {
		return getError(dsb_web.NewGetServiceInstanceDefault(500), err, 500)
	}

	return dsb_web.NewGetServiceInstanceOK().WithPayload(d)

}

func getBindingInfoForInstance(k8s *k8sClient.Client, instanceId string) (*models.ServiceInstanceDetails, error) {
	volumeName := getVolumeNameFromServiceInstanceId(instanceId)
	log.Printf("testing volume %s ", volumeName)
	isReady, _, err := k8s.TestVolume(volumeName)
	if (err != nil) {
		return nil, err
	}

	// Still creating
	if (!isReady) {
		return &models.ServiceInstanceDetails{
			InstanceID: instanceId,
			State: "CREATING",
		}, nil
	}

	bindingInfo := make(map[string]string)
	bindingInfo["volumeName"] = volumeName
	return &models.ServiceInstanceDetails{
		InstanceID: instanceId,
		State: "RUNNING",
		Binding: bindingInfo,
		Size:10,
		StorageType: "Kubernetes Persistent Volume",
	}, nil
}