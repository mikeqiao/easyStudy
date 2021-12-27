package attribute

import (
	"fmt"
	"github.com/mikeqiao/easyStudy/game/baseStruct"
	"github.com/mikeqiao/easyStudy/game/common"
)

//创建一个新的avatar结构
// avatarInfo 为传递来的格式实体格式
func MakeNewAvatarStruct(avatarInfo *baseStruct.AvatarStruct) (code int32, errDes error) {
	code = common.CodeOk
	errDes = nil
	if nil == EntityStructManager {
		code = common.CodeErr
		errDes = fmt.Errorf("Nil  EntityStructManager!")
		return
	}

	id := EntityStructManager.GetAvatarIdByName(avatarInfo.Name)
	if 0 != id {
		code = common.CodeInvalidParam
		errDes = fmt.Errorf("invalid  avatarName!")
		return
	}
	//验证结构的合法性
	for k, v := range avatarInfo.Data {
		if nil != v {
			s := EntityStructManager.GetStruct(v.KeyID)
			if nil == s {
				code = common.CodeInvalidParam
				errDes = fmt.Errorf("not have this EntityStruct! key: %v", k)
				return
			}
		} else {
			code = common.CodeErr
			errDes = fmt.Errorf("invalid EntityStruct! key: %v", k)
			return
		}
	}
	//合法的结构加入 列表
	avatarInfo.KeyID = EntityStructManager.GetAvatarId()
	EntityStructManager.AddAvatar(avatarInfo)
	return
}
