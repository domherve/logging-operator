// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package output_test

import (
	"testing"

	"github.com/banzaicloud/logging-operator/pkg/sdk/model/output"
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/render"
	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/require"
)

func TestFormatSingleValueConfig(t *testing.T) {
	CONFIG := []byte(`
path: /tmp/logs/${tag}/%Y/%m/%d.%H.%M
format:
  type: single_value
  add_newline: true
  message_key: msg
buffer:
  timekey: 1m
  timekey_wait: 30s
  timekey_use_utc: true
`)
	expected := `
  <match **>
	@type file
	@id test
	add_path_suffix true
	path /tmp/logs/${tag}/%Y/%m/%d.%H.%M
    <buffer tag,time>
      @type file
	  chunk_limit_size 8MB
      path /buffers/test.*.buffer
      retry_forever true
      timekey 1m
      timekey_use_utc true
      timekey_wait 30s
    </buffer>
    <format>
      @type single_value
      add_newline true
      message_key msg
    </format>
  </match>
`
	f := &output.FileOutputConfig{}
	require.NoError(t, yaml.Unmarshal(CONFIG, f))
	test := render.NewOutputPluginTest(t, f)
	test.DiffResult(expected)
}
