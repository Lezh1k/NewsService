package clconfig

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
server_settings:
  echo_addr: ":5000"

nats_settings:
  client_id: "client"
  cluster_id: "test-cluster"
  connect_timeout: "2s"
  ack_timeout: "30s"
  default_ack_prefix: "_STAN.acks"
  discover_prefix: "_STAN.discover"
  max_pub_ack_in_flight: 100
  ping_interval: 5
  ping_max_out: 3
  nats_url: "nats://nats:4222"
`
	dir, err := ioutil.TempDir("", "client_config_test")
	require.NoError(t, err)
	defer os.RemoveAll(dir) // clean up
	tmpfn := filepath.Join(dir, "client.yml")
	err = ioutil.WriteFile(tmpfn, []byte(fileContentStr), 0666)
	require.NoError(t, err)
	os.Clearenv()
	c, err := Get("", dir)
	require.NoError(t, err)
	assert.Equal(t, "nats://nats:4222", c.NATSSettings.NATSURL)
	assert.Equal(t, "client", c.NATSSettings.ClientID)
	assert.Equal(t, ":5000", c.ServerSettings.EchoAddress)
}

func TestSuccesGetConfigWithFileName(t *testing.T) {
	fileContentStr := `
---
nats_settings:
  client_id: "client"
  cluster_id: "test-cluster"
  connect_timeout: "2s"
  ack_timeout: "30s"
  default_ack_prefix: "_STAN.acks"
  discover_prefix: "_STAN.discover"
  max_pub_ack_in_flight: 100
  ping_interval: 5
  ping_max_out: 3
  nats_url: "nats://nats:4222"
`
	dir, err := ioutil.TempDir("", "client_config_test")
	require.NoError(t, err)
	defer os.RemoveAll(dir) // clean up
	tmpfn := filepath.Join(dir, "client.yml")
	err = ioutil.WriteFile(tmpfn, []byte(fileContentStr), 0666)
	require.NoError(t, err)
	os.Clearenv()
	c, err := Get(tmpfn, dir)
	require.NoError(t, err)
	assert.Equal(t, "nats://nats:4222", c.NATSSettings.NATSURL)
	assert.Equal(t, "client", c.NATSSettings.ClientID)
}

func TestEmptyFolderConfig(t *testing.T) {
	dir, err := ioutil.TempDir("", "client_config_test")
	require.NoError(t, err)
	defer os.RemoveAll(dir) // clean up
	_, err = Get(dir)
	assert.Error(t, err)
}
