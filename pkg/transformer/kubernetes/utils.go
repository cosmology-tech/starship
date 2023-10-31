package kubernetes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cosmology-tech/starship/pkg/types"
	"gopkg.in/yaml.v3"
	api "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

const Selector = "apps.kubernetes.io/name"

// ConfigLabels configures label name alone
func ConfigLabels(name string) map[string]string {
	return map[string]string{Selector: name}
}

func ResourceRequirementsFromConfig(r types.Resource) (api.ResourceRequirements, error) {
	rr := api.ResourceRequirements{
		Requests: api.ResourceList{},
		Limits:   api.ResourceList{},
	}

	if r.CPU != "" {
		cpu, err := resource.ParseQuantity(r.CPU)
		if err != nil {
			return rr, err
		}
		rr.Requests[api.ResourceCPU] = cpu
		rr.Limits[api.ResourceCPU] = cpu
	}

	if r.Memory != "" {
		mem, err := resource.ParseQuantity(r.Memory)
		if err != nil {
			return rr, err
		}
		rr.Requests[api.ResourceMemory] = mem
		rr.Limits[api.ResourceMemory] = mem
	}

	return rr, nil
}

func InitWaitContainer(nodeConfig types.NodeConfig) (api.Container, error) {
	cmdfmt := `
while [ $(curl -sw '%s' http://%s.$NAMESPACE.svc.cluster.local:%s/%s -o /dev/null) -ne 200 ]; do
	echo "Genesis validator does not seem to be ready for: %s. Waiting for it to start..."
	sleep 10;
done
`
	cmd := ""
	for _, deps := range nodeConfig.DependsOn {
		// note: add exposer port for waiting for chain
		// todo: make this dependent on NodeConfig readiness information
		cmd = cmd + "\n" + fmt.Sprintf(cmdfmt, "%{http_code}", deps.ContainerName, "8081", "node_id", deps.ContainerName)
	}

	// set default resources to 100m and 100Mi for container
	resources, err := ResourceRequirementsFromConfig(types.Resource{CPU: "100m", Memory: "100Mi"})
	if err != nil {
		return api.Container{}, err
	}

	wait := api.Container{
		Name:            "wait",
		Image:           "curlimages/curl",
		ImagePullPolicy: api.PullPolicy(nodeConfig.ImagePullPolicy),
		Command: []string{
			"bash",
			"-c",
			cmd,
		},
		Env: []api.EnvVar{
			{
				Name: "NAMESPACE",
				ValueFrom: &api.EnvVarSource{
					FieldRef: &api.ObjectFieldSelector{
						FieldPath: "metadata.namespace",
					},
				},
			},
		},
		Resources: resources,
	}

	return wait, nil
}

func EnvVarFromNodeConfig(nodeConfig types.NodeConfig) []api.EnvVar {
	envs := []api.EnvVar{}
	for _, nenv := range nodeConfig.Environment {
		envs = append(envs, api.EnvVar{
			Name:  nenv.Name,
			Value: nenv.Value,
		})
	}
	return envs
}

func VolumeMountsFromNodeConfig(nodeConfig types.NodeConfig) []api.VolumeMount {
	volumeMounts := []api.VolumeMount{}
	for _, mount := range nodeConfig.Mounts {
		vm := api.VolumeMount{
			Name:      mount.Name,
			MountPath: mount.MountPath,
		}
		volumeMounts = append(volumeMounts, vm)
	}

	return volumeMounts
}

func InitContainers(nodeConfig types.NodeConfig) ([]api.Container, error) {
	initContainers := []api.Container{}

	res, err := ResourceRequirementsFromConfig(nodeConfig.Resources)
	if err != nil {
		return nil, err
	}

	volumeMounts := VolumeMountsFromNodeConfig(nodeConfig)

	// add init-wait containers based on nodeConfig.DependsOn
	waitInitContainer, err := InitWaitContainer(nodeConfig)
	if err != nil {
		return nil, err
	}
	initContainers = append(initContainers, waitInitContainer)

	for _, init := range nodeConfig.Init {
		envs := []api.EnvVar{}
		for _, nenv := range nodeConfig.Environment {
			envs = append(envs, api.EnvVar{
				Name:  nenv.Name,
				Value: nenv.Value,
			})
		}
		initContainer := api.Container{
			Name:            init.Name,
			Image:           init.Image,
			ImagePullPolicy: api.PullPolicy(nodeConfig.ImagePullPolicy),
			Env:             EnvVarFromNodeConfig(nodeConfig),
			Command:         init.Command,
			Resources:       res,
			VolumeMounts:    volumeMounts,
		}
		initContainers = append(initContainers, initContainer)
	}
	return initContainers, nil
}

// marshal object runtime.Object and return byte array
func marshal(obj runtime.Object, jsonFormat bool, indent int) (data []byte, err error) {
	// convert data to yaml or json
	if jsonFormat {
		data, err = json.MarshalIndent(obj, "", "  ")
	} else {
		data, err = marshalWithIndent(obj, indent)
	}
	if err != nil {
		data = nil
	}
	return
}

// Convert JSON to YAML.
func jsonToYaml(j []byte, spaces int) ([]byte, error) {
	// Convert the JSON to an object.
	var jsonObj interface{}
	// We are using yaml.Unmarshal here (instead of json.Unmarshal) because the
	// Go JSON library doesn't try to pick the right number type (int, float,
	// etc.) when unmarshling to interface{}, it just picks float64
	// universally. go-yaml does go through the effort of picking the right
	// number type, so we can preserve number type throughout this process.
	err := yaml.Unmarshal(j, &jsonObj)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer

	encoder := yaml.NewEncoder(&b)
	encoder.SetIndent(spaces)
	if err := encoder.Encode(jsonObj); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func marshalWithIndent(o interface{}, indent int) ([]byte, error) {
	j, err := json.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("error marshaling into JSON: %s", err.Error())
	}

	y, err := jsonToYaml(j, indent)
	if err != nil {
		return nil, fmt.Errorf("error converting JSON to YAML: %s", err.Error())
	}

	return y, nil
}
