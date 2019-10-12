package dir

import (
	"reflect"
	"testing"
)

func TestCurrentDirectory(t *testing.T) {
	tests := []struct {
		name    string
		wantD   *Directory
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, err := CurrentDirectory()
			if (err != nil) != tt.wantErr {
				t.Errorf("CurrentDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotD, tt.wantD) {
				t.Errorf("CurrentDirectory() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}

func TestNewDirectory(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Directory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDirectory(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirectory_AddFile(t *testing.T) {
	type args struct {
		f *File
	}
	tests := []struct {
		name string
		d    *Directory
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.AddFile(tt.args.f)
		})
	}
}

func TestDirectory_AddDirectory(t *testing.T) {
	type args struct {
		dir *Directory
	}
	tests := []struct {
		name string
		d    *Directory
		args args
		want *Directory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AddDirectory(tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Directory.AddDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirectory_AddDirectoryName(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		d    *Directory
		args args
		want *Directory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AddDirectoryName(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Directory.AddDirectoryName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirectory_Write(t *testing.T) {
	tests := []struct {
		name    string
		d       *Directory
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Write(); (err != nil) != tt.wantErr {
				t.Errorf("Directory.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDirectory_WriteAll(t *testing.T) {
	tests := []struct {
		name    string
		d       *Directory
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.WriteAll(); (err != nil) != tt.wantErr {
				t.Errorf("Directory.WriteAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDirectory_AllChildren(t *testing.T) {
	tests := []struct {
		name  string
		d     *Directory
		wantC []Child
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := tt.d.AllChildren(); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("Directory.AllChildren() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestDirectory_Parents(t *testing.T) {
	tests := []struct {
		name  string
		d     *Directory
		wantS []*Directory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := tt.d.Parents(); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("Directory.Parents() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestDirectory_Path(t *testing.T) {
	tests := []struct {
		name string
		d    *Directory
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Path(); got != tt.want {
				t.Errorf("Directory.Path() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirectory_String(t *testing.T) {
	tests := []struct {
		name string
		d    *Directory
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Directory.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recurseChild(t *testing.T) {
	type args struct {
		c Child
		s *[]Child
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recurseChild(tt.args.c, tt.args.s)
		})
	}
}
