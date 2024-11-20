package services

import "testing"

func TestGetVerify(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Vaild Email",
			args:    args{email: "11130403@me.mcu.edu.tw"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetVerify(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("GetVerify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
