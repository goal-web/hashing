package hashing

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
)

type Md5 struct {
	salt string
}

func (md5 *Md5) mixWithSalt(value string) string {
	return value + md5.salt
}

func (md5 *Md5) Info(_ string) contracts.Fields {
	return nil
}

func (md5 *Md5) Make(value string, _ contracts.Fields) string {
	return utils.Md5(md5.mixWithSalt(value))
}

func (md5 *Md5) Check(value, hashedValue string, _ contracts.Fields) bool {
	return md5.Make(value, nil) == hashedValue
}
