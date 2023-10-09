package todo

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/domain"
	mockPackageTodo "github.com/ArMo-Team/GoLangPractices_todoList/mocks/internal_/app/todo"
)

func TestNewTodoService(t *testing.T) {
	type args struct {
		repository TodoRepository
	}
	mockRepository := mockPackageTodo.NewTodoRepository(t)
	tests := []struct {
		name string
		args args
		want TodoService
	}{
		{
			name: "NewTodoService_Success",
			args: args{repository: mockRepository},
			want: &todoService{Repository: mockRepository},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTodoService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodoService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoService_GetTodos(t *testing.T) {
	type fields struct {
		Repository TodoRepository
	}
	type testStruct struct {
		name    string
		fields  fields
		want    []*domain.Todo
		wantErr bool
	}
	tests := []testStruct{}
	{
		GetTodos_Error := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			mockRepository.On("GetAll").Return(nil, errors.New("dummy error"))

			return testStruct{
				name:    "GetTodos_Error",
				fields:  fields{Repository: mockRepository},
				want:    nil,
				wantErr: true,
			}
		}

		GetTodos_Success := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			want := []*domain.Todo{
				{ID: 1, Title: "dummy title", Description: "dummy description", Completed: true},
				{ID: 2, Title: "dummy title 2", Description: "dummy description 2", Completed: false},
			}
			mockRepository.On("GetAll").Return(want, nil)

			return testStruct{
				name:    "GetTodos_Success",
				fields:  fields{Repository: mockRepository},
				want:    want,
				wantErr: false,
			}
		}

		tests = append(tests, GetTodos_Error())
		tests = append(tests, GetTodos_Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &todoService{
				Repository: tt.fields.Repository,
			}
			got, err := s.GetTodos()
			if (err != nil) != tt.wantErr {
				t.Errorf("todoService.GetTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoService.GetTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoService_GetTodoByID(t *testing.T) {
	type fields struct {
		Repository TodoRepository
	}
	type args struct {
		id string
	}
	type testStruct struct {
		name    string
		fields  fields
		args    args
		want    *domain.Todo
		wantErr bool
	}
	tests := []testStruct{}
	{
		GetTodoByID_Error_GetByID := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			id := "23"
			mockRepository.On("GetByID", id).Return(nil, errors.New("dummy error"))

			return testStruct{
				name:    "GetTodoByID_Error_GetByID",
				fields:  fields{Repository: mockRepository},
				args:    args{id: id},
				want:    nil,
				wantErr: true,
			}
		}
		GetTodoByID_Error_TodoNotFound := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			id := "43"
			mockRepository.On("GetByID", id).Return(nil, nil)

			return testStruct{
				name:    "GetTodoByID_Error_TodoNotFound",
				fields:  fields{Repository: mockRepository},
				args:    args{id: id},
				want:    nil,
				wantErr: false,
			}
		}
		GetTodoByID_Success := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			id := "43"
			want := &domain.Todo{ID: 23, Title: "dummy title"}
			mockRepository.On("GetByID", id).Return(want, nil)

			return testStruct{
				name:    "GetTodoByID_Success",
				fields:  fields{Repository: mockRepository},
				args:    args{id: id},
				want:    want,
				wantErr: false,
			}
		}

		tests = append(tests, GetTodoByID_Error_GetByID())
		tests = append(tests, GetTodoByID_Error_TodoNotFound())
		tests = append(tests, GetTodoByID_Success())
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &todoService{
				Repository: tt.fields.Repository,
			}
			got, err := s.GetTodoByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoService.GetTodoByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoService.GetTodoByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoService_CreateTodo(t *testing.T) {
	type fields struct {
		Repository TodoRepository
	}
	type args struct {
		todo *domain.Todo
	}
	type testStruct struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	tests := []testStruct{}
	{
		CreateTodo_Error := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			todo := &domain.Todo{}
			mockRepository.On("Create", todo).Return(errors.New("dummy error"))

			return testStruct{
				name:    "CreateTodo_Error",
				fields:  fields{Repository: mockRepository},
				args:    args{todo: todo},
				wantErr: true,
			}
		}

		CreateTodo_Success := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			todo := &domain.Todo{}
			mockRepository.On("Create", todo).Return(nil)

			return testStruct{
				name:    "CreateTodo_Success",
				fields:  fields{Repository: mockRepository},
				args:    args{todo: todo},
				wantErr: false,
			}
		}

		tests = append(tests, CreateTodo_Error())
		tests = append(tests, CreateTodo_Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &todoService{
				Repository: tt.fields.Repository,
			}
			if err := s.CreateTodo(tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("todoService.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_todoService_UpdateTodo(t *testing.T) {
	type fields struct {
		Repository TodoRepository
	}
	type args struct {
		todo *domain.Todo
	}
	type testStruct struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	tests := []testStruct{}
	{
		UpdateTodo_Error := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			todo := &domain.Todo{}
			mockRepository.On("Update", todo).Return(errors.New("dummy error"))

			return testStruct{
				name:    "UpdateTodo_Error",
				fields:  fields{Repository: mockRepository},
				args:    args{todo: todo},
				wantErr: true,
			}
		}

		UpdateTodo_Success := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			todo := &domain.Todo{}
			mockRepository.On("Update", todo).Return(nil)

			return testStruct{
				name:    "UpdateTodo_Success",
				fields:  fields{Repository: mockRepository},
				args:    args{todo: todo},
				wantErr: false,
			}
		}

		tests = append(tests, UpdateTodo_Error())
		tests = append(tests, UpdateTodo_Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &todoService{
				Repository: tt.fields.Repository,
			}
			if err := s.UpdateTodo(tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("todoService.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_todoService_UpdateTodoCompleted(t *testing.T) {
	type fields struct {
		Repository TodoRepository
	}
	type args struct {
		todoID    string
		completed bool
	}
	type testStruct struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	tests := []testStruct{}
	{
		UpdateTodoCompleted_Error_GetByID := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			todoID := "231"
			mockRepository.On("GetByID", todoID).Return(nil, errors.New("dummy error"))

			return testStruct{
				name:    "UpdateTodoCompleted_Error_GetByID",
				fields:  fields{Repository: mockRepository},
				args:    args{todoID: todoID, completed: false},
				wantErr: true,
			}
		}

		UpdateTodoCompleted_Error_Update := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			todoID := "231"
			todo := &domain.Todo{}
			mockRepository.On("GetByID", todoID).Return(todo, nil)

			isCompleted := false
			todo.Completed = isCompleted
			mockRepository.On("Update", todo).Return(errors.New("dummy error"))

			return testStruct{
				name:    "UpdateTodoCompleted_Error_Update",
				fields:  fields{Repository: mockRepository},
				args:    args{todoID: todoID, completed: isCompleted},
				wantErr: true,
			}
		}

		UpdateTodoCompleted_Success := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			todoID := "231"
			todo := &domain.Todo{}
			mockRepository.On("GetByID", todoID).Return(todo, nil)

			isCompleted := false
			todo.Completed = isCompleted
			mockRepository.On("Update", todo).Return(nil)

			return testStruct{
				name:    "UpdateTodoCompleted_Success",
				fields:  fields{Repository: mockRepository},
				args:    args{todoID: todoID, completed: isCompleted},
				wantErr: false,
			}
		}

		tests = append(tests, UpdateTodoCompleted_Error_GetByID())
		tests = append(tests, UpdateTodoCompleted_Error_Update())
		tests = append(tests, UpdateTodoCompleted_Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &todoService{
				Repository: tt.fields.Repository,
			}
			if err := s.UpdateTodoCompleted(tt.args.todoID, tt.args.completed); (err != nil) != tt.wantErr {
				t.Errorf("todoService.UpdateTodoCompleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_todoService_DeleteTodo(t *testing.T) {
	type fields struct {
		Repository TodoRepository
	}
	type args struct {
		id string
	}
	type testStruct struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	tests := []testStruct{}
	{
		DeleteTodo_Error := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			id := "234"
			mockRepository.On("Delete", id).Return(errors.New("dummy error"))

			return testStruct{
				name:    "DeleteTodo_Error",
				fields:  fields{Repository: mockRepository},
				args:    args{id: id},
				wantErr: true,
			}
		}

		DeleteTodo_Success := func() testStruct {
			mockRepository := mockPackageTodo.NewTodoRepository(t)
			id := "234"
			mockRepository.On("Delete", id).Return(nil)

			return testStruct{
				name:    "DeleteTodo_Success",
				fields:  fields{Repository: mockRepository},
				args:    args{id: id},
				wantErr: false,
			}
		}

		tests = append(tests, DeleteTodo_Error())
		tests = append(tests, DeleteTodo_Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &todoService{
				Repository: tt.fields.Repository,
			}
			if err := s.DeleteTodo(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("todoService.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
