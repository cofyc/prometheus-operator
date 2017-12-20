// Copyright 2016 The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v2

import (
	v1 "github.com/coreos/prometheus-operator/pkg/client/monitoring/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
)

// v2.Prometheus has some fields v1.Prometheus don't have, so we need to manully write this conversion function through we can simply drop these fields.
func Convert_v2_PrometheusSpec_To_v1_PrometheusSpec(in *PrometheusSpec, out *v1.PrometheusSpec, s conversion.Scope) error {
	return autoConvert_v2_PrometheusSpec_To_v1_PrometheusSpec(in, out, s)
}
