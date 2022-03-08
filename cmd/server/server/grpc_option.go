package grpcserver

import (
	"crypto/tls"
	"time"

	"github.com/Notch-Technologies/wizy/cert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

func NewGrpcServerOption(
	certConfig *cert.CertConfig,
) ([]grpc.ServerOption, error) {
	var opts []grpc.ServerOption

	// config default option
	//
	enforcementPol := keepalive.EnforcementPolicy{
		MinTime:             10 * time.Second,
		PermitWithoutStream: true,
	}

	serverParams := keepalive.ServerParameters{
		MaxConnectionIdle:     10 * time.Second,
		MaxConnectionAgeGrace: 5 * time.Second,
		Time:                  5 * time.Second,
		Timeout:               3 * time.Second,
	}
	opts = append(opts, grpc.KeepaliveEnforcementPolicy(enforcementPol), grpc.KeepaliveParams(serverParams))

	// grpc tls config with custom
	//
	certManager := certConfig.CreateCertManager()
	if certConfig.Domain != "" {
		t := credentials.NewTLS(certManager.TLSConfig())
		opts = append(opts, grpc.Creds(t))
	} else if certConfig.File != "" && certConfig.Key != "" {
		tlsConfig, err := createTLSConfig(certConfig.File, certConfig.Key)
		if err != nil {
			return nil, err
		}
		transportCredentials := credentials.NewTLS(tlsConfig)
		opts = append(opts, grpc.Creds(transportCredentials))
	}

	return opts, nil
}

func createTLSConfig(file, key string) (*tls.Config, error) {
	serverCert, err := tls.LoadX509KeyPair(file, key)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return config, nil
}

