package security

// CacheHashKey ->
func CacheHashKey(tokenID string) string {
	return "app:redis:" + tokenID
}

// CacheHashField ->
func CacheHashField() string {
	return "token"
}
