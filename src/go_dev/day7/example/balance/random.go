package balance

import (
	"errors"
	"math/rand"
)
func init(){
	RegisterBalancer("random",&RandomBalance{})
}
//
type RandomBalance struct{
}
//随机算法函数 
func (p *RandomBalance) DoBalance(insts []*Instance, key ...string)(inst *Instance,err error){
	if len(insts)==0{
		err =errors.New("No instance,没有主机!")
		return
	}
	lens:=len(insts)
	index :=rand.Intn(lens)
	inst=insts[index]
	return
}