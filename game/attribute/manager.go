package attribute

import (
	"github.com/mikeqiao/easyStudy/game/baseStruct"
	"sync"
)

var EntityStructManager *StructManager

func init() {
	EntityStructManager = new(StructManager)
	//读取现有的数据，初始化信息
	EntityStructManager.LoadData()
}

type StructManager struct {
	nowStructId    int32
	nowAvatarId    int32
	structList     sync.Map // map[int32]*msg.EntityStruct
	avatarList     sync.Map //map[int32]*msg.AvatarStruct
	avatarNameList sync.Map //map[string]int32
}

//加载所有定义好的结构数据到 manager
func (s *StructManager) LoadData() {
	s.nowStructId = 1
	s.nowAvatarId = 1
	//从存储加载数据，如果没有数据就不执行，有数据就替换 nowId
	return
}

//创建一个新的结构Id(唯一)
//id 为返回的结构id
func (s *StructManager) GetStructId() (id int32) {
	s.nowStructId += 1
	id = s.nowStructId
	return
}

//创建一个新的结构Id(唯一)
//id 为返回的结构id
func (s *StructManager) GetAvatarId() (id int32) {
	s.nowAvatarId += 1
	id = s.nowAvatarId
	return
}

//根据key获取现有的一个实体结构，并返回指针
//entityId 为实体结构的唯一Key
func (s *StructManager) GetStruct(entityId int32) *baseStruct.EntityStruct {
	v, ok := s.structList.Load(entityId)
	if ok && nil != v {
		d, ok := v.(*baseStruct.EntityStruct)
		if ok {
			return d
		}
	}
	return nil
}

//将一个新的实体结构加入到控制器中
//entityInfo为新的实体结构
func (s *StructManager) AddStruct(entityInfo *baseStruct.EntityStruct) {
	if nil != entityInfo {
		s.structList.Store(entityInfo.KeyID, entityInfo)
	}
}

//将一个现有的实体结构从控制器中删除
//entityId 为实体结构的唯一Key
func (s *StructManager) DelStruct(entityId int32) {
	_, ok := s.structList.Load(entityId)
	if ok {
		s.structList.Delete(entityId)
	}
}

//根据key获取现有的一个实体avatar，并返回指针
//avatarId 为实体结构的唯一Key
func (s *StructManager) GetAvatar(avatarId int32) *baseStruct.AvatarStruct {
	v, ok := s.avatarList.Load(avatarId)
	if ok && nil != v {
		d, ok := v.(*baseStruct.AvatarStruct)
		if ok {
			return d
		}
	}
	return nil
}

//将一个新的实体avatar加入到控制器中
//entityInfo为新的实体结构
func (s *StructManager) AddAvatar(avatarInfo *baseStruct.AvatarStruct) {
	if nil != avatarInfo {
		s.avatarList.Store(avatarInfo.KeyID, avatarInfo)
		s.avatarNameList.Store(avatarInfo.Name, avatarInfo.KeyID)
	}
}

//将一个现有的实体结构从控制器中删除
//entityId 为实体结构的唯一Key
func (s *StructManager) DelAvatar(avatarId int32) {
	v, ok := s.avatarList.Load(avatarId)
	if ok {
		s.avatarList.Delete(avatarId)
	}
	if ok && nil != v {
		d, ok := v.(*baseStruct.AvatarStruct)
		if ok {
			s.avatarNameList.Delete(d.Name)
		}
	}
}

//通过名字查找一个avatar
//name 为实体结构的唯一Key
func (s *StructManager) GetAvatarIdByName(name string) int32 {
	v, ok := s.avatarNameList.Load(name)
	if ok && nil != v {
		d, ok := v.(int32)
		if ok {
			return d
		}
	}
	return 0
}
