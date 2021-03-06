// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
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

package cmd

import (
	"github.com/spf13/cobra"
)

// initializeSnapstoreFlags adds the snapstore related flags to <cmd>
func initializeSnapstoreFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&storageProvider, "storage-provider", "", "snapshot storage provider")
	cmd.Flags().StringVar(&storageContainer, "store-container", "", "container which will be used as snapstore")
	cmd.Flags().StringVar(&storagePrefix, "store-prefix", "", "prefix or directory inside container under which snapstore is created")
	cmd.Flags().IntVar(&maxParallelChunkUploads, "max-parallel-chunk-uploads", 5, "maximum number of parallel chunk uploads allowed ")
	cmd.Flags().StringVar(&snapstoreTempDir, "snapstore-temp-directory", "/tmp", "temporary directory for processing")
}
