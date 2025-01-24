// Copyright 2025 Harness, Inc.
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

package split

import (
	"github.com/harness/gitness/app/auth/authz"
	"github.com/harness/gitness/app/store"
	"github.com/harness/gitness/audit"
	"github.com/harness/gitness/store/database/dbtx"
	"github.com/harness/gitness/types"
)

type WorkspaceOutput struct {
	types.SplitWorkspace
}

type EnvironmentOutput struct {
	types.SplitEnvironment
}

type TrafficTypeOutput struct {
	types.SplitTrafficType
}

type SegmentOutput struct {
	types.SplitSegment
}

type Controller struct {
	authorizer authz.Authorizer
	// urlProvider           url.Provider
	auditService        audit.Service
	tx                  dbtx.Transactor
	splitWorkspaceStore store.SplitWorkspaceStore
	// splitEnvironmentStore store.SplitEnvironmentStore
	// splitTrafficTypeStore store.SplitTrafficTypeStore
	// splitSegmentStore     store.SplitSegmentStore
}

func newController(
	authorizer authz.Authorizer,
	// urlProvider url.Provider,
	auditService audit.Service,
	tx dbtx.Transactor,
	splitWorkspaceStore store.SplitWorkspaceStore,
	// splitEnvironmentStore store.SplitEnvironmentStore,
	// splitTrafficTypeStore store.SplitTrafficTypeStore,
	// splitSegmentStore store.SplitSegmentStore,
) *Controller {
	return &Controller{
		authorizer: authorizer,
		// urlProvider:           urlProvider,
		auditService:        auditService,
		tx:                  tx,
		splitWorkspaceStore: splitWorkspaceStore,
		// splitEnvironmentStore: splitEnvironmentStore,
		// splitTrafficTypeStore: splitTrafficTypeStore,
		// splitSegmentStore:     splitSegmentStore,
	}
}
