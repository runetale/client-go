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
	setupKeyPrefix = "w_"
)

// SetupKey key types.
type SetupKeyType string

const (
	defaultKey SetupKeyType = "default"
)

// structure of jwt in SetupKey
type customClaims struct {
	keytype    SetupKeyType `json:"key_type"`
	revoked    bool         `json:"revoked"`
	createdAt  time.Time    `json:"created_at"`
	lastusedAt time.Time    `json:"lastused_at"`
	*jwt.StandardClaims
}

type SetupKey struct {
	key          string
	signedString string
}

// create a SetupKey using the base64-encoded value of the server private key.
func NewSetupKey(b64key string) (*SetupKey, error) {
	var (
		now = time.Now().Unix()
		// expires 2 weak of setupkey
		exp = time.Now().Add(time.Hour * 24 * 14).Unix()
		sb  strings.Builder
	)

	id := strings.ToUpper(uuid.New().String())

	claims := customClaims{
		defaultKey,
		false,
		time.Now(),
		time.Now(),
		&jwt.StandardClaims{
			Id:        id,
			Issuer:    "wizy",
			IssuedAt:  now,
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(b64key))
	if err != nil {
		return nil, err
	}

	sb.WriteString(setupKeyPrefix)
	sb.WriteString(ss)

	return &SetupKey{
		key:          sb.String(),
		signedString: b64key,
	}, nil
}

func (s *SetupKey) SetupKey() string {
	setupkey := strings.Trim(s.key, setupKeyPrefix)
	return setupkey
}

func (s *SetupKey) KeyType() (SetupKeyType, error) {
	c, err := getCustomClaims(s.key, s.signedString)
	return c.keytype, err
}

func (s *SetupKey) IsRevoked() (bool, error) {
	c, err := getCustomClaims(s.key, s.signedString)
	return c.revoked, err
}

func (s *SetupKey) IsExpired() (bool, error) {
	c, err := getCustomClaims(s.key, s.signedString)
	return time.Now().After(time.Unix(c.ExpiresAt, 0)), err
}

func (s *SetupKey) ID() (string, error) {
	c, err := getCustomClaims(s.key, s.signedString)
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
