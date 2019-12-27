/*******************************************************
 *  File        :   balance.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/23 2:46 下午
 *  Notes       :	负载均衡接口
 *******************************************************/
package balance

type Balancer interface {
	DoBalance([]*Instance, ...string) (*Instance, error)
}
