/*******************************************************
 *  File        :   balance.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/23 2:46 下午
 *  Notes       :
 *******************************************************/
package balance

type Balancer interface {
	DoBalance([]*Instance) (*Instance, error)
}
