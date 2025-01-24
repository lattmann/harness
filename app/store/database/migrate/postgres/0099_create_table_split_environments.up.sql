CREATE TABLE split_environments (
    split_environment_id                          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    split_environment_name                        TEXT NOT NULL,
    split_environment_production                  BOOLEAN NOT NULL DEFAULT FALSE,
    split_environment_workspace_id                UUID NOT NULL,
    split_environment_principal_id                INTEGER NOT NULL,
    split_environment_status                      TEXT NOT NULL,
    split_environment_created                     BIGINT NOT NULL,
	split_environment_updated                     BIGINT

    ,CONSTRAINT fk_split_environment_principal_id FOREIGN KEY (split_environment_principal_id)
    REFERENCES principals (principal_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE CASCADE

    ,CONSTRAINT fk_split_environment_workspace_id FOREIGN KEY (split_environment_workspace_id)
    REFERENCES split_workspaces (split_workspace_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE CASCADE
);
