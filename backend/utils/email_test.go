package utils

import "testing"

func TestSendEmail(t *testing.T) {
	type args struct {
		email      string
		verifycode string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "vaildEmail",
			args:    args{email: "11130403@me.mcu.edu.tw", verifycode: "123456"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendEmail(tt.args.email, tt.args.verifycode); (err != nil) != tt.wantErr {
				t.Errorf("SendEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
