CREATE TABLE IF NOT EXISTS logs
(
    Timestamp DateTime,
    DatasetId LowCardinality(String),
    TraceId String,
    ParentId String,
    TeamId LowCardinality(String),
    Level LowCardinality(String),
    Resource LowCardinality(String),
    Body String,
    Metadata Map(LowCardinality(String), String),
    RetentionDays UInt32 DEFAULT 90,
 
    INDEX idx_trace_id TraceId TYPE bloom_filter GRANULARITY 1,
    INDEX idx_parent_id ParentId TYPE bloom_filter GRANULARITY 1,
    INDEX idx_log_attr_key mapKeys(Metadata) TYPE bloom_filter GRANULARITY 1,
    INDEX idx_log_attr_value mapValues(Metadata) TYPE bloom_filter GRANULARITY 1,
    INDEX idx_body Body TYPE tokenbf_v1(32768, 3, 0) GRANULARITY 1
)
ENGINE = MergeTree
PARTITION BY toDate(Timestamp)
ORDER BY (TeamId, DatasetId, Timestamp)
TTL Timestamp + toIntervalDay(RetentionDays)
