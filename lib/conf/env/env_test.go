package env

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestDefaultString(t *testing.T) {
	v := DefaultString("a", "test")
	if v != "test" {
		t.Fatal("v must be test")
	}
	if err := os.Setenv("a", "test1"); err != nil {
		t.Fatal(err)
	}
	v = DefaultString("a", "test")
	if v != "test1" {
		t.Fatal("v must be test1")
	}
}

func TestEnv(t *testing.T) {
	tests := []struct {
		flag string
		env  string
		def  string
		val  *string
	}{
		{
			"region",
			"REGION",
			_region,
			&System.Region,
		},
		{
			"zone",
			"ZONE",
			_zone,
			&System.Zone,
		},
		{
			"deploy.env",
			"DEPLOY_ENV",
			_deployEnv,
			&System.DeployEnv,
		},
		{
			"appid",
			"APP_ID",
			"",
			&System.AppID,
		},
		{
			"deploy.color",
			"DEPLOY_COLOR",
			"",
			&System.Color,
		},
	}
	for _, test := range tests {
		// flag set value
		t.Run(fmt.Sprintf("%s: flag set", test.env), func(t *testing.T) {
			fs := flag.NewFlagSet("", flag.ContinueOnError)
			addFlag(fs)
			err := fs.Parse([]string{fmt.Sprintf("-%s=%s", test.flag, "test")})
			if err != nil {
				t.Fatal(err)
			}
			if *test.val != "test" {
				t.Fatal("val must be test")
			}
		})
		// flag not set, env set
		t.Run(fmt.Sprintf("%s: flag not set, env set", test.env), func(t *testing.T) {
			*test.val = ""
			os.Setenv(test.env, "test2")
			fs := flag.NewFlagSet("", flag.ContinueOnError)
			addFlag(fs)
			err := fs.Parse([]string{})
			if err != nil {
				t.Fatal(err)
			}
			if *test.val != "test2" {
				t.Fatal("val must be test")
			}
		})
		// flag not set, env not set
		t.Run(fmt.Sprintf("%s: flag not set, env not set", test.env), func(t *testing.T) {
			*test.val = ""
			os.Setenv(test.env, "")
			fs := flag.NewFlagSet("", flag.ContinueOnError)
			addFlag(fs)
			err := fs.Parse([]string{})
			if err != nil {
				t.Fatal(err)
			}
			if *test.val != test.def {
				t.Fatal("val must be test")
			}
		})
	}
}
