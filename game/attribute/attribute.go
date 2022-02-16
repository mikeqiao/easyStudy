package attribute

import (
	"github.com/mikeqiao/easyStudy/game/baseStruct"
	"github.com/mikeqiao/easyStudy/game/common"
)

//添加一种新的属性类型
//a 为属性的结构参数
func AddNewAttribute(a *baseStruct.Attribute) (code int32) {
	code = common.CodeOk
	return
}

//删除一种存在的属性类型
//id 为属性的唯一Id
func DelAttribute(id int32) (code int32) {
	code = common.CodeOk
	return
}

//修改一种存在的属性类型
//a 为属性的结构参数
func SetAttribute(a *baseStruct.Attribute) (code int32) {
	code = common.CodeOk
	return
}
