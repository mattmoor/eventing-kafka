package env

// Package Constants
const (
	// Eventing-Kafka Configuration
	ServiceAccountEnvVarKey = "SERVICE_ACCOUNT"
	MetricsPortEnvVarKey    = "METRICS_PORT"
	HealthPortEnvVarKey     = "HEALTH_PORT"

	// Kafka Authorization
	KafkaBrokerEnvVarKey   = "KAFKA_BROKERS"
	KafkaUsernameEnvVarKey = "KAFKA_USERNAME"
	KafkaPasswordEnvVarKey = "KAFKA_PASSWORD"

	// Kafka Configuration
	KafkaProviderEnvVarKey                   = "KAFKA_PROVIDER"
	KafkaOffsetCommitMessageCountEnvVarKey   = "KAFKA_OFFSET_COMMIT_MESSAGE_COUNT"
	KafkaOffsetCommitDurationMillisEnvVarKey = "KAFKA_OFFSET_COMMIT_DURATION_MILLIS"
	KafkaTopicEnvVarKey                      = "KAFKA_TOPIC"

	// Dispatcher Configuration
	ChannelKeyEnvVarKey           = "CHANNEL_KEY"
	ServiceNameEnvVarKey          = "SERVICE_NAME"
	ExponentialBackoffEnvVarKey   = "EXPONENTIAL_BACKOFF"
	InitialRetryIntervalEnvVarKey = "INITIAL_RETRY_INTERVAL"
	MaxRetryTimeEnvVarKey         = "MAX_RETRY_TIME"
)

