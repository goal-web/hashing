package hashing

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/logs"
	"github.com/goal-web/supports/utils"
)

type Factory struct {
	config  contracts.Config
	hashes  map[string]contracts.Hasher
	drivers map[string]contracts.HasherProvider
}

func (this *Factory) Info(hashedValue string) contracts.Fields {
	return this.Driver("default").Info(hashedValue)
}

func (this *Factory) Make(value string, options contracts.Fields) string {
	return this.Driver("default").Make(value, options)
}

func (this *Factory) Check(value, hashedValue string, options contracts.Fields) bool {
	return this.Driver("default").Check(value, hashedValue, options)
}

func (this Factory) getConfig(name string) contracts.Fields {
	return this.config.GetFields(
		utils.IfString(name == "default", "hashing", fmt.Sprintf("hashing.hashes.%s", name)),
	)
}

func (this *Factory) Driver(name string) contracts.Hasher {
	if hashed, existsHashed := this.hashes[name]; existsHashed {
		return hashed
	}

	config := this.getConfig(name)
	driver := utils.GetStringField(config, "driver", "bcrypt")
	driveProvider, existsProvider := this.drivers[driver]

	if !existsProvider {
		logs.WithFields(nil).Fatal(fmt.Sprintf("不支持的哈希驱动：%s", driver))
	}

	this.hashes[name] = driveProvider(config)

	return this.hashes[name]
}

func (this *Factory) Extend(driver string, hashedProvider contracts.HasherProvider) {
	this.drivers[driver] = hashedProvider
}
