package types

type StarshipObject struct {
	ServiceConfigs map[string]NodeConfig
}

type EnvVar struct {
	Name  string
	Value string
}

type Ports struct {
	Name          string
	HostPort      int32
	ContainerPort int32
	HostIP        string
	Protocol      string
}

// Init describes precontainer creation steps. These can be either as docker containers
// or just scripts. These are supposed to run to completion
type Init struct {
	Name        string
	Image       string // optional
	Command     []string
	WorkingDir  string
	ScriptFiles []string
}

type Resource struct {
	CPU    string `name:"cpu" json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory string `name:"memory" json:"memory,omitempty" yaml:"memory,omitempty"`
}

type File struct {
	Name      string
	Path      string
	Content   []byte
	MountPath string
}

// NodeConfig describes containers and objects without specifics of underlying infra
type NodeConfig struct {
	Name            string
	ContainerName   string
	Controller      string
	Image           string
	Environment     []EnvVar
	EnvFile         []string
	Port            []Ports
	Command         []string
	ScriptFiles     []string
	WorkingDir      string
	Init            []Init
	DependsOn       []*NodeConfig // Pointer to serviceconfig this would depend on
	Replicas        int
	Labels          map[string]string
	Annotations     map[string]string
	Sidecars        []*NodeConfig // pointers to serviceconfig that are run as sidecars
	Resources       Resource
	ImagePullPolicy string
	Files           []File
}
