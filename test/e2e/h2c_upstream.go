/*
Copyright 2017 Frederic Branczyk All rights reserved.

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

package e2e

import (
	"testing"

	"github.com/brancz/kube-rbac-proxy/test/kubetest"
)

func testH2CUpstream(s *kubetest.Suite) kubetest.TestSuite {
	return func(t *testing.T) {
		command := `curl --connect-timeout 5 -v -s -k --fail -H "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" https://kube-rbac-proxy.default.svc.cluster.local:8443/metrics`

		kubetest.Scenario{
			Name: "With H2C Upstream",

			Given: kubetest.Setups(
				kubetest.CreatedManifests(
					s.KubeClient,
					"h2c-upstream/clusterRole.yaml",
					"h2c-upstream/clusterRoleBinding.yaml",
					"h2c-upstream/deployment.yaml",
					"h2c-upstream/service.yaml",
					"h2c-upstream/serviceAccount.yaml",
					"h2c-upstream/clusterRole-client.yaml",
					"h2c-upstream/clusterRoleBinding-client.yaml",
				),
			),
			When: kubetest.Conditions(
				kubetest.PodsAreReady(
					s.KubeClient,
					1,
					"app=kube-rbac-proxy",
				),
				kubetest.ServiceIsReady(
					s.KubeClient,
					"kube-rbac-proxy",
				),
			),
			Then: kubetest.Checks(
				ClientSucceeds(
					s.KubeClient,
					command,
					nil,
				),
			),
		}.Run(t)
	}
}
