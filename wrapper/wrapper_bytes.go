package wrp

import (
	cnv "github.com/fcfcqloow/go-advance/convert"
)

func NewResource(byts []byte) *Resource {
	ins := AsResource(byts)
	return &ins
}
func AsResource(byts []byte) Resource {
	return byts
}
func EncodeJson(stc interface{}) (Resource, error) {
	byts, err := cnv.MarshalJson(stc)
	return AsResource(byts), err
}

func (self Resource) AsWrap() *Resource {
	return &self
}

func (self *Resource) AsString() *String {
	return MakeString(string(self.AsPrimitive()))
}
func (self *Resource) AsPrimitive() []byte {
	return *self
}
func (self *Resource) Copy() (*Resource, error) {
	var tmp interface{}
	err := self.DecodeJson(&tmp)
	if err != nil {
		return nil, err
	}
	ins, err := EncodeJson(tmp)
	if err != nil {
		return nil, err
	}
	return &ins, err
}
func (self *Resource) DecodeJson(stc interface{}) error {
	return cnv.UnMarshalJson(self.AsPrimitive(), stc)
}

func (self *Resource) IndentJson() (string, error) {
	return cnv.IndentJson(self.AsPrimitive())
}
