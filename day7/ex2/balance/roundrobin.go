/*******************************************************
 *  File        :   roundrobin.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/23 2:59 下午
 *  Notes       :	轮询负载均衡算法
 *******************************************************/
package balance

import "errors"

type RoundRobinBalance struct {
	curIndex int //当前位置
}

func init() {
	RegisterBalance("round", &RoundRobinBalance{})
}

func (p *RoundRobinBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("NO instance")
		return
	}
	lens := len(insts)
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	inst = insts[p.curIndex]
	p.curIndex++
	return
}
