/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package prebind

import (
	"fmt"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

// StatelessPrebindExample is an example of a simple plugin that has no state
// and implements only one hook for prebind.
type StatelessPrebindExample struct{}

var _ = framework.PrebindPlugin(StatelessPrebindExample{})

// Name is the name of the plugin used in Registry and configurations.
const Name = "stateless-prebind-plugin-example"

// Name returns name of the plugin. It is used in logs, etc.
func (sr StatelessPrebindExample) Name() string {
	return Name
}

// Prebind is the functions invoked by the framework at "prebind" extension point.
func (sr StatelessPrebindExample) Prebind(pc *framework.PluginContext, pod *v1.Pod, nodeName string) *framework.Status {
	if pod == nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("pod cannot be nil"))
	}
	if pod.Namespace != "foo" {
		return framework.NewStatus(framework.Unschedulable, "only pods from 'foo' namespace are allowed")
	}
	return nil
}

// New initializes a new plugin and returns it.
func New(_ *runtime.Unknown, _ framework.Framework) (framework.Plugin, error) {
	return &StatelessPrebindExample{}, nil
}
