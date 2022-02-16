package common

//enum.go 是公共文件路径下的 定义枚举值文件

const (
	//错误码
	Code = 0
	//成功
	CodeOk = 1
	//未知错误
	CodeErr = 2
	//参数错误
	CodeInvalidParam = 3
	//参数错误
	CodeInvalidType = 4
)
const (
	//属性的数据类型
	AttributeType = 0
	//string
	AttributeString = 1
	//int8
	AttributeInt8 = 2
	//uint8
	AttributeUint8 = 3
	//int16
	AttributeInt16 = 4
	//uint16
	AttributeUint16 = 5
	//int32
	AttributeInt32 = 6
	//uint32
	AttributeUint32 = 7
	//int64
	AttributeInt64 = 8
	//uint64
	AttributeUint64 = 9
	//float32
	AttributeFloat32 = 10
	//float64
	AttributeFloat64 = 11
	//byte
	AttributeByte = 12
	//bool
	AttributeBool = 13
)
