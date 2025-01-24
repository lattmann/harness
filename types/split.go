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

package types

import "github.com/google/uuid"

type SplitWorkspace struct {
	ID                       uuid.UUID `json:"id"`
	Name                     string    `json:"name"`
	RequiresTitleAndComments bool      `json:"requires_title_and_comments"`
	Created                  int64     `json:"created"`
	Updated                  int64     `json:"updated"`
}

type SplitEnvironment struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Production  bool      `json:"production"`
	WorkspaceID uuid.UUID `json:"workspace_id"`
	PrincipalID int64     `json:"principal_id"`
	Status      string    `json:"status"`
	Created     int64     `json:"created"`
	Updated     int64     `json:"updated"`
}

type SplitTrafficType struct {
	ID                 uuid.UUID `json:"id"`
	Name               string    `json:"name"`
	WorkspaceID        uuid.UUID `json:"workspace_id"`
	DisplayAttributeID string    `json:"display_attribute_id"`
	Created            int64     `json:"created"`
	Updated            int64     `json:"updated"`
}

type SplitSegment struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	EnvironmentID uuid.UUID `json:"environment_id"`
	TrafficTypeID uuid.UUID `json:"traffic_type_id"`
	Created       int64     `json:"created"`
	Updated       int64     `json:"updated"`
}
