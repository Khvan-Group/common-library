package constants

const (
	// Headers

	APPLICATION_JSON      = "application/json"
	CONTENT_TYPE          = "Content-type"
	AUTHORIZATION         = "Authorization"
	BEARER                = "Bearer "
	X_IS_INTERNAL_SERVICE = "X-Is-Internal-Service"

	// User Roles

	ADMIN     = "ADMIN"
	MODERATOR = "MODERATOR"
	USER      = "USER"

	// RabbitMQ Queues

	RABBIT_CREATE_WALLET_QUEUE = "RABBIT_CREATE_WALLET_QUEUE"
	RABBIT_UPDATE_WALLET_QUEUE = "RABBIT_UPDATE_WALLET_QUEUE"

	MINIO_BUCKET_AVATARS = "MINIO_BUCKET_AVATARS"
	MINIO_URL            = "MINIO_URL"
	MINIO_USER           = "MINIO_USER"
	MINIO_PASSWORD       = "MINIO_PASSWORD"
)
