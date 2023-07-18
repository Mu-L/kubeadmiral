//go:build exclude
/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

This file may have been modified by The KubeAdmiral Authors
("KubeAdmiral Modifications"). All KubeAdmiral Modifications
are Copyright 2023 The KubeAdmiral Authors.
*/

package util

import (
	"regexp"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restclient "k8s.io/client-go/rest"

	"github.com/kubewharf/kubeadmiral/pkg/stats"
)

// FederationNamespaces defines the namespace configuration shared by
// most federation controllers.
type FederationNamespaces struct {
	FedSystemNamespace string
	TargetNamespace    string
}

// ControllerConfig defines the configuration common to federation
// controllers.
type ControllerConfig struct {
	FederationNamespaces
	KubeConfig                            *restclient.Config
	ClusterAvailableDelay                 time.Duration
	ClusterUnavailableDelay               time.Duration
	MinimizeLatency                       bool
	SkipAdoptingResources                 bool
	WorkerCount                           int
	NamespaceAutoPropagationExcludeRegexp *regexp.Regexp
	CreateCrdForFtcs                      bool

	Metrics stats.Metrics
}

func (c *ControllerConfig) LimitedScope() bool {
	return c.FederationNamespaces.TargetNamespace != metav1.NamespaceAll
}
