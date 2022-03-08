package flagtype

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type portValue struct{ n *uint16 }

const DefaultGrpcServerPort = 443
const DefaultSignalingServerPort = 10000
const DefaultApiPort = 8080

func PortValue(dst *uint16, defaultPort uint16) flag.Value {
	*dst = defaultPort
	return portValue{dst}
}

func (p portValue) String() string {
	if p.n == nil {
		return ""
	}
	return fmt.Sprint(*p.n)
}

func (p portValue) Set(v string) error {
	if v == "" {
		return errors.New("can't be the empty string")
	}
	if strings.Contains(v, ":") {
		return errors.New("expecting just a port number, without a colon")
	}
	n, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return fmt.Errorf("not a valid number")
	}
	if n > math.MaxUint16 {
		return errors.New("out of range for port number")
	}
	*p.n = uint16(n)
	return nil
}
