package config

import (
	"github.com/mediaexchange-io/mx-assert"
	"testing"
)

type Postgres struct {
	Username string `json:"username" yaml:"username" env:"POSTGRES_USERNAME"`
	Password string `json:"password" yaml:"password" env:"POSTGRES_PASSWORD"`
	Hostname string `json:"hostname" yaml:"hostname" env:"POSTGRES_HOSTNAME"`
	Database string `json:"database" yaml:"database" env:"POSTGRES_DATABASE"`
}

type Server struct {
	Host string `json:"host" yaml:"host" env:"SERVER_HOST"`
	Port int    `json:"port" yaml:"port" env:"SERVER_PORT"`
	Ssl  bool   `json:"ssl"  yaml:"ssl"  env:"SERVER_SSL"`
}

type TestConfig struct {
	Postgres Postgres `json:"postgres" yaml:"postgres"`
	Server   Server   `json:"server"   yaml:"server"`
	Version  float64  `json:"version"  yaml:"version"  env:"VERSION"`
	Unsigned uint64   `json:"unsigned" yaml:"unsigned" env:"UNSIGNED"`
}

func TestRead_WithBadFile(t *testing.T) {
	var c TestConfig
	err := WithFile("bad-file-name", &c)
	if err == nil {
		t.Errorf("WithFile should have failed with invalid file")
	}
}

func TestRead_WithJsonFile(t *testing.T) {
	var c TestConfig
	err := WithFile("testdata/appconfig_test.json", &c)

	Assert := assert.With(t)
	Assert.That(err).IsOk()
	Assert.That(c.Postgres.Username).IsEqualTo("pg_user_json")
	Assert.That(c.Postgres.Password).IsEqualTo("pg_pass_json")
	Assert.That(c.Postgres.Hostname).IsEqualTo("pg_host_json")
	Assert.That(c.Postgres.Database).IsEqualTo("pg_db_json")
	Assert.That(c.Server.Host).IsEqualTo("host_json")
	Assert.That(c.Server.Port).IsEqualTo(8000)
	Assert.That(c.Server.Ssl).IsTrue()
	Assert.That(c.Version).IsEqualTo(3.14159)
	Assert.That(c.Unsigned).IsEqualTo(uint64(65536))
}

func TestRead_WithJsonFileAndEnvOverride(t *testing.T) {
	t.Setenv("SERVER_PORT", "9999")

	var c TestConfig
	err := WithFile("testdata/appconfig_test.json", &c)

	Assert := assert.With(t)
	Assert.That(err).IsOk()
	Assert.That(c.Postgres.Username).IsEqualTo("pg_user_json")
	Assert.That(c.Postgres.Password).IsEqualTo("pg_pass_json")
	Assert.That(c.Postgres.Hostname).IsEqualTo("pg_host_json")
	Assert.That(c.Postgres.Database).IsEqualTo("pg_db_json")
	Assert.That(c.Server.Host).IsEqualTo("host_json")
	Assert.That(c.Server.Port).IsEqualTo(9999)
	Assert.That(c.Server.Ssl).IsTrue()
	Assert.That(c.Version).IsEqualTo(3.14159)
	Assert.That(c.Unsigned).IsEqualTo(uint64(65536))
}

func TestRead_WithYamlFile(t *testing.T) {
	var c TestConfig
	err := WithFile("testdata/appconfig_test.yaml", &c)

	Assert := assert.With(t)
	Assert.That(err).IsOk()
	Assert.That(c.Postgres.Username).IsEqualTo("pg_user_yaml")
	Assert.That(c.Postgres.Password).IsEqualTo("pg_pass_yaml")
	Assert.That(c.Postgres.Hostname).IsEqualTo("pg_host_yaml")
	Assert.That(c.Postgres.Database).IsEqualTo("pg_db_yaml")
	Assert.That(c.Server.Host).IsEqualTo("host_yaml")
	Assert.That(c.Server.Port).IsEqualTo(8080)
	Assert.That(c.Server.Ssl).IsTrue()
	Assert.That(c.Version).IsEqualTo(3.14159)
	Assert.That(c.Unsigned).IsEqualTo(uint64(65536))
}

func TestRead_WithYamlFileAndEnvOverride(t *testing.T) {
	t.Setenv("SERVER_PORT", "9999")

	var c TestConfig
	err := WithFile("testdata/appconfig_test.yaml", &c)

	Assert := assert.With(t)
	Assert.That(err).IsOk()
	Assert.That(c.Postgres.Username).IsEqualTo("pg_user_yaml")
	Assert.That(c.Postgres.Password).IsEqualTo("pg_pass_yaml")
	Assert.That(c.Postgres.Hostname).IsEqualTo("pg_host_yaml")
	Assert.That(c.Postgres.Database).IsEqualTo("pg_db_yaml")
	Assert.That(c.Server.Host).IsEqualTo("host_yaml")
	Assert.That(c.Server.Port).IsEqualTo(9999)
	Assert.That(c.Server.Ssl).IsTrue()
	Assert.That(c.Version).IsEqualTo(3.14159)
	Assert.That(c.Unsigned).IsEqualTo(uint64(65536))
}

func TestRead_WithEnvOnly(t *testing.T) {
	t.Setenv("POSTGRES_USERNAME", "pg_user_env")
	t.Setenv("POSTGRES_PASSWORD", "pg_pass_env")
	t.Setenv("POSTGRES_HOSTNAME", "pg_host_env")
	t.Setenv("POSTGRES_DATABASE", "pg_db_env")
	t.Setenv("SERVER_HOST", "host_env")
	t.Setenv("SERVER_PORT", "8888")
	t.Setenv("SERVER_SSL", "false")
	t.Setenv("VERSION", "3.14159")
	t.Setenv("UNSIGNED", "65536")

	var c TestConfig
	err := WithFile("", &c)

	Assert := assert.With(t)
	Assert.That(err).IsOk()
	Assert.That(c.Postgres.Username).IsEqualTo("pg_user_env")
	Assert.That(c.Postgres.Password).IsEqualTo("pg_pass_env")
	Assert.That(c.Postgres.Hostname).IsEqualTo("pg_host_env")
	Assert.That(c.Postgres.Database).IsEqualTo("pg_db_env")
	Assert.That(c.Server.Host).IsEqualTo("host_env")
	Assert.That(c.Server.Port).IsEqualTo(8888)
	Assert.That(c.Server.Ssl).IsFalse()
	Assert.That(c.Version).IsEqualTo(3.14159)
	Assert.That(c.Unsigned).IsEqualTo(uint(65536))
}
