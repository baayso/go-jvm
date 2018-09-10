package runtimedataarea

import "math"

// 局部变量表
type LocalVars []Slot

// 创建LocalVars实例
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}

	return nil
}

// int
// int变量直接存取即可

func (lv LocalVars) SetInt(index uint, val int32) {
	lv[index].num = val
}

func (lv LocalVars) GetInt(index uint) int32 {
	return lv[index].num
}

// float
// float变量可以先转成int类型，然后按int变量来处理

func (lv LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)

	lv.SetInt(index, int32(bits))
}

func (lv LocalVars) GetFloat(index uint) float32 {
	bits := uint32(lv.GetInt(index))

	return math.Float32frombits(bits)
}

// long
// long变量则需要拆成两个int变量

func (lv LocalVars) SetLong(index uint, val int64) {
	lv.SetInt(index, int32(val))
	lv.SetInt(index+1, int32(val>>32))
}

func (lv LocalVars) GetLong(index uint) int64 {
	low := uint32(lv.GetInt(index))
	high := uint32(lv.GetInt(index + 1))

	return int64(high)<<32 | int64(low)
}

// double
// double变量可以先转成long类型，然后按照long变量来处理

func (lv LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)

	lv.SetLong(index, int64(bits))
}

func (lv LocalVars) GetDouble(index uint) float64 {
	bits := uint64(lv.GetLong(index))

	return math.Float64frombits(bits)
}

// 引用类型
// 引用类型的值直接存取即可

func (lv LocalVars) SetRef(index uint, ref *Object) {
	lv[index].ref = ref
}

func (lv LocalVars) GetRef(index uint) *Object {
	return lv[index].ref
}
