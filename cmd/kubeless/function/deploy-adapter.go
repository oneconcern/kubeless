package function

import (
	"fmt"
	"log"

	kubelessApi "github.com/kubeless/kubeless/pkg/apis/kubeless/v1beta1"
	kubelessUtils "github.com/kubeless/kubeless/pkg/utils"
)

func DeployModel(modelName, ns, handler, file, deps, runtime, runtimeImage, mem, cpu, timeout string, imagePullPolicy string, port int32, headless bool, envs, labels []string, secrets []string) error {
	defaultFunctionSpec := kubelessApi.Function{}
	defaultFunctionSpec.ObjectMeta.Labels = map[string]string{
		"created-by": "kubeless",
		"function":   modelName,
	}

	f, err := getFunctionDescription(modelName, ns, handler, file, deps, runtime, runtimeImage, mem, cpu, timeout, imagePullPolicy, port, headless, envs, labels, secrets, defaultFunctionSpec)
	if err != nil {
		return fmt.Errorf("model description %v ", err)
	}

	kubelessClient, err := kubelessUtils.GetKubelessClientOutCluster()
	if err != nil {
		return fmt.Errorf("get kubeless client out cluster %v ", err)
	}

	log.Println("deploying model...")
	err = kubelessUtils.CreateFunctionCustomResource(kubelessClient, f)
	if err != nil {
		return fmt.Errorf("failed to deploy %s. Received:\n%s", modelName, err)
	}
	log.Printf("model %s submitted for deployment", modelName)
	log.Printf("check the deployment status executing 'datamon model ls %s'", modelName)

	return nil
}
