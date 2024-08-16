package cluster

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	opt "github.com/alibaba/kt-connect/pkg/kt/command/options"
)

// CreateRouterPod create router pod
func (k *Kubernetes) CreateRouterPod(name string, labels, annotations map[string]string, ports map[int]int) (*coreV1.Pod, error) {
	targetPorts := map[string]int{}
	for _, remotePort := range ports {
		targetPorts[fmt.Sprintf("router-%d", remotePort)] = remotePort
	}
	metaAndSpec := &PodMetaAndSpec{&ResourceMeta{
		Name:        name,
		Namespace:   opt.Get().Global.Namespace,
		Labels:      labels,
		Annotations: annotations,
	}, opt.Get().Mesh.RouterImage, map[string]string{}, targetPorts, true}
	pod := createPod(metaAndSpec)
	if _, err := k.Clientset.CoreV1().Pods(metaAndSpec.Meta.Namespace).
		Create(context.TODO(), pod, metav1.CreateOptions{}); err != nil {
		return nil, err
	}
	SetupHeartBeat(metaAndSpec.Meta.Name, metaAndSpec.Meta.Namespace, k.UpdatePodHeartBeat)
	log.Info().Msgf("Router pod %s created", name)
	return k.WaitPodReady(name, opt.Get().Global.Namespace, opt.Get().Global.PodCreationTimeout)
}

// CreateRectifierPod create pod for rectify time difference
func (k *Kubernetes) CreateRectifierPod(name string) (*coreV1.Pod, error) {
	metaAndSpec := &PodMetaAndSpec{&ResourceMeta{
		Name:        name,
		Namespace:   opt.Get().Global.Namespace,
		Labels:      map[string]string{},
		Annotations: map[string]string{},
	}, opt.Get().Global.Image, map[string]string{}, map[string]int{}, true}
	pod := createPod(metaAndSpec)
	pod.Spec.Containers[0].Command = []string{"tail", "-f", "/dev/null"}
	if _, err := k.Clientset.CoreV1().Pods(metaAndSpec.Meta.Namespace).
		Create(context.TODO(), pod, metav1.CreateOptions{}); err != nil {
		return nil, err
	}
	log.Debug().Msgf("Rectify pod %s created", name)
	return k.WaitPodReady(name, opt.Get().Global.Namespace, opt.Get().Global.PodCreationTimeout)
}
