package key

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// we can create setupkey with server state private key.
//

const (
	setupKeyPrefix = "w_"
)

// key usage types.
type SetupKeyType string

const (
	DefaultKey SetupKeyType = "default"
)

type PermissionType string

// like unix permission
const (
	RWXKey PermissionType = "admin"
	RWKey  PermissionType = "manager"
	RKey   PermissionType = "default"
)

// structure of jwt in SetupKey
type customClaims struct {
	KeyType    SetupKeyType		`json:"key_type"`
	Group 	   string 			`json:"group"`
	Job 	   string 			`json:"job"`
	Permission PermissionType	`json:"permission"`
	Revoked    bool         	`json:"revoked"`

	CreatedAt  time.Time    	`json:"created_at"`
	LastusedAt time.Time    	`json:"lastused_at"`

	*jwt.StandardClaims
}

type SetupKey struct {
	Key          string
	SignedString string
}

func NewSetupKey(sub, group, job string, permissionType PermissionType) (*SetupKey, error) {
	var (
		now = time.Now().Unix()
		// expires 1 weak of setupkey
		exp = time.Now().Add(time.Hour * 24 * 7).Unix()
		sb  strings.Builder
	)

	id := strings.ToUpper(uuid.New().String())

	claims := customClaims{
		DefaultKey,
		group,
		job,
		permissionType,
		false,
		time.Now(),
		time.Now(),
		&jwt.StandardClaims{
			Id:        id,
			Issuer:    sub,
			IssuedAt:  now,
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString := os.Getenv("JWT_SECRET")
	ss, err := token.SignedString([]byte(signedString))
	if err != nil {
		return nil, err
	}

	sb.WriteString(setupKeyPrefix)
	sb.WriteString(ss)

	return &SetupKey{
		Key:          sb.String(),
		SignedString: signedString,
	}, nil
}

func (s *SetupKey) SetupKey() string {
	setupkey := strings.Trim(s.Key, setupKeyPrefix)
	return setupkey
}

func (s *SetupKey) KeyType() (SetupKeyType, error) {
	c, err := getCustomClaims(s.Key, s.SignedString)
	return c.KeyType, err
}

func (s *SetupKey) IsRevoked() (bool, error) {
	c, err := getCustomClaims(s.Key, s.SignedString)
	return c.Revoked, err
}

func (s *SetupKey) IsExpired() (bool, error) {
	c, err := getCustomClaims(s.Key, s.SignedString)
	return time.Now().After(time.Unix(c.ExpiresAt, 0)), err
}

func (s *SetupKey) ID() (string, error) {
	c, err := getCustomClaims(s.Key, s.SignedString)
	return c.Id, err
}

func getCustomClaims(setupkey, signedString string) (*customClaims, error) {
	sk := strings.Trim(setupkey, setupKeyPrefix)

	token, err := jwt.ParseWithClaims(sk, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signedString), nil
	})

	if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
