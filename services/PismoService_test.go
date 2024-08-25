package services

import (
	"context"
	"github.com/sudipidus/pismo-test/storage"
	"testing"
)

func TestPismoServiceImpl_CreateAccount(t *testing.T) {
	type fields struct {
		storage storage.Storage
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PismoServiceImpl{
				storage: tt.fields.storage,
			}
			got, err := s.CreateAccount(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}
