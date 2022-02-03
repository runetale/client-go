package redis

type StoreKey string

const (
	UserStoreKey StoreKey = "users"
	NetworkStoreKey StoreKey = "networks"
	OrgGroupStoreKey StoreKey = "org_groups"
	SetupKeyStoreKey StoreKey = "setup_keys"
)
