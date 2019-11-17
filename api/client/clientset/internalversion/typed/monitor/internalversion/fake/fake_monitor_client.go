/*
 * Copyright 2019 THL A29 Limited, a Tencent company.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	internalversion "tkestack.io/tke/api/client/clientset/internalversion/typed/monitor/internalversion"
)

type FakeMonitor struct {
	*testing.Fake
}

func (c *FakeMonitor) ConfigMaps() internalversion.ConfigMapInterface {
	return &FakeConfigMaps{c}
}

func (c *FakeMonitor) Metrics() internalversion.MetricInterface {
	return &FakeMetrics{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMonitor) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
