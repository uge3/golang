package balance

//主机实例
type Instance struct {
	host string
	port int
}

//新建主机方法
func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

//外部接口方法,获取主机
func (p *Instance) Gethost() string {
	return p.host
}

//外部接口方法,获取端口
func (p *Instance) Getport() int {
	return p.port
}
