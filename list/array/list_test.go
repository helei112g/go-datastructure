package array

import (
	"testing"
)

func TestSqList_ClearList(t *testing.T) {
	l1 := InitList()
	l1.Insert(10, 1)
	l1.Insert(20, 2)
	l1.Insert(30, 3)
	tests := []struct {
		name   string
		fields *SqList
		want   bool
	}{
		{
			name:   "test01",
			fields: l1,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.ClearList(); got != tt.want {
				t.Errorf("ClearList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqList_Delete(t *testing.T) {
	l1 := InitList()
	l1.Insert(11, 1)
	l1.Insert(21, 2)
	l1.Insert(31, 3)
	l1.Insert(51, 4)
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields *SqList
		args   args
		wantE  ElemType
		wantOk bool
	}{
		{
			name:   "test01",
			fields: l1,
			args:   args{i: 4},
			wantE:  51,
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SqList{
				list: tt.fields.list,
				len:  tt.fields.len,
			}
			gotE, gotOk := l.Delete(tt.args.i)
			if gotE != tt.wantE {
				t.Errorf("Delete() gotE = %v, want %v", gotE, tt.wantE)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Delete() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSqList_GetElem(t *testing.T) {
	l1 := InitList()
	l1.Insert(11, 1)
	l1.Insert(21, 2)
	l1.Insert(31, 3)
	l1.Insert(51, 4)
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields *SqList
		args   args
		wantE  ElemType
		wantOk bool
	}{
		{
			name:   "test01",
			fields: l1,
			args:   args{i: 4},
			wantE:  51,
			wantOk: true,
		},
		{
			name:   "test02",
			fields: l1,
			args:   args{i: 10},
			wantE:  0,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotE, gotOk := tt.fields.GetElem(tt.args.i)
			if gotE != tt.wantE {
				t.Errorf("GetElem() gotE = %v, want %v", gotE, tt.wantE)
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetElem() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSqList_GetElemLocation(t *testing.T) {
	l1 := InitList()
	l1.Insert(10, 1)
	l1.Insert(40, 2)
	l1.Insert(20, 3)
	type args struct {
		e ElemType
	}
	tests := []struct {
		name   string
		fields *SqList
		args   args
		wantI  int
		wantOk bool
	}{
		{
			name:   "test01",
			fields: l1,
			args:   args{e: 20},
			wantI:  3,
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, gotOk := tt.fields.GetElemLocation(tt.args.e)
			if gotI != tt.wantI {
				t.Errorf("GetElemLocation() gotI = %v, want %v", gotI, tt.wantI)
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetElemLocation() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSqList_GetLen(t *testing.T) {
	l1 := InitList()
	l2 := InitList()
	l2.Insert(1, 1)
	l2.Insert(2, 2)
	tests := []struct {
		name   string
		fields *SqList
		want   int
	}{
		{
			name:   "test01",
			fields: l1,
			want:   0,
		},
		{
			name:   "test02",
			fields: l2,
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetLen(); got != tt.want {
				t.Errorf("GetLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqList_Insert(t *testing.T) {
	type fields struct {
		list [MaxLen]ElemType
		len  int
	}
	type args struct {
		e ElemType
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test01",
			fields: fields{
				list: [20]ElemType{},
				len:  0,
			},
			args: args{
				e: 10,
				i: 1,
			},
			want: true,
		},
		{
			name: "test02",
			fields: fields{
				list: [20]ElemType{},
				len:  0,
			},
			args: args{
				e: 1,
				i: 2,
			},
			want: true,
		},
	}
	l := InitList()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := l.Insert(tt.args.e, tt.args.i); got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
			t.Log(l.list)
		})
	}
}

func TestSqList_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		fields *SqList
		want   bool
	}{
		{
			name:   "test01",
			fields: InitList(),
			want:   true,
		},
		{
			name:   "test02",
			fields: InitList(),
			want:   false,
		},
	}
	l := InitList()
	for _, tt := range tests {
		if tt.name == "test01" {
			t.Run(tt.name, func(t *testing.T) {
				if got := l.IsEmpty(); got != tt.want {
					t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
				}
			})
		} else {
			l.Insert(1, 1)
			l.Insert(2, 2)
			l.Insert(3, 3)
			t.Run(tt.name, func(t *testing.T) {
				if got := l.IsEmpty(); got != tt.want {
					t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
