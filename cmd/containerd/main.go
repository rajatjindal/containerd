/*
   Copyright The containerd Authors.

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

package main

import (
	"crypto"
	"fmt"
	"os"

	_ "github.com/containerd/containerd/cmd/containerd/builtins"
	"github.com/containerd/containerd/cmd/containerd/command"
	"github.com/containerd/containerd/pkg/hasher"
	"github.com/containerd/containerd/pkg/seed" //nolint:staticcheck // Global math/rand seed is deprecated, but still used by external dependencies
)

func init() {
	//nolint:staticcheck // Global math/rand seed is deprecated, but still used by external dependencies
	seed.WithTimeAndRand()
	crypto.RegisterHash(crypto.SHA256, hasher.NewSHA256)
}

func main() {
	app := command.App()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "containerd: %s\n", err)
		os.Exit(1)
	}
}
