/*
 * Copyright (C) 2016 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package client

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/skydive-project/skydive/graffiti/api/client"
	"github.com/skydive-project/skydive/graffiti/logging"
)

func createMetadataJSONPatch(addMetadata, removeMetadata []string) (patch client.JSONPatch, _ error) {
	for _, add := range addMetadata {
		split := strings.SplitN(add, "=", 2)
		if len(split) < 2 {
			return nil, fmt.Errorf("metadata to add should be of the form k1=v1, got %s", add)
		}

		var value interface{}
		if err := json.Unmarshal([]byte(split[1]), &value); err != nil {
			value = split[1]
		}
		patch = append(patch, client.NewPatchOperation("add", "/Metadata/"+strings.Replace(split[0], "/", ".", -1), value))
	}

	for _, remove := range removeMetadata {
		patch = append(patch, client.NewPatchOperation("remove", "/Metadata/"+strings.Replace(remove, "/", ".", -1)))
	}

	return patch, nil
}

func printJSON(obj interface{}) {
	s, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		logging.GetLogger().Error(err)
		os.Exit(1)
	}
	fmt.Println(string(s))
}
