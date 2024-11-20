package services

import (
	"testing"
)

func TestGenrateFmailyQRCode(t *testing.T) {
	type args struct {
		familyId uint
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []byte
	}{
		{
			name:  "Generate QR Code for family ID 1",
			args:  args{familyId: 1},
			want:  "https://example.com/qrcode?familyId=1",
			want1: []byte{0x12, 0x34, 0x56, 0x78},
		},
		{
			name:  "Generate QR Code for family ID 42",
			args:  args{familyId: 42},
			want:  "https://example.com/qrcode?familyId=42",
			want1: []byte{0x12, 0x34, 0x56, 0x78},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GenrateFmailyQRCode(tt.args.familyId)
			t.Log(got)
			t.Log(got1)
		})
	}
}
