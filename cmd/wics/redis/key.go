package redis

type StoreKey string

const (
	userStoreKey StoreKey = "users"
	networkStoreKey StoreKey = "networks"
	orgGroupStoreKey StoreKey = "org_groups"
	setupKeyStoreKey StoreKey = "setup_keys"
)
