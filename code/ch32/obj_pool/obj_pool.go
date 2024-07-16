package object_pool

import (
	"errors"
	"time"
)

type ReusableObj struct {
}

// 对象池结构体
type ObjPool struct {
	bufChan chan *ReusableObj // 用于缓冲可重用对象的通道
}

// 创建对象池的构造函数
func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj) // 创建缓冲通道, 容量为 numOfObj
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{} // 初始化对象池, 填充可重用对象
	}
	return &objPool
}

// 从对象池获取一个对象的方法
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan: // 尝试从缓冲通道获取对象
		return ret, nil
	case <-time.After(timeout): // 超时控制
		return nil, errors.New("time out") // 获取对象超时, 返回错误
	}
}

// 归还一个对象到对象池的方法
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj: // 尝试将对象放回缓冲通道
		return nil
	default: // 如果缓冲通道已满
		return errors.New("overflow") // 返回错误"overflow"
	}
}
