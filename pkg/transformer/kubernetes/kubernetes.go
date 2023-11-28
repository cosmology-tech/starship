package kubernetes

import (
	"fmt"
	"github.com/cosmology-tech/starship/pkg/types"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Kubernetes implements Transformer interface and represents Kubernetes transformer
type Kubernetes struct {
	logger *zap.Logger
	// the user provided options from the command line
	Opt types.ConvertOptions
}

func NewKubernetes(logger *zap.Logger) Kubernetes {
	return Kubernetes{logger: logger}
}

const (
	// DeploymentController is controller type for Deployment
	DeploymentController = "deployment"
	// DaemonSetController is controller type for DaemonSet
	DaemonSetController = "daemonset"
	// StatefulStateController is controller type for StatefulSet
	StatefulStateController = "statefulset"
)

func (k *Kubernetes) configAllLabels(name string, labels map[string]string) map[string]string {
	base := ConfigLabels(name)
	for key, v := range labels {
		newKey := fmt.Sprintf("apps.kubernetes.io/%s", key)
		base[newKey] = v
	}
	return base
}

func (k *Kubernetes) nodeConfigToStatefulSets(nodeConfig types.NodeConfig) (*appsv1.StatefulSet, error) {
	podSpec, err := k.PodSpecWithNodeConfig(nodeConfig)
	if err != nil {
		return nil, err
	}
	rp := int32(nodeConfig.Replicas)
	ds := &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: nodeConfig.ContainerName,
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: &rp,
			Template: api.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels:      k.configAllLabels(nodeConfig.ContainerName, nodeConfig.Labels),
					Annotations: nodeConfig.Annotations,
				},
				Spec: podSpec,
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: k.configAllLabels(nodeConfig.ContainerName, nodeConfig.Labels),
			},
			ServiceName: nodeConfig.ContainerName,
		},
	}
	return ds, nil
}

func (k *Kubernetes) ContainerFromNodeConfig(nodeConfig types.NodeConfig) (api.Container, error) {
	resources, err := ResourceRequirementsFromConfig(nodeConfig.Resources)
	if err != nil {
		return api.Container{}, err
	}

	vms := VolumeMountsFromNodeConfig(nodeConfig)

	return api.Container{
		Name:            nodeConfig.ContainerName,
		Image:           nodeConfig.Image,
		Command:         nodeConfig.Command,
		WorkingDir:      nodeConfig.WorkingDir,
		Env:             EnvVarFromNodeConfig(nodeConfig),
		Resources:       resources,
		VolumeMounts:    vms,
		ImagePullPolicy: api.PullPolicy(nodeConfig.ImagePullPolicy),
	}, nil
}

func (k *Kubernetes) PodSpecWithNodeConfig(nodeConfig types.NodeConfig) (api.PodSpec, error) {
	initContainers, err := InitContainers(nodeConfig)
	if err != nil {
		return api.PodSpec{}, err
	}

	container, err := k.ContainerFromNodeConfig(nodeConfig)
	pod := api.PodSpec{
		RestartPolicy:  api.RestartPolicyOnFailure,
		InitContainers: initContainers,
		Containers:     []api.Container{container},
	}

	// return early incase of no sidecars
	if nodeConfig.Sidecars == nil {
		return pod, nil
	}

	for _, sidecar := range nodeConfig.Sidecars {
		sidecarContainer, err := k.ContainerFromNodeConfig(*sidecar)
		if err != nil {
			return api.PodSpec{}, err
		}
		pod.Containers = append(pod.Containers, sidecarContainer)
	}

	return pod, nil
}

func (k *Kubernetes) Transform(nodeConfigs []types.NodeConfig, opts types.ConvertOptions) (RuntimeObjects, error) {
	var allobjects []runtime.Object

	for _, nodeConfig := range nodeConfigs {
		if nodeConfig.Controller == StatefulStateController {
			nss, err := k.nodeConfigToStatefulSets(nodeConfig)
			if err != nil {
				return RuntimeObjects{}, err
			}
			k.logger.Info("converted nodeconfig to statefulset", zap.Any("statefulset", nss))
			allobjects = append(allobjects, nss)
		}
	}

	return RuntimeObjects{allobjects}, nil
}
