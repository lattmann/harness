CREATE TABLE split_segments (
    split_segment_id                            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    split_segment_name                          TEXT NOT NULL,
    split_segment_description                   TEXT DEFAULT '',
    split_segment_environment_id                UUID NOT NULL,
    split_segment_traffic_type_id               UUID NOT NULL,
    split_segment_created                       BIGINT NOT NULL,
    split_segment_updated                       BIGINT

    ,CONSTRAINT fk_split_segment_environment_id FOREIGN KEY (split_segment_environment_id)
        REFERENCES split_environments (split_environment_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE

    ,CONSTRAINT fk_split_segment_traffic_type_id FOREIGN KEY (split_segment_traffic_type_id)
        REFERENCES split_traffic_types (split_traffic_type_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);
