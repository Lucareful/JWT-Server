package service

import (
	"context"
	"testing"
)

func TestGenerateAuthorizationToken(t *testing.T) {
	type args struct {
		ctx      context.Context
		ClientID string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestGenerateAuthorizationToken", args{context.Background(), "000001"}, "N2JMOTBJOTUTMJFINY0ZYWVJLTG5MJCTMGIXNTMXOTBJM2FM", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateAuthorizationToken(tt.args.ctx, tt.args.ClientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAuthorizationCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateAuthorizationCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateAuthorizationCode(t *testing.T) {
	type args struct {
		ctx      context.Context
		ClientID string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestGenerateAuthorizationToken", args{context.Background(), "000001"}, "N2JMOTBJOTUTMJFINY0ZYWVJLTG5MJCTMGIXNTMXOTBJM2FM", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateAuthorizationCode(tt.args.ctx, tt.args.ClientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAuthorizationCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateAuthorizationCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
