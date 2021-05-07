package infrastructure

import (
	"context"
	"encoding/json"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/murilocosta/agartha/internal/core"
)

func LoadConfigurationFromServer(configServerURL string, config *core.Config) error {
	// Create Etcd client instance
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{configServerURL},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}
	defer etcdCli.Close()

	// Try to retrieve application configuration from Etcd service
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := etcdCli.Get(ctx, "config")
	close()
	if err != nil {
		return err
	}

	// Read configuration from key-value pair retrieved
	err = json.Unmarshal(resp.Kvs[0].Value, &config)
	if err != nil {
		return err
	}

	return nil
}
