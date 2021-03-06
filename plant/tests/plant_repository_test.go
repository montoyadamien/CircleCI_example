package tests

import (
	"github.com/stretchr/testify/assert"
	"os"
	"plant/repositories"
	"testing"
)

func setUp(){
	repositories.InitTestDatabase()
}

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestFindAll(t *testing.T) {
	var res = repositories.FindAll()
	assert.Equal(t, 2, len(res))
	assert.Equal(t, 1, res[0].Id)
	assert.Equal(t, "Sunflower", res[0].Name)
	assert.Equal(t, 2, res[1].Id)
	assert.Equal(t, "Poppy", res[1].Name)
}
