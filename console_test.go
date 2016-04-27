package glog

import (
	glog "."
	"flag"
	"testing"
)

// Test Errorf color
func TestErrorColor(t *testing.T) {
	flag.Set("logtostderr", "true")

	if !flag.Parsed() {
		flag.Parse()
	}

	t.Logf("begin\n")
	glog.Errorf("hello world\n")
	t.Logf("end\n")
}

// Test Warningf color
func TestWarningColor(t *testing.T) {
	flag.Set("logtostderr", "true")

	if !flag.Parsed() {
		flag.Parse()
	}

	t.Logf("begin\n")
	glog.Warningf("hello world\n")
	t.Logf("end\n")
}
