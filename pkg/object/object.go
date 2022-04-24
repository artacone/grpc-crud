package object

import "fmt"

type Object struct {
	Id   uint64
	Data Data
}

type Data struct {
	Ts   int64
	Name string
}

func New(id uint64, ts int64, name string) *Object {
	return &Object{Id: id, Data: Data{Ts: ts, Name: name}}
}

func (obj Object) String() string {
	return fmt.Sprintf("ID: <%v>, TS: <%v>, NAME:<%v>",
		obj.Id, obj.Data.Ts, obj.Data.Name)
}
