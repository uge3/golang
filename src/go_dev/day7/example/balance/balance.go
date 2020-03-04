package balance

//多态接口
type Balancer interface {
	DoBalance([]*Instance, ...string) (inst *Instance, err error)
}
