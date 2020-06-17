package main

import "testing"

func Test_parseURL(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name         string
		args         args
		wantAddr     string
		wantCipher   string
		wantPassword string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAddr, gotCipher, gotPassword, err := parseURL(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAddr != tt.wantAddr {
				t.Errorf("parseURL() gotAddr = %v, want %v", gotAddr, tt.wantAddr)
			}
			if gotCipher != tt.wantCipher {
				t.Errorf("parseURL() gotCipher = %v, want %v", gotCipher, tt.wantCipher)
			}
			if gotPassword != tt.wantPassword {
				t.Errorf("parseURL() gotPassword = %v, want %v", gotPassword, tt.wantPassword)
			}
		})
	}
}
