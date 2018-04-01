package Graphics

import (
	"reflect"
)

type NotifyEvent func(sender interface{})
type GDispatchObj struct {
	TargetObject interface{}
	DispatchHandler reflect.Value
}

type IInheritedObject interface {
	SubChild(idx int) interface{}
	SubChildCount() int
	Destroy()
	Owner()interface{}
	OwnerChildCount()int
	OwnerChild(idx int)interface{}
	RemoveOwnerChild(obj interface{})
	SetOwner(aowner interface{})
	AddOwnerChild(obj interface{})
}

type GObject struct {
	fSubChilds []interface{} //子对象
	fOwner     interface{}
	fOwnerList []interface{}
	TagData		interface{}
}

func (obj *GObject)SetOwner(aowner interface{})  {
	var targetobj interface{}
	if i := obj.SubChildCount() - 1; i >= 0{
		targetobj = obj.SubChild(i)
	}else{
		targetobj = obj
	}
	if obj.fOwner != nil{
		obj.fOwner.(IInheritedObject).RemoveOwnerChild(targetobj)
	}
	obj.fOwner = aowner
	if aowner != nil{
		aowner.(IInheritedObject).AddOwnerChild(targetobj)
	}
}

func (obj *GObject)AddOwnerChild(childobj interface{})  {
	if obj.fOwnerList == nil{
		obj.fOwnerList =make([]interface{},1)
		obj.fOwnerList[0] = childobj
	}else{
		obj.fOwnerList = append(obj.fOwnerList,childobj)
	}
}

func (obj *GObject)Owner()interface{}  {
	return obj.fOwner
}

func (obj *GObject)RemoveOwnerChild(delobj interface{})  {
	if obj.fOwnerList!= nil{
		delobj.(*GObject).fOwner =nil
		for k,v := range obj.fOwnerList{
			if v == delobj{
				obj.fOwnerList = append(obj.fOwnerList[:k],obj.fOwnerList[k+1:])
			}
		}
	}
}

func (obj *GObject) OwnerChild(idx int) interface{} {
	if obj.fOwnerList == nil {
		return nil
	}
	if idx >= 0 && idx < len(obj.fOwnerList) {
		return obj.fOwnerList[idx]
	}
	return nil
}

func (obj *GObject)OwnerChildCount()int  {
	if obj.fOwner != nil{
		return len(obj.fOwnerList)
	}
	return 0
}

func (obj *GObject)Free()  {
	//执行Destroy过程
	if i := obj.SubChildCount() - 1; i >= 0{
		obj.SubChild(i).(IInheritedObject).Destroy()

	}else{
		obj.Destroy()
	}
}

func (obj *GObject)Destroy()  {
	//释放的过程，后面继承的重写此方法则可
	if obj.fOwnerList != nil {
		for i := 0;i<len(obj.fOwnerList);i++{
			subObj := obj.fOwnerList[i]
			subObj.(IInheritedObject).Destroy()
		}
	}
	if obj.fOwner != nil{
		obj.fOwner.(IInheritedObject).RemoveOwnerChild(obj)
	}
}

func (obj *GObject) SubInit(subObj interface{}) {
	if obj.fSubChilds == nil {
		obj.fSubChilds = make([]interface{}, 0)
	}
	for _, v := range obj.fSubChilds {
		if v == subObj {
			return
		}
	}
	obj.fSubChilds = append(obj.fSubChilds, subObj)
}

func (obj *GObject) SubChild(idx int) interface{} {
	if obj.fSubChilds == nil {
		return nil
	}
	if idx >= 0 && idx < len(obj.fSubChilds) {
		return obj.fSubChilds[idx]
	}
	return nil
}

func (obj *GObject)RealObject()interface{}  {
	if i := obj.SubChildCount() -1;i>=0{
		return obj.SubChild(i)
	}else{
		return obj
	}
}

func (obj *GObject) SubChildCount() int {
	if obj.fSubChilds == nil {
		return 0
	}
	return len(obj.fSubChilds)
}
