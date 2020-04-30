package files

import "testing"

func TestWalkFiles(t *testing.T) {
	// create file structure

	type args struct {
		root string
		fn   WalkerFunc
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WalkFiles(tt.args.root, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("WalkFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}