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

package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/harness/gitness/app/store"
	"github.com/harness/gitness/store/database"
	"github.com/harness/gitness/store/database/dbtx"
	"github.com/harness/gitness/types"
	"github.com/jmoiron/sqlx"
)

type SplitWorkspaceStore struct {
	db *sqlx.DB
}

func NewSplitWorkspaceStore(db *sqlx.DB) *SplitWorkspaceStore {
	return &SplitWorkspaceStore{db: db}
}

type SplitWorkspace struct {
	ID                       uuid.UUID `db:"split_workspace_id"`
	Name                     string    `db:"split_workspace_name"`
	RequiresTitleAndComments bool      `db:"split_workspace_requires_title_and_comments"`
	Created                  int64     `db:"split_workspace_created"`
	Updated                  int64     `db:"split_workspace_updated"`
}

var _ store.SplitWorkspaceStore = (*SplitWorkspaceStore)(nil)

const (
	splitWorkspaceColumns = `
	split_workspace_id
	,split_workspace_name
	,split_workspace_requires_title_and_comments
	,split_workspace_created
	,split_workspace_updated`

	splitWorkspaceSelectBase = `SELECT ` + splitWorkspaceColumns + ` FROM split_workspaces`
)

// Create Creates a split workspace in the database.
func (s *SplitWorkspaceStore) Create(
	ctx context.Context,
	// splitWorkspaceID uuid.UUID,
	workspace *types.SplitWorkspace,
) error {
	const sqlQuery = `
	INSERT INTO split_workspaces (
		split_workspace_name
		,split_workspace_requires_title_and_comments
		,split_workspace_created
		,split_workspace_updated
	) values (
		:split_workspace_name
		,:split_workspace_requires_title_and_comments
		,:split_workspace_created
		,:split_workspace_updated
	) RETURNING split_workspace_id`

	db := dbtx.GetAccessor(ctx, s.db)

	query, arg, err := db.BindNamed(sqlQuery, workspace)
	if err != nil {
		return database.ProcessSQLErrorf(ctx, err, "Failed to bind split workspace object")
	}

	if err = db.QueryRowContext(ctx, query, arg...).Scan(&workspace.ID); err != nil {
		return database.ProcessSQLErrorf(ctx, err, "Failed to insert split workspace")
	}

	return nil
}

// func mapInternalSplitWorkspace(u *types.SplitWorkspace, splitWorkspaceID uuid.UUID) *SplitWorkspace {
// 	return &SplitWorkspace{
// 		ID:                       splitWorkspaceID,
// 		RequiresTitleAndComments: u.RequiresTitleAndComments,
// 		Created:                  u.Created,
// 		Updated:                  u.Updated,
// 	}
// }
