package dir

import (
	"os"
	"reflect"
	"testing"
)

func TestNewFile(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFile(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name  string
		args  args
		wantF *File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotF := NewFilePath(tt.args.path); !reflect.DeepEqual(gotF, tt.wantF) {
				t.Errorf("NewFilePath() = %v, want %v", gotF, tt.wantF)
			}
		})
	}
}

func TestFile_WithContent(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		f    *File
		args args
		want *File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.WithContent(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.WithContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_WithParent(t *testing.T) {
	type args struct {
		parent *Directory
	}
	tests := []struct {
		name string
		f    *File
		args args
		want *File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.WithParent(tt.args.parent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.WithParent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_FullName(t *testing.T) {
	tests := []struct {
		name string
		f    *File
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.FullName(); got != tt.want {
				t.Errorf("File.FullName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Read(t *testing.T) {
	tests := []struct {
		name    string
		f       *File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.Read(); (err != nil) != tt.wantErr {
				t.Errorf("File.Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_Write(t *testing.T) {
	tests := []struct {
		name    string
		f       *File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.Write(); (err != nil) != tt.wantErr {
				t.Errorf("File.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_WritePerm(t *testing.T) {
	type args struct {
		perm os.FileMode
	}
	tests := []struct {
		name    string
		f       *File
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.WritePerm(tt.args.perm); (err != nil) != tt.wantErr {
				t.Errorf("File.WritePerm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_Parents(t *testing.T) {
	tests := []struct {
		name  string
		f     *File
		wantS []*Directory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := tt.f.Parents(); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("File.Parents() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestFile_Path(t *testing.T) {
	tests := []struct {
		name string
		f    *File
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Path(); got != tt.want {
				t.Errorf("File.Path() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_String(t *testing.T) {
	tests := []struct {
		name string
		f    *File
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.String(); got != tt.want {
				t.Errorf("File.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
