package stconfig

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuccesGetConfig(t *testing.T) {

	fileContentStr := `
---
db_settings:
  #  connection_str: "root:root@tcp(db:3306)/news?parseTime=true"
  connection_str: "root:root@tcp(localhost:3308)/news?parseTime=true"
`
	dir, err := ioutil.TempDir("", "storage_config_test")
	require.NoError(t, err)
	defer os.RemoveAll(dir) // clean up
	tmpfn := filepath.Join(dir, "storage.yml")
	err = ioutil.WriteFile(tmpfn, []byte(fileContentStr), 0666)
	require.NoError(t, err)
	os.Clearenv()
	c, err := Get("", dir)
	require.NoError(t, err)
	assert.Equal(t, "root:root@tcp(localhost:3308)/news?parseTime=true", c.DBSettings.ConnectionString)
}

func TestSuccesGetConfigWithFileName(t *testing.T) {
	fileContentStr := `
---
db_settings:
  #  connection_str: "root:root@tcp(db:3306)/news?parseTime=true"
  connection_str: "root:root@tcp(localhost:3308)/news?parseTime=true"
`
	dir, err := ioutil.TempDir("", "storage_config_test")
	require.NoError(t, err)
	defer os.RemoveAll(dir) // clean up
	tmpfn := filepath.Join(dir, "storage.yml")
	err = ioutil.WriteFile(tmpfn, []byte(fileContentStr), 0666)
	require.NoError(t, err)
	os.Clearenv()
	c, err := Get(tmpfn, dir)
	require.NoError(t, err)
	assert.Equal(t, "root:root@tcp(localhost:3308)/news?parseTime=true", c.DBSettings.ConnectionString)
}

func TestEmptyFolderConfig(t *testing.T) {
	dir, err := ioutil.TempDir("", "client_config_test")
	require.NoError(t, err)
	defer os.RemoveAll(dir) // clean up
	_, err = Get(dir)
	assert.Error(t, err)
}
