/*******************************************************
 *  File        :   instance.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/23 2:48 下午
 *  Notes       :	实例主机
 *******************************************************/
package balance

//实例
type Instance struct {
	host string
	port int
}

//创建新的实例
func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

//获取主机地址
func (p *Instance) GetHOst() string {
	return p.host
}

//获取主机端口
func (p *Instance) GetPort() int {
	return p.port
}
