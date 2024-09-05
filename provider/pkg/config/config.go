package config

import (
	"context"
	"fmt"
	"os"

	goprovider "github.com/pulumi/pulumi-go-provider"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type OPNSenseProviderConfig struct {
	Address string `pulumi:"address" provider:"address"`
	Key     string `pulumi:"key" provider:"secret"`
	Secret  string `pulumi:"secret" provider:"secret"`
}

func (c *OPNSenseProviderConfig) Annotate(a infer.Annotator) {
	a.Describe(&c.Address, "OPNSense IP Address and Port")
	a.Describe(&c.Key, "OPNSense API Key")
	a.Describe(&c.Secret, "OPNSense API Secret")
}

func (c *OPNSenseProviderConfig) Configure(ctx context.Context) error {
	goprovider.GetLogger(ctx).Infof("Configuring OPNSense provider")
	if c.Address == "" {
		Address, exists := os.LookupEnv("OPNSENSE_ADDRESS")
		if !exists {
			return fmt.Errorf("address is required for OPNSense provider")
		}
		c.Address = Address
	}
	if c.Key == "" {
		Key, exists := os.LookupEnv("OPNSENSE_KEY")
		if !exists {
			return fmt.Errorf("key is required for OPNSense provider")
		}
		c.Key = Key
	}
	if c.Secret == "" {
		Secret, exists := os.LookupEnv("OPNSENSE_SECRET")
		if !exists {
			return fmt.Errorf("secret is required for OPNSense provider")
		}
		c.Secret = Secret
	}
	return nil
}
