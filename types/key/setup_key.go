package key

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// we can create setupkey with server state private key.
//

const (
	SetupKeyPrefix = "w_"
)

// key usage types.
type SetupKeyType int8

const (
	DefaultKey SetupKeyType = 3
)

type PermissionType int8

// like unix permission
const (
	RWXKey PermissionType = 3
	RWKey  PermissionType = 2
	RKey   PermissionType = 1
)

// structure of jwt in SetupKey
type customClaims struct {
	UserID     uint         `json:"user_id"`
	ProviderID string       `json:"provider_id"`
	KeyType    SetupKeyType `json:"key_type"`

	UserGroupID uint           `json:"user_group_id"`
	OrgGroupID  uint           `json:"org_group_id"`
	Job         string         `json:"job"`
	Permission  PermissionType `json:"permission"`
	Revoked     bool           `json:"revoked"`

	CreatedAt  time.Time `json:"created_at"`
	LastusedAt time.Time `json:"lastused_at"`

	*jwt.StandardClaims
}

type SetupKey struct {
	Key          string
	signedString string
}

func NewSetupKey(
	userID uint, providerID, 
	job string, userGroupID, 
	orgGroupID uint, permissionType PermissionType,
	issuer, audience string,
	signedString string,
) (*SetupKey, error) {
	var (
		now = time.Now().Unix()
		// expires 1 weak of setupkey
		exp = time.Now().Add(time.Hour * 24 * 7).Unix()
		sb  strings.Builder
	)

	id := strings.ToUpper(uuid.New().String())

	claims := customClaims{
		userID,
		providerID,
		DefaultKey,
		userGroupID,
		orgGroupID,
		job,
		permissionType,
		false,
		time.Now(),
		time.Now(),
		&jwt.StandardClaims{
			Id:        id,
			Issuer:    issuer,
			Audience:  audience,
			IssuedAt:  now,
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(signedString))
	if err != nil {
		return nil, err
	}

	sb.WriteString(SetupKeyPrefix)
	sb.WriteString(ss)

	return &SetupKey{
		Key:          sb.String(),
		signedString: signedString,
	}, nil
}

func (s *SetupKey) SetupKey() string {
	setupkey := strings.Trim(s.Key, SetupKeyPrefix)
	return setupkey
}

func (s *SetupKey) KeyType() (SetupKeyType, error) {
	c, err := GetCustomClaims(s.Key, s.signedString)
	return c.KeyType, err
}

func (s *SetupKey) IsRevoked() (bool, error) {
	c, err := GetCustomClaims(s.Key, s.signedString)
	return c.Revoked, err
}

func (s *SetupKey) IsExpired() (bool, error) {
	c, err := GetCustomClaims(s.Key, s.signedString)
	return time.Now().After(time.Unix(c.ExpiresAt, 0)), err
}

func (s *SetupKey) ID() (string, error) {
	c, err := GetCustomClaims(s.Key, s.signedString)
	return c.Id, err
}

func GetCustomClaims(setupkey, signedString string) (*customClaims, error) {
	sk := strings.Trim(setupkey, SetupKeyPrefix)

	token, err := jwt.ParseWithClaims(sk, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signedString), nil
	})

	if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
