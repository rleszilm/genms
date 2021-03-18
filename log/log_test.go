package log_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/rleszilm/genms/log"
	testLog "github.com/rleszilm/genms/log"
)

func TestLevels(t *testing.T) {
	testcases := []struct {
		desc   string
		level  testLog.Level
		expect []byte
	}{
		{
			desc:   "trace",
			level:  testLog.LvlTrace,
			expect: []byte("[Trace](trace): Trace\n[Debug](trace): Debug\n[Info](trace): Info\n[Warning](trace): Warning\n[Error](trace): Error\n"),
		},
		{
			desc:   "debug",
			level:  testLog.LvlDebug,
			expect: []byte("[Debug](debug): Debug\n[Info](debug): Info\n[Warning](debug): Warning\n[Error](debug): Error\n"),
		},
		{
			desc:   "info",
			level:  testLog.LvlInfo,
			expect: []byte("[Info](info): Info\n[Warning](info): Warning\n[Error](info): Error\n"),
		},
		{
			desc:   "warning",
			level:  testLog.LvlWarning,
			expect: []byte("[Warning](warning): Warning\n[Error](warning): Error\n"),
		},
		{
			desc:   "error",
			level:  testLog.LvlError,
			expect: []byte("[Error](error): Error\n"),
		},
		{
			desc:   "disable",
			level:  testLog.LvlDisable,
			expect: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			log.SetLevel(log.LvlError)
			buffer := &bytes.Buffer{}

			ch := testLog.NewChannel(tc.desc)
			ch.WithFlags(0)
			ch.WithLevel(tc.level)
			ch.WithOutput(buffer)

			ch.Trace("Trace")
			ch.Debug("Debug")
			ch.Info("Info")
			ch.Warning("Warning")
			ch.Error("Error")

			if !reflect.DeepEqual(buffer.Bytes(), tc.expect) {
				t.Errorf("Output is not as expected\nExpect:%s\nActual:%s\n", string(tc.expect), string(buffer.Bytes()))
				t.Errorf("Raw Output\nExpect:%v\nActual:%v\n", tc.expect, buffer.Bytes())
			}
		})
	}
}

func TestGlobalLevels(t *testing.T) {
	testcases := []struct {
		desc   string
		level  testLog.Level
		expect []byte
	}{
		{
			desc:   "trace",
			level:  testLog.LvlTrace,
			expect: []byte("[Trace](trace): Trace\n[Debug](trace): Debug\n[Info](trace): Info\n[Warning](trace): Warning\n[Error](trace): Error\n"),
		},
		{
			desc:   "debug",
			level:  testLog.LvlDebug,
			expect: []byte("[Debug](debug): Debug\n[Info](debug): Info\n[Warning](debug): Warning\n[Error](debug): Error\n"),
		},
		{
			desc:   "info",
			level:  testLog.LvlInfo,
			expect: []byte("[Info](info): Info\n[Warning](info): Warning\n[Error](info): Error\n"),
		},
		{
			desc:   "warning",
			level:  testLog.LvlWarning,
			expect: []byte("[Info](warning): Info\n[Warning](warning): Warning\n[Error](warning): Error\n"),
		},
		{
			desc:   "error",
			level:  testLog.LvlError,
			expect: []byte("[Info](error): Info\n[Warning](error): Warning\n[Error](error): Error\n"),
		},
		{
			desc:   "disable",
			level:  testLog.LvlDisable,
			expect: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			log.SetLevel(tc.level)
			buffer := &bytes.Buffer{}

			ch := testLog.NewChannel(tc.desc)
			ch.WithFlags(0)
			ch.WithLevel(log.LvlInfo)
			ch.WithOutput(buffer)

			ch.Trace("Trace")
			ch.Debug("Debug")
			ch.Info("Info")
			ch.Warning("Warning")
			ch.Error("Error")

			if !reflect.DeepEqual(buffer.Bytes(), tc.expect) {
				t.Errorf("Output is not as expected\nExpect:%s\nActual:%s\n", string(tc.expect), string(buffer.Bytes()))
				t.Errorf("Raw Output\nExpect:%v\nActual:%v\n", tc.expect, buffer.Bytes())
			}
		})
	}
}
