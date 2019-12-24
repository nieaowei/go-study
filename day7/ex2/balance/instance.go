/*******************************************************
 *  File        :   instance.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/23 2:48 下午
 *  Notes       :
 *******************************************************/
package balance

type Instance struct {
	host string
	port int
}

func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

func (p *Instance) GetHOst() string {
	return p.host
}

func (p *Instance) GetPort() int {
	return p.port
}
