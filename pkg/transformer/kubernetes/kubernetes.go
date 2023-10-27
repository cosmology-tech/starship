package kubernetes

import (
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
		base[key] = v
	}
	return base
}

func (k *Kubernetes) nodeConfigToStatefulSets(nodeConfig types.NodeConfig) (*appsv1.StatefulSet, error) {
	podSpec, err := k.InitPodSpecWithNodeConfig(nodeConfig)
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
			Name:   nodeConfig.ContainerName,
			Labels: k.configAllLabels(nodeConfig.ContainerName, nodeConfig.Labels),
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: &rp,
			Template: api.PodTemplateSpec{
				Spec: podSpec,
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: ConfigLabels(nodeConfig.ContainerName),
			},
			ServiceName: nodeConfig.ContainerName,
		},
	}
	return ds, nil
}

func (k *Kubernetes) InitPodSpecWithNodeConfig(nodeConfig types.NodeConfig) (api.PodSpec, error) {
	initContainers, err := InitToContainer(nodeConfig)
	if err != nil {
		return api.PodSpec{}, err
	}
	pod := api.PodSpec{
		RestartPolicy:  api.RestartPolicyOnFailure,
		InitContainers: initContainers,
		Containers: []api.Container{
			{
				Name:            nodeConfig.ContainerName,
				Image:           nodeConfig.Image,
				ImagePullPolicy: api.PullPolicy(nodeConfig.ImagePullPolicy),
			},
		},
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
