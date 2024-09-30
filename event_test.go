package sse

import (
	"testing"
)

func TestEvent_SetID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		e    *Event
		args args
	}{
		{
			name: "testSetID",
			e:    &Event{},
			args: args{
				id: "1111",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.SetID(tt.args.id)
			if tt.e.ID != tt.args.id {
				t.Errorf("SetID() = %v, want %v", tt.e.ID, tt.args.id)
			}
		})
	}
}

func TestEvent_SetData(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		e    *Event
		args args
	}{
		{
			name: "testSetData",
			e:    &Event{},
			args: args{
				data: "Hello, SSE!",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.SetData(tt.args.data)
			if tt.e.Data != tt.args.data {
				t.Errorf("SetID() = %v, want %v", tt.e.Data, tt.args.data)
			}
		})
	}
}
