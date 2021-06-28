package diffcopy

import (
	"reflect"
	"testing"
)

func TestFindWaitingFiles(t *testing.T) {
	var tests = []struct {
		name string
		src  string
		dest string
		want []string
	}{
		{
			"test0",
			"../../test/src",
			"../../test/dest",
			[]string{"../../test/src/dir0/file1", "../../test/src/dir0/file2"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindWaitingFiles(tt.src, tt.dest)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("given(%s, %s): expected %s, actual %s", tt.src, tt.dest, tt.want, got)
			}
		})
	}
}
