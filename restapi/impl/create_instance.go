package impl

import (
	"ocopea/k8svolumedsb/models"
	"ocopea/k8svolumedsb/restapi/operations/dsb_web"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"ocopea/kubernetes/client/v1"
	k8sClient "ocopea/kubernetes/client"
	"fmt"
	"ocopea/kubernetes/client/resource"
	"ocopea/kubernetes/client/inf"
)

func CreateInstanceResponse(
k8s *k8sClient.Client,
params *dsb_web.CreateServiceInstanceParams) middleware.Responder {

	err := createDsbInstance(k8s, params.ServiceSettings)
	if (err != nil) {
		return getError(dsb_web.NewCreateServiceInstanceDefault(500), err, 500)
	}

	return dsb_web.NewCreateServiceInstanceOK().WithPayload(
		&models.ServiceInstance{
			InstanceID: *params.ServiceSettings.InstanceID,
		})
}

func getError(failedResponse ErrorResponse, err error, errCode int) middleware.Responder {
	log.Printf("error occured %d - %s", errCode, err.Error())
	modelErrorInt := int32(errCode)
	errStr := err.Error()
	modelsError := models.Error{Code: &modelErrorInt, Message: &errStr}
	failedResponse.SetPayload(&modelsError)
	return failedResponse

}

//todo:better
func getVolumeNameFromServiceInstanceId(instanceId string) string {
	l := len(instanceId)
	if l > 125 {
		return "v-" + instanceId[l - 125:l]
	} else {
		return "v-" + instanceId
	}
}

func createDsbInstance(k8sClient *k8sClient.Client, serviceInstanceInfo *models.CreateServiceInstance) error {
	plan := serviceInstanceInfo.InstanceSettings["plan"];
	if (len(plan) == 0) {
		return fmt.Errorf("plan was not defined when trying to create service %s", *serviceInstanceInfo.InstanceID)
	}

	log.Printf("Creating DSB %s with plan %s", *serviceInstanceInfo.InstanceID, plan);
	if (plan != "persistent-volume-claim") {
		return fmt.Errorf("Unsupported plan %s when creating service %s", plan, *serviceInstanceInfo.InstanceID)

	} else {
		return createDockerPersistentVolumeInstance(k8sClient, serviceInstanceInfo)
	}
}
func createDockerPersistentVolumeInstance(k8sClient *k8sClient.Client, instance *models.CreateServiceInstance) error {
	name := getVolumeNameFromServiceInstanceId(*instance.InstanceID)

	_, err := k8sClient.CreatePersistentVolume(
		&v1.PersistentVolume{
			ObjectMeta: v1.ObjectMeta{
				Name: name,
				Labels: map[string]string{
					"type":"local",
				},
			},
			Spec: v1.PersistentVolumeSpec{
				PersistentVolumeSource: v1.PersistentVolumeSource{
					HostPath: &v1.HostPathVolumeSource{
						Path:"/tmp/data/" + *instance.InstanceID,
					},
				},
				Capacity: map[v1.ResourceName]resource.Quantity{
					"storage": {
						Amount: inf.NewDec(10, 0),
						Format: resource.BinarySI,
					},
				},
				AccessModes: []v1.PersistentVolumeAccessMode{
					v1.ReadWriteOnce,
				},
			},
		},
		true)
	return err
}



