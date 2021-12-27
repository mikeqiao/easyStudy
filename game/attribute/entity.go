package attribute

import (
	"fmt"
	"github.com/mikeqiao/easyStudy/game/baseStruct"
	"github.com/mikeqiao/easyStudy/game/common"
)

//创建一个新的实体结构
// entityInfo 为传递来的格式实体格式
func MakeNewEntityStruct(entityInfo *baseStruct.EntityStruct) (code int32, errDes error) {
	code = common.CodeOk
	errDes = nil
	if nil == EntityStructManager {
		code = common.CodeErr
		errDes = fmt.Errorf("Nil  EntityStructManager!")
		return
	}
	nameList := make(map[string]int32)
	//验证结构的合法性
	for k, v := range entityInfo.Data {
		if nil != v {
			if "" == v.Name {
				code = common.CodeInvalidParam
				errDes = fmt.Errorf("nil attributeName! key: %v", k)
				return
			}
			if _, ok := nameList[v.Name]; ok {
				code = common.CodeInvalidParam
				errDes = fmt.Errorf("repeat attributeName! key: %v", k)
				return
			}

			ok := CheckAttributeType(v.AttributeType)
			if !ok {
				code = common.CodeInvalidParam
				errDes = fmt.Errorf("invalid attributeType! key: %v, type: %v", k, v.AttributeType)
				return
			}
			nameList[v.Name] = k
		} else {
			code = common.CodeErr
			errDes = fmt.Errorf("invalid attribute! key: %v", k)
			return
		}
	}

	//合法的结构加入 列表
	entityInfo.KeyID = EntityStructManager.GetStructId()
	EntityStructManager.AddStruct(entityInfo)
	return
}

func CheckAttributeType(attributeType int32) bool {
	switch attributeType {
	//string
	case common.AttributeString:
	//int8
	case common.AttributeInt8:
	//uint8
	case common.AttributeUint8:
	//int16
	case common.AttributeInt16:
	//uint16
	case common.AttributeUint16:
	//int32
	case common.AttributeInt32:
	//uint32
	case common.AttributeUint32:
	//int64
	case common.AttributeInt64:
	//uint64
	case common.AttributeUint64:
	//float32
	case common.AttributeFloat32:
	//float64
	case common.AttributeFloat64:
	//byte
	case common.AttributeByte:
	//bool
	case common.AttributeBool:

	default:
		return false
	}

	return true
}
