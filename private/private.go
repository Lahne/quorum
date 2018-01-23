package private

import (
	"os"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private/constellation"
)

type PrivateTransactionManager interface {
	Send(data []byte, from string, to []string) ([]byte, error)
	Receive(data []byte) ([]byte, error)
}

func FromEnvironmentOrNil(name string) PrivateTransactionManager {
	cfgPath := os.Getenv(name)
	log.Info("cfgPath: " + cfgPath)
	if cfgPath == "" {
		return nil
	}
	return constellation.MustNew(cfgPath)
}

var P = FromEnvironmentOrNil("PRIVATE_CONFIG")
