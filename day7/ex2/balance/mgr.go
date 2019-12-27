/*******************************************************
 *  File        :   mgr.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/23 3:15 下午
 *  Notes       :	实例管理器
 *******************************************************/
package balance

import "fmt"

type BalanceMgr struct {
	allBalance map[string]Balancer
}

var (
	//默认管理器实体
	mgr BalanceMgr = BalanceMgr{
		allBalance: make(map[string]Balancer),
	}
)

//在管理器中注册均衡器
func (p *BalanceMgr) registerBalancer(name string, b Balancer) {
	p.allBalance[name] = b
}

//在管理器中注册实例
func RegisterBalance(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balance, ok := mgr.allBalance[name]
	if !ok {
		err = fmt.Errorf("not found ", name)
		fmt.Println("not found ", name)
		return
	}
	inst, err = balance.DoBalance(insts)
	if err != nil {
		err = fmt.Errorf(" errors", name)
		fmt.Println(" errors", name)
		return
	}
	return
}
