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
	"github.com/google/wire"
	"github.com/harness/gitness/app/auth/authz"
	"github.com/harness/gitness/app/store"

	// "github.com/harness/gitness/app/url"
	"github.com/harness/gitness/audit"
	"github.com/harness/gitness/store/database/dbtx"
)

var WireSet = wire.NewSet(
	ProvideController,
)

func ProvideController(
	authorizer authz.Authorizer,
	// urlProvider url.Provider,
	auditService audit.Service,
	tx dbtx.Transactor,
	splitWorkspaceStore store.SplitWorkspaceStore,
	// splitEnvironmentStore store.SplitEnvironmentStore,
	// splitTrafficTypeStore store.SplitTrafficTypeStore,
	// splitSegmentStore store.SplitSegmentStore,
) *Controller {
	return newController(
		authorizer,
		// urlProvider,
		auditService,
		tx,
		splitWorkspaceStore,
		// splitEnvironmentStore,
		// splitTrafficTypeStore,
		// splitSegmentStore,
	)
}
