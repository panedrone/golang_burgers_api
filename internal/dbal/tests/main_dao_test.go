package dbal

import (
	"app/internal/dbal"
	"context"
	"os"
	"path"
	"runtime"
	"testing"
)

var ctx context.Context

func TestMain(m *testing.M) {
	{
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Join(path.Dir(filename), "../../..")
		err := os.Chdir(dir)
		if err != nil {
			panic(err)
		}
	}

	err := dbal.OpenDb()
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		_ = dbal.CloseDb()
	}()
	ctx = context.Background()
	code := m.Run()
	// .................... clean up
	os.Exit(code)
}
