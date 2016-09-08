package Graphics

import "reflect"

type GDispatchObj struct {
	TargetObject interface{}
	DispatchHandler reflect.Value
}

type IInheritedObject interface {
	SubChild(idx int) interface{}
	SubChildCount() int
	SubInit(subObj interface{})
	Destroy()
}

type GObject struct {
	fSubChilds []interface{} //子对象
	fRealFreeDispatch *GDispatchObj
}

func (obj *GObject)Free()  {
	//执行Destroy过程
	if obj.fRealFreeDispatch == nil{
		for i := obj.SubChildCount() - 1; i >= 0; i-- {
			subObj := obj.SubChild(i)
			pType := reflect.TypeOf(subObj)
			if mnd, ok := pType.MethodByName("Destroy"); ok {
				pType = mnd.Func.Type()
				if pType.NumIn() == 1 && pType.NumOut() == 0 {
					obj.fRealFreeDispatch = new(GDispatchObj)
					obj.fRealFreeDispatch.DispatchHandler = mnd.Func
					obj.fRealFreeDispatch.TargetObject = subObj
					break
				}
			}
		}
	}
	if obj.fRealFreeDispatch!=nil{
		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(obj.fRealFreeDispatch.TargetObject)
		obj.fRealFreeDispatch.DispatchHandler.Call(in)
	}
}

func (obj *GObject)Destroy()  {
	//释放的过程，后面继承的重写此方法则可
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

func (obj *GObject) SubChildCount() int {
	if obj.fSubChilds == nil {
		return 0
	}
	return len(obj.fSubChilds)
}
