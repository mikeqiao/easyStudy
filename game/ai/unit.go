package ai

const (
	AIUnit_State   = 0
	AIUnit_Idle    = 1 //等待闲置 idle 状态
	AIUnit_Working = 2 //工作状态
	AIUnit_Stop    = 3 //暂停状态
	AIUnit_Closing = 4 //关闭中状态
	AIUnit_Closed  = 5 //已经关闭状态
)

//一个ai 单元
type AIUnit struct {
	UnitID   int32 //唯一Id
	UnitType int32 //ai单元 类型
	State    int32 //ai单元的 状态
}

//初始化
func (a *AIUnit) Init() {

}

//开启启动
func (a *AIUnit) Start() {

}

//暂停
func (a *AIUnit) Stop() {

}

//关闭
func (a *AIUnit) Close() {

}
