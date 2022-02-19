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

package plugins

import (
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/defaultbinder"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/defaultpreemption"
	plfeature "github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/feature"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/imagelocality"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/interpodaffinity"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/nodeaffinity"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/nodename"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/nodeports"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/noderesources"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/nodeunschedulable"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/nodevolumelimits"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/podtopologyspread"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/queuesort"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/selectorspread"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/tainttoleration"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/volumebinding"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/volumerestrictions"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/plugins/volumezone"
	"github.com/QuarfotPrice/sched.dev/pkg/scheduler/framework/runtime"
	"k8s.io/apiserver/pkg/util/feature"
	"k8s.io/kubernetes/pkg/features"
)

// NewInTreeRegistry builds the registry with all the in-tree plugins.
// A scheduler that runs out of tree plugins can register additional plugins
// through the WithFrameworkOutOfTreeRegistry option.
func NewInTreeRegistry() runtime.Registry {
	fts := plfeature.Features{
		EnablePodAffinityNamespaceSelector: feature.DefaultFeatureGate.Enabled(features.PodAffinityNamespaceSelector),
		EnablePodDisruptionBudget:          feature.DefaultFeatureGate.Enabled(features.PodDisruptionBudget),
		EnablePodOverhead:                  feature.DefaultFeatureGate.Enabled(features.PodOverhead),
		EnableReadWriteOncePod:             feature.DefaultFeatureGate.Enabled(features.ReadWriteOncePod),
		EnableVolumeCapacityPriority:       feature.DefaultFeatureGate.Enabled(features.VolumeCapacityPriority),
		EnableCSIStorageCapacity:           feature.DefaultFeatureGate.Enabled(features.CSIStorageCapacity),
	}

	return runtime.Registry{
		selectorspread.Name:                  selectorspread.New,
		imagelocality.Name:                   imagelocality.New,
		tainttoleration.Name:                 tainttoleration.New,
		nodename.Name:                        nodename.New,
		nodeports.Name:                       nodeports.New,
		nodeaffinity.Name:                    nodeaffinity.New,
		podtopologyspread.Name:               podtopologyspread.New,
		nodeunschedulable.Name:               nodeunschedulable.New,
		noderesources.FitName:                runtime.FactoryAdapter(fts, noderesources.NewFit),
		noderesources.BalancedAllocationName: runtime.FactoryAdapter(fts, noderesources.NewBalancedAllocation),
		volumebinding.Name:                   runtime.FactoryAdapter(fts, volumebinding.New),
		volumerestrictions.Name:              runtime.FactoryAdapter(fts, volumerestrictions.New),
		volumezone.Name:                      volumezone.New,
		nodevolumelimits.CSIName:             runtime.FactoryAdapter(fts, nodevolumelimits.NewCSI),
		nodevolumelimits.EBSName:             runtime.FactoryAdapter(fts, nodevolumelimits.NewEBS),
		nodevolumelimits.GCEPDName:           runtime.FactoryAdapter(fts, nodevolumelimits.NewGCEPD),
		nodevolumelimits.AzureDiskName:       runtime.FactoryAdapter(fts, nodevolumelimits.NewAzureDisk),
		nodevolumelimits.CinderName:          runtime.FactoryAdapter(fts, nodevolumelimits.NewCinder),
		interpodaffinity.Name:                runtime.FactoryAdapter(fts, interpodaffinity.New),
		queuesort.Name:                       queuesort.New,
		defaultbinder.Name:                   defaultbinder.New,
		defaultpreemption.Name:               runtime.FactoryAdapter(fts, defaultpreemption.New),
	}
}
