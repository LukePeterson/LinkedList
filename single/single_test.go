package single

import (
	"reflect"
	"testing"
)

func TestLinkedList_append(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
	}
	type args struct {
		newNode *Node
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantHead *Node
		wantTail *Node
	}{
		{
			name:     "Append to an empty list",
			fields:   fields{head: nil, tail: nil},
			args:     args{newNode: &Node{value: 1}},
			wantHead: &Node{value: 1},
			wantTail: &Node{value: 1},
		},
		{
			name:     "Append to a list with one node",
			fields:   fields{head: &Node{value: 1}, tail: &Node{value: 1}},
			args:     args{newNode: &Node{value: 2}},
			wantHead: &Node{value: 1},
			wantTail: &Node{value: 2},
		},
		{
			name:     "Append to a list with multiple nodes",
			fields:   fields{head: &Node{value: 1}, tail: &Node{value: 2}},
			args:     args{newNode: &Node{value: 3}},
			wantHead: &Node{value: 1},
			wantTail: &Node{value: 3},
		},
		{
			name:     "Append a nil node",
			fields:   fields{head: &Node{value: 1}, tail: &Node{value: 1}},
			args:     args{newNode: nil},
			wantHead: &Node{value: 1},
			wantTail: &Node{value: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			list.append(tt.args.newNode)
			if (list.head != nil && list.head.value != tt.wantHead.value) ||
				(list.tail != nil && list.tail.value != tt.wantTail.value) {
				t.Errorf("LinkedList.append() head = %v, tail = %v; want head = %v, tail = %v",
					list.head.value, list.tail.value, tt.wantHead.value, tt.wantTail.value)
			}
		})
	}
}

func TestLinkedList_toSlice(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "Convert an empty list to a slice",
			fields: fields{
				head: nil,
				tail: nil,
			},
			want: []int{},
		},
		{
			name: "Convert a single-element list to a slice",
			fields: fields{
				head: &Node{value: 1},
				tail: &Node{value: 1},
			},
			want: []int{1},
		},
		{
			name: "Convert a multi-element list to a slice",
			fields: fields{
				head: &Node{value: 1, next: &Node{value: 2, next: &Node{value: 3}}},
				tail: &Node{value: 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			if got := list.toSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.toSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_toList(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
	}
	type args struct {
		input []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name:   "Empty input",
			fields: fields{head: nil, tail: nil},
			args:   args{input: []int{}},
			want:   []int{},
		},
		{
			name:   "Single element",
			fields: fields{head: nil, tail: nil},
			args:   args{input: []int{5}},
			want:   []int{5},
		},
		{
			name:   "Multiple elements",
			fields: fields{head: nil, tail: nil},
			args:   args{input: []int{1, 2, 3}},
			want:   []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			list.toList(tt.args.input)
			got := list.toSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.toList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_reverse(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "Empty input",
			fields: fields{
				head: nil,
				tail: nil,
			},
			want: []int{},
		},
		{
			name: "Single element",
			fields: fields{
				head: &Node{value: 1, next: nil},
				tail: nil,
			},
			want: []int{1},
		},
		{
			name: "Multiple elements",
			fields: fields{
				head: &Node{value: 1, next: &Node{value: 2, next: &Node{value: 3, next: nil}}},
				tail: nil,
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			list.reverse()

			got := []int{}
			for node := list.head; node != nil; node = node.next {
				got = append(got, node.value)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
