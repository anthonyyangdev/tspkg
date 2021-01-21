package api

func Ok(statusCode int) bool {
	return 200 <= statusCode && statusCode < 300
}