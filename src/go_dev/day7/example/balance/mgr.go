package balance

import "fmt"

//方法管理
type BalancerMgr struct {
	allBalancer map[string]Balancer
}

//map初始化
var mgr = BalancerMgr{
	allBalancer: make(map[string]Balancer),
}

//接口方法管理
func (p *BalancerMgr) registerBalancer(name string, b Balancer) {
	p.allBalancer[name] = b
}

//方法注册 外部可调用
func RegisterBalancer(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	//判断 map中的KEY
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("Not fond %s balancer", name)
		return
	}
	fmt.Printf("use %s balance", name)
	inst, err = balancer.DoBalance(insts)
	return
}
