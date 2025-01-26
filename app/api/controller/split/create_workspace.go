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
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/harness/gitness/app/auth"
	"github.com/harness/gitness/types"
)

type CreateWorkspaceInput struct {
	Name                     string `json:"name"`
	RequiresTitleAndComments bool   `json:"requiresTitleAndComments"`
}

func (c *Controller) CreateWorkspace(
	ctx context.Context,
	session *auth.Session,
	in *CreateWorkspaceInput,
) (*WorkspaceOutput, error) {
	if err := c.sanitizeCreateInput(in); err != nil {
		return nil, fmt.Errorf("failed to sanitize input: %w", err)
	}

	var err error
	var workspace *types.SplitWorkspace

	err = c.tx.WithTx(ctx, func(ctx context.Context) error {
		workspace, err = c.createWorkspaceInnerInTX(ctx, session, in)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &WorkspaceOutput{SplitWorkspace: *workspace}, nil
}

func (c *Controller) createWorkspaceInnerInTX(
	ctx context.Context,
	session *auth.Session,
	in *CreateWorkspaceInput,
) (*types.SplitWorkspace, error) {

	if session.Principal.ID == 0 {
		return nil, fmt.Errorf("principal id is required")
	}

	now := time.Now().UnixMilli()
	workspace := &types.SplitWorkspace{
		Name:                     in.Name,
		RequiresTitleAndComments: in.RequiresTitleAndComments,
		Created:                  now,
		Updated:                  now,
	}

	err := c.splitWorkspaceStore.Create(ctx, workspace)
	if err != nil {
		return nil, fmt.Errorf("failed to create workspace: %w", err)
	}

	return workspace, nil
}

func (c *Controller) sanitizeCreateInput(in *CreateWorkspaceInput) error {

	// TODO: default value may need to be in config
	in.RequiresTitleAndComments = in.RequiresTitleAndComments || true
	return nil
}

func (c *Controller) GetWorkspace(
	ctx context.Context,
	session *auth.Session,
	id uuid.UUID,
) (*WorkspaceOutput, error) {
	workspace, err := c.splitWorkspaceStore.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find workspace: %w", err)
	}

	return &WorkspaceOutput{SplitWorkspace: *workspace}, nil
}
