package configs

import (
	"reflect"
	"testing"
)

func TestInitEnvVariables(t *testing.T) {
	tests := []struct {
		name    string
		want    *AppConfig
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitEnvVariables()
			if (err != nil) != tt.wantErr {
				t.Errorf("InitEnvVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitEnvVariables() = %v, want %v", got, tt.want)
			}
		})
	}
}
