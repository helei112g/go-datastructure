package trie

import "testing"

func TestTrieNode_Insert(t1 *testing.T) {
	type fields struct {
		data      rune
		children  [SIZE]*TrieNode
		isEndChar bool
	}
	type args struct {
		test string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOk bool
	}{
		{
			name: "test01",
			fields: fields{
				data:      0,
				children:  [26]*TrieNode{},
				isEndChar: false,
			},
			args:   args{test: "her"},
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TrieNode{
				data:      tt.fields.data,
				children:  tt.fields.children,
				isEndChar: tt.fields.isEndChar,
			}
			if gotOk := t.Insert(tt.args.test); gotOk != tt.wantOk {
				t1.Errorf("Insert() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestTrieNode_Find(t1 *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{pattern: "hi"},
			want: false,
		},
		{
			name: "test02",
			args: args{pattern: "her"},
			want: true,
		},
		{
			name: "test03",
			args: args{pattern: "hello"},
			want: true,
		},
		{
			name: "test04",
			args: args{pattern: "lover"},
			want: true,
		},
	}
	t := &TrieNode{}
	t.Insert("hello")
	t.Insert("her")
	t.Insert("lover")
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := t.Find(tt.args.pattern); got != tt.want {
				t1.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
