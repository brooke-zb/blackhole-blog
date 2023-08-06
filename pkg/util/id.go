package util

import "github.com/yitter/idgenerator-go/idgen"

func initIdGenerator() {
	idgen.SetIdGenerator(idgen.NewIdGeneratorOptions(0))
}

func NextId() uint64 {
	return uint64(idgen.NextId())
}
