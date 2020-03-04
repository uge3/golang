package main

//一次性哈希
import (
	"fmt"
	"go_dev/day7/example/balance"
	"hash/crc32"
	"math/rand"
)

type HashBalance struct {
}

func init() {
	balance.RegisterBalancer("hash", &HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {
	var defKey string = fmt.Sprintf("%d", rand.Int()) //默认的key
	if len(key) > 0 {
		defKey = key[0]
		//err = fmt.Errorf("hash balance must pass hash key")
	}

	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("No backend instance")
		return
	}
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens //取余
	inst = insts[index]
	return
}
