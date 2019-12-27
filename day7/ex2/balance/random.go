/*******************************************************
 *  File        :   random.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/23 2:52 下午
 *  Notes       :	随机负载均衡算法
 *******************************************************/
package balance

import (
	"errors"
	"math/rand"
)

func init() {
	RegisterBalance("random", &RandomBalance{})
}

//随机负载均衡算法
type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}
