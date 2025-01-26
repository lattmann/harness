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
	"github.com/pkg/errors"
)

// SPLIT WORKSPACE STORE
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

	query, arg, err := db.BindNamed(sqlQuery, mapInternalSplitWorkspace(workspace, uuid.New()))
	if err != nil {
		return database.ProcessSQLErrorf(ctx, err, "Failed to bind split workspace object")
	}

	if err = db.QueryRowContext(ctx, query, arg...).Scan(&workspace.ID); err != nil {
		return database.ProcessSQLErrorf(ctx, err, "Failed to insert split workspace")
	}

	return nil
}

// Find Finds a split workspace in the database.
func (s *SplitWorkspaceStore) Find(
	ctx context.Context,
	id uuid.UUID,
) (*types.SplitWorkspace, error) {

	stmt := database.Builder.
		Select(splitWorkspaceColumns).
		From("split_workspaces").
		Where("split_workspace_id = ?", id)

	db := dbtx.GetAccessor(ctx, s.db)

	dst := new(splitWorkspace)
	sql, args, err := stmt.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to convert query to sql")
	}

	if err = db.GetContext(ctx, dst, sql, args...); err != nil {
		return nil, database.ProcessSQLErrorf(ctx, err, "Failed to find split workspace")
	}

	return s.mapToSplitWorkspace(ctx, dst)
}

type splitWorkspace struct {
	ID                       uuid.UUID `db:"split_workspace_id"`
	Name                     string    `db:"split_workspace_name"`
	RequiresTitleAndComments bool      `db:"split_workspace_requires_title_and_comments"`
	Created                  int64     `db:"split_workspace_created"`
	Updated                  int64     `db:"split_workspace_updated"`
}

func mapInternalSplitWorkspace(u *types.SplitWorkspace, splitWorkspaceID uuid.UUID) *splitWorkspace {
	return &splitWorkspace{
		ID:                       splitWorkspaceID,
		Name:                     u.Name,
		RequiresTitleAndComments: u.RequiresTitleAndComments,
		Created:                  u.Created,
		Updated:                  u.Updated,
	}
}

func (s *SplitWorkspaceStore) mapToSplitWorkspace(ctx context.Context, dst *splitWorkspace) (*types.SplitWorkspace, error) {
	return &types.SplitWorkspace{
		ID:                       dst.ID,
		Name:                     dst.Name,
		RequiresTitleAndComments: dst.RequiresTitleAndComments,
		Created:                  dst.Created,
		Updated:                  dst.Updated,
	}, nil
}

// SPLIT ENVIRONMENT STORE

// SPLIT TRAFFIC TYPE STORE

// SPLIT SEGMENT STORE
