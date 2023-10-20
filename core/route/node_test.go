package route

import (
	"reflect"
	"testing"
)

func Test_node_insert(t *testing.T) {
	type fields struct {
		component string
		wildChild bool
		children  []*node
		route     *route
	}
	type args struct {
		components []string
		route      *route
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				component: tt.fields.component,
				wildChild: tt.fields.wildChild,
				children:  tt.fields.children,
				route:     tt.fields.route,
			}
			n.insert(tt.args.components, tt.args.route)
		})
	}
}

func Test_node_search(t *testing.T) {
	type fields struct {
		component string
		wildChild bool
		children  []*node
		route     *route
	}
	type args struct {
		components []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *route
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				component: tt.fields.component,
				wildChild: tt.fields.wildChild,
				children:  tt.fields.children,
				route:     tt.fields.route,
			}
			if got := n.search(tt.args.components); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
