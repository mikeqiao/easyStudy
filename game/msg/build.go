package msg

//go:generate protoc -I. -I../ -I../baseStruct -I../msg --go_out=. --go_opt=paths=source_relative msg.proto
