package function

import (
	"fmt"
	"github.com/oneconcern/kubeless/pkg/utils"
	"os"
)

func ListAdapter(out, namespace string, args []string) error {
	if namespace == "" {
		namespace = utils.GetDefaultNamespace()
	}

	kubelessClient, err := utils.GetKubelessClientOutCluster()
	if err != nil {
		return fmt.Errorf(" can not list models: %v", err)
	}

	apiV1Client := utils.GetClientOutOfCluster()

	if err := doList(os.Stdout, kubelessClient, apiV1Client, namespace, out, args); err != nil {
		return err
	}

	return nil
}
