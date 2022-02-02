package redis

type StoreKey string

const (
	userStoreKey StoreKey = "users"
	networkStoreKey StoreKey = "networks"
	groupStoreKey StoreKey = "groups"
	setupKeyStoreKey StoreKey = "setup_keys"
)
