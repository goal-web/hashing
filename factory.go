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

func (factory *Factory) Info(hashedValue string) contracts.Fields {
	return factory.Driver("default").Info(hashedValue)
}

func (factory *Factory) Make(value string, options contracts.Fields) string {
	return factory.Driver("default").Make(value, options)
}

func (factory *Factory) Check(value, hashedValue string, options contracts.Fields) bool {
	return factory.Driver("default").Check(value, hashedValue, options)
}

func (factory *Factory) getConfig(name string) contracts.Fields {
	fields, _ := factory.config.Get(
		utils.IfString(name == "default", "hashing", fmt.Sprintf("hashing.hashes.%s", name)),
	).(contracts.Fields)
	return fields
}

func (factory *Factory) Driver(name string) contracts.Hasher {
	if hashed, existsHashed := factory.hashes[name]; existsHashed {
		return hashed
	}

	config := factory.getConfig(name)
	driver := utils.GetStringField(config, "driver", "bcrypt")
	driveProvider, existsProvider := factory.drivers[driver]

	if !existsProvider {
		logs.WithFields(nil).Fatal(fmt.Sprintf("不支持的哈希驱动：%s", driver))
	}

	factory.hashes[name] = driveProvider(config)

	return factory.hashes[name]
}

func (factory *Factory) Extend(driver string, hashedProvider contracts.HasherProvider) {
	factory.drivers[driver] = hashedProvider
}
