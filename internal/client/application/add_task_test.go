package application

import (
	"chatgo/internal/client/application/mocks"
	"chatgo/internal/client/infrastructure/storage"
	"github.com/gofrs/uuid"
	"reflect"
	"testing"
)

func TestAddTask_Handle(t *testing.T) {
	type args struct {
		title       string
		description string
	}
	tests := []struct {
		name    string
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		{
			name: "new task",
			args: args{
				title:       "a",
				description: "b",
			},
			want:    uuid.Must(uuid.FromString("00000000-0000-0000-0000-000000000001")),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := storage.NewTaskRepository()
			chatgoMock := mocks.NewChatGoService(t)
			chatgoMock.On("AddTask", "a", "b").Return(
				uuid.Must(uuid.FromString("00000000-0000-0000-0000-000000000001")),
				nil,
			)
			h := NewAddTask(repository, chatgoMock)
			got, err := h.Handle(tt.args.title, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}
