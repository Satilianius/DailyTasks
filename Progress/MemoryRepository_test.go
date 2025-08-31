package Progress

import (
	"DailyTasks/Tasks"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMemoryRepository_AddTask(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
	}
	type args struct {
		task Tasks.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Add valid boolean task",
			fields: fields{
				booleanTaskProgress:  make(map[uuid.UUID]map[time.Time]bool),
				numberTaskProgress:   make(map[uuid.UUID]map[time.Time]float64),
				durationTaskProgress: make(map[uuid.UUID]map[time.Time]time.Duration),
			},
			args: args{
				task: Tasks.NewTask(Tasks.BooleanTask, "one"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
			}
			if err := r.AddTask(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("AddTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryRepository_GetAllProgress(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    PrintableProgress
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			got, got1, err := r.GetAllProgress(tt.args.taskUuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllProgress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllProgress() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetAllProgress() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMemoryRepository_GetBooleanProgressAll(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *BooleanProgress
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			got, got1, err := r.GetBooleanProgressAll(tt.args.taskUuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBooleanProgressAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBooleanProgressAll() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetBooleanProgressAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMemoryRepository_GetBooleanProgressBetweenDates(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
		from     time.Time
		to       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *BooleanProgress
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			got, got1, err := r.GetBooleanProgressBetweenDates(tt.args.taskUuid, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBooleanProgressBetweenDates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBooleanProgressBetweenDates() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetBooleanProgressBetweenDates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMemoryRepository_GetDurationProgressAll(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *DurationProgress
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			got, got1, err := r.GetDurationProgressAll(tt.args.taskUuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDurationProgressAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDurationProgressAll() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetDurationProgressAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMemoryRepository_GetDurationProgressBetweenDates(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
		from     time.Time
		to       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *DurationProgress
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			got, got1, err := r.GetDurationProgressBetweenDates(tt.args.taskUuid, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDurationProgressBetweenDates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDurationProgressBetweenDates() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetDurationProgressBetweenDates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMemoryRepository_GetNumberProgressAll(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NumberProgress
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			got, got1, err := r.GetNumberProgressAll(tt.args.taskUuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumberProgressAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNumberProgressAll() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetNumberProgressAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMemoryRepository_GetNumberProgressBetweenDates(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
		from     time.Time
		to       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NumberProgress
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			got, got1, err := r.GetNumberProgressBetweenDates(tt.args.taskUuid, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumberProgressBetweenDates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNumberProgressBetweenDates() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetNumberProgressBetweenDates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMemoryRepository_GetProgressBetweenDates(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
		from     time.Time
		to       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    PrintableProgress
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			got, got1, err := r.GetProgressBetweenDates(tt.args.taskUuid, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProgressBetweenDates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProgressBetweenDates() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetProgressBetweenDates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMemoryRepository_RemoveBoolean(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
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
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			r.RemoveBoolean(tt.args.taskUuid)
		})
	}
}

func TestMemoryRepository_RemoveDuration(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
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
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			r.RemoveDuration(tt.args.taskUuid)
		})
	}
}

func TestMemoryRepository_RemoveNumber(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
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
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			r.RemoveNumber(tt.args.taskUuid)
		})
	}
}

func TestMemoryRepository_RemoveTaskAndProgress(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			if err := r.RemoveTaskAndProgress(tt.args.taskUuid); (err != nil) != tt.wantErr {
				t.Errorf("RemoveTaskAndProgress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryRepository_UpdateBooleanProgress(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
		date     time.Time
		done     bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			if err := r.UpdateBooleanProgress(tt.args.taskUuid, tt.args.date, tt.args.done); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBooleanProgress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryRepository_UpdateDurationProgress(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
		date     time.Time
		value    time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			if err := r.UpdateDurationProgress(tt.args.taskUuid, tt.args.date, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDurationProgress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryRepository_UpdateNumberProgress(t *testing.T) {
	type fields struct {
		booleanTaskProgress  map[uuid.UUID]map[time.Time]bool
		numberTaskProgress   map[uuid.UUID]map[time.Time]float64
		durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration
		booleanMutex         sync.RWMutex
		numberMutex          sync.RWMutex
		durationMutex        sync.RWMutex
	}
	type args struct {
		taskUuid uuid.UUID
		date     time.Time
		value    float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepository{
				booleanTaskProgress:  tt.fields.booleanTaskProgress,
				numberTaskProgress:   tt.fields.numberTaskProgress,
				durationTaskProgress: tt.fields.durationTaskProgress,
				booleanMutex:         tt.fields.booleanMutex,
				numberMutex:          tt.fields.numberMutex,
				durationMutex:        tt.fields.durationMutex,
			}
			if err := r.UpdateNumberProgress(tt.args.taskUuid, tt.args.date, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("UpdateNumberProgress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewMemoryRepository(t *testing.T) {
	tests := []struct {
		name string
		want *MemoryRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemoryRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemoryRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
