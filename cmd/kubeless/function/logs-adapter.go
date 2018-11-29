package function

import (
	"fmt"
	"github.com/oneconcern/kubeless/pkg/utils"
	"io"
	"k8s.io/api/core/v1"
	"os"
)

func LogAdapter(model, ns string, follow bool) error {
	if ns == "" {
		ns = utils.GetDefaultNamespace()
	}

	k8sClient := utils.GetClientOutOfCluster()

	pods, err := utils.GetPodsByLabel(k8sClient, ns, "function", model)
	if err != nil {
		return fmt.Errorf(" can't find the model pod: %v", err)
	}
	readyPod, err := utils.GetReadyPod(pods)
	if err != nil {
		return fmt.Errorf("no model pod is running: %v", err)
	}
	podLog := &v1.PodLogOptions{
		Container: model,
		Follow:    follow,
	}
	req := k8sClient.Core().Pods(ns).GetLogs(readyPod.Name, podLog)

	readCloser, err := req.Stream()
	if err != nil {
		return fmt.Errorf("getting log failed: %v", err)
	}
	defer readCloser.Close()
	io.Copy(os.Stdout, readCloser)

	return nil
}
