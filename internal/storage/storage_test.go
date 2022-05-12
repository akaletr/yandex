package storage

import (
	"testing"
)

func Test_storage_Read(t *testing.T) {
	type fields struct {
		data map[string]string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "#1",
			fields:  fields{map[string]string{}},
			args:    args{"anyString"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "#2",
			fields:  fields{map[string]string{"anotherString": "anyResult"}},
			args:    args{"anotherString"},
			want:    "anyResult",
			wantErr: false,
		},
		{
			name:    "#3",
			fields:  fields{map[string]string{}},
			args:    args{""},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &storage{
				data: tt.fields.data,
			}
			got, err := storage.Read(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storage_Write(t *testing.T) {
	type fields struct {
		data map[string]string
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "#1",
			fields:  fields{map[string]string{}},
			args:    args{"asdkfjkh", "google.com"},
			wantErr: false,
		},
		{
			name:    "#2",
			fields:  fields{map[string]string{}},
			args:    args{"key1", ""},
			wantErr: false,
		},
		{
			name:    "#3",
			fields:  fields{map[string]string{}},
			args:    args{"", "google.com"},
			wantErr: false,
		},
		{
			name:    "#4",
			fields:  fields{map[string]string{"key1": "someURL"}},
			args:    args{"key1", "google.com"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &storage{
				data: tt.fields.data,
			}
			err := storage.Write(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
