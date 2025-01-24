CREATE TABLE split_traffic_types (
    split_traffic_type_id                         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    split_traffic_type_name                       TEXT NOT NULL,
    split_traffic_type_workspace_id               UUID NOT NULL,
    split_traffic_type_display_attribute_id       TEXT,
    split_traffic_type_created                    BIGINT NOT NULL,
    split_traffic_type_updated                    BIGINT

    ,CONSTRAINT fk_split_traffic_type_workspace_id FOREIGN KEY (split_traffic_type_workspace_id)
        REFERENCES split_workspaces (split_workspace_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);
