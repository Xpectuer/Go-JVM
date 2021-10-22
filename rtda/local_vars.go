package rtda

import "math"

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (lv LocalVars) SetInt(index uint, val int32) {
	lv[index].num = val
}

func (lv LocalVars) GetInt(index uint) int32 {
	return lv[index].num
}

func (lv LocalVars) SetFloat(index uint, val float32) {
	lv[index].num = int32(math.Float32bits(val))
}

func (lv LocalVars) GetFloat(index uint) float32 {
	return math.Float32frombits((uint32)(lv[index].num))
}

func (lv LocalVars) SetLong(index uint, val int64) {
	// lower 32 bits
	lv[index].num = int32(val)
	// higher 32 bits
	lv[index+1].num = int32(val >> 32)
}

func (lv LocalVars) GetLong(index uint) int64 {
	low := uint32(lv[index].num)
	high := uint32(lv[index+1].num)
	return int64(high)<<32 | int64(low)
}

// Reuses Long Type's code
func (lv LocalVars) SetDouble(index uint, val float64) {
	lv.SetLong(index, int64(math.Float64bits(val)))
}

func (lv LocalVars) GetDouble(index uint) float64 {
	bits := uint64(lv.GetLong(index))
	return math.Float64frombits(bits)
}

func (lv LocalVars) SetRef(index uint, ref *Object) {
	lv[index].ref = ref
}

func (lv LocalVars) GetRef(index uint) *Object {
	return lv[index].ref
}
