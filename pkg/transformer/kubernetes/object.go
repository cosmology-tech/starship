package kubernetes

import (
	"errors"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
	"path/filepath"
)

// RuntimeObjects holds the return type for the Transformer, which implements
// transformer.Objects interface
type RuntimeObjects struct {
	Objects []runtime.Object
}

// WriteToFile implements the transformer.Object interface
func (ro *RuntimeObjects) WriteToFile(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	// create files
	for _, obj := range ro.Objects {
		data, err := marshal(obj, false, 2)
		if err != nil {
			return err
		}

		mo, ok := obj.(metav1.Object)
		if !ok {
			fmt.Println("unstructured object", mo)
			return errors.New("unable to change interface")
		}
		file := filepath.Join(dir, mo.GetName()) + ".yaml"
		if err := os.WriteFile(file, data, 0644); err != nil {
			return errors.New("Failed to write: " + err.Error())
		}
	}

	return nil
}

func (ro *RuntimeObjects) Validate() error {
	return nil
}
