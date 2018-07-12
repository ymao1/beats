// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package http

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "http", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzMlDFT+zAMxfd8Cl323v2HThn+MxPH0I1jcOPXxm1iG0sp5NtzbkhoUpejB+XwGFm/92QpWtAeXUGViM+IxEiNgvK71eohz4g0uAzGi3G2oP8ZEVEMUeN0WyMjCqihGAVtVUbEEDF2ywU95sx1/pQRbQxqzcUxd0FWNRjV4pHOx+zg2uFLQnNKOSUFPLdgGb+ngBehwzmW9E4iYzcuNCrePLk2159UA6URZszehlvvUMos9KkXolWF3lHPZWJYSQo3kMrppPAe3YsL89hXlXswtQydFF473d1A1quudkpPC/7oNHtnGT/S6h71B3sdUMIcLjx76TRu8OwsSlpO0fGqGh8XwvLfMunIV0HxNZ5G4r0T2rjWftN0wsAvDelZpwbdHU+m6eohjfnxDwymZJzO03TXDufSZmSEw2RSrzZyRhgNrCEpC28BAAD//9KJlTU="
}