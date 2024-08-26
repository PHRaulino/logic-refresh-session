package config

import (
	"reflect"
	"testing"
	"time"
)

func TestGetRefreshSteps(t *testing.T) {
	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{
			"Carrega as configurações iniciais",
			&Config{
				RefreshTime:   4 * time.Minute,
				SessionTime:   15 * time.Minute,
				QueueFilename: "queue.txt",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRefreshSteps() expect error ? %v, error receive: '%v'", tt.wantErr, err)
			} else if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRefreshSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
