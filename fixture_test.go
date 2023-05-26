package protoc_gen_go_grpc_fixture

import (
	"testing"
)

type TestObject2 struct {
	Field string
}

type TestObject1 struct {
	Obj        *TestObject2
	Obj2       TestObject2
	Arr1       []*TestObject2
	Arr2       []TestObject2
	Circular   *TestObject1
	StrSlice   []string
	IntSlice   []int64
	FloatSlice []float32
	StrArr     [10]string
	IntArr     [10]string
	FloatArr   [10]string
}

func TestNilObj(t *testing.T) {
	obj := &TestObject1{}
	Populate(obj)
	if obj.Obj == nil {
		t.Fatalf("obj.Obj is nil")
	}
	if obj.Obj.Field == "" {
		t.Fatal("obj.Obj.field is empty")
	}
	var ref = obj
	for i := 0; i < MaxDepth; i++ {
		if ref.Circular == nil {
			t.Fatalf("ref.Circular is nil after %d iterations", i)
		}
		ref = obj.Circular
	}
	if ref.Circular != nil {
		t.Fatal("depth check failed")
	}
	if len(obj.StrSlice) != ArraySize {
		t.Fatal("StrSlice array size check failed")
	}
	if len(obj.IntSlice) != ArraySize {
		t.Fatal("IntSlice array size check failed")
	}
	if len(obj.FloatSlice) != ArraySize {
		t.Fatal("FloatSlice array size check failed")
	}
	for i := 0; i < ArraySize; i++ {
		if obj.StrSlice[i] == "" {
			t.Fatal("StrSlice array element check failed")
		}
		if obj.IntSlice[i] == 0 {
			t.Fatal("IntSlice array element check failed")
		}
		if obj.FloatSlice[i] == 0 {
			t.Fatal("FloatSlice array element check failed")
		}
	}
}
