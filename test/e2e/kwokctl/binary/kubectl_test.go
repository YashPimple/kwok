/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package binary_test

import (
	"runtime"
	"testing"

	"sigs.k8s.io/e2e-framework/pkg/envconf"

	"sigs.k8s.io/kwok/test/e2e"
)

func TestHack(t *testing.T) {
	f0 := e2e.CaseHack(kwokctlPath, clusterName, envconf.RandomName("node", 16)).
		Feature()
	testEnv.Test(t, f0)
}

func TestRecording(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping test on windows")
	}
	f0 := e2e.CaseRecording(kwokctlPath, clusterName, pwd).
		Feature()
	testEnv.Test(t, f0)
}

func TestRecordingExternal(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping test on windows")
	}
	f0 := e2e.CaseRecordingExternal(kwokctlPath, clusterName, pwd).
		Feature()
	testEnv.Test(t, f0)
}

func TestPortForward(t *testing.T) {
	f0 := e2e.CasePortForward(kwokctlPath, clusterName, envconf.RandomName("node", 16), namespace).
		Feature()
	testEnv.Test(t, f0)
}

func TestLogs(t *testing.T) {
	f0 := e2e.CaseLogs(kwokctlPath, clusterName, envconf.RandomName("node", 16), namespace, pwd).
		Feature()
	testEnv.Test(t, f0)
}

func TestAttach(t *testing.T) {
	f0 := e2e.CaseAttach(kwokctlPath, clusterName, envconf.RandomName("node", 16), namespace, pwd).
		Feature()
	testEnv.Test(t, f0)
}

func TestExec(t *testing.T) {
	f0 := e2e.CaseExec(kwokctlPath, clusterName, envconf.RandomName("node", 16), namespace).
		Feature()
	testEnv.Test(t, f0)
}

func TestRestart(t *testing.T) {
	f0 := e2e.CaseRestart(kwokctlPath, clusterName).
		Feature()
	testEnv.Test(t, f0)
}

func TestSnapshot(t *testing.T) {
	f0 := e2e.CaseSnapshot(kwokctlPath, clusterName, pwd).
		Feature()
	testEnv.Test(t, f0)
}

func TestResourceUsage(t *testing.T) {
	f0 := e2e.CaseResourceUsage(kwokctlPath, clusterName).
		Feature()
	testEnv.Test(t, f0)
}
