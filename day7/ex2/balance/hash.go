/*******************************************************
 *  File        :   hash.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/25 3:47 下午
 *  Notes       :
 *******************************************************/
package balance

import (
	"errors"
	"fmt"
	"hash/crc32"
	"math/rand"
)

func init() {
	RegisterBalance("hash", &HashBalancer{
		key: fmt.Sprintf("%d", rand.Int31()),
	})
}

type HashBalancer struct {
	key string
}

func (p *HashBalancer) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(key) != 0 {
		p.key = key[0]
	}
	lens := len(insts)
	if lens == 0 {
		err = errors.New("no backend instance")
		return
	}
	table := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(p.key), table)
	index := int(hashVal) % lens
	inst = insts[index]
	fmt.Println("key val: %s ", p.key)
	return
}
