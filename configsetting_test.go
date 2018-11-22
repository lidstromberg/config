package config

import (
	"context"
	"testing"
)

func Test_ConfigSetGet(t *testing.T) {
	ctx := context.Background()

	var (
		testKey   = "testkey"
		testValue = "testvalue"
		bc        ConfigSetting
	)

	//get a config map
	bc = NewConfig(ctx)

	//set a value
	bc.SetConfigValue(ctx, testKey, testValue)

	result := bc.GetConfigValue(ctx, testKey)

	if result != testValue {
		t.Fatalf("did not locate %s", testKey)
	}
}
func Test_ConfigMapLoad(t *testing.T) {
	ctx := context.Background()

	var bc ConfigSetting

	//get a config map
	bc = NewConfig(ctx)

	mp := make(map[string]string)

	mp["one"] = "one"
	mp["two"] = "two"
	mp["three"] = "three"
	mp["four"] = "four"
	mp["five"] = "five"

	//set a value
	bc.LoadConfigMap(ctx, mp)

	result := bc.GetConfigValue(ctx, "five")

	if result != "five" {
		t.Fatal("did not locate five")
	}
}
func Test_ConfigDump(t *testing.T) {
	ctx := context.Background()

	var bc ConfigSetting

	//get a config map
	bc = NewConfig(ctx)

	mp := make(map[string]string)

	mp["one"] = "one"
	mp["two"] = "two"
	mp["three"] = "three"
	mp["four"] = "four"
	mp["five"] = "five"

	//set a value
	bc.LoadConfigMap(ctx, mp)

	//dump contents
	bc.DumpConfigMap(ctx)
}
