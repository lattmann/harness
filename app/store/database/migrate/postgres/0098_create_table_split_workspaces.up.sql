CREATE TABLE split_workspaces (
    split_workspace_id                          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    split_workspace_name                        TEXT NOT NULL,
    split_workspace_requires_title_and_comments BOOLEAN NOT NULL DEFAULT TRUE,
    split_workspace_created                     BIGINT NOT NULL,
	split_workspace_updated                     BIGINT
);
