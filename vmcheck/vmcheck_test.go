// Copyright 2016-2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vmcheck

import (
	"testing"

	"github.com/vmware/vmw-guestinfo/internal"
)

func TestIsVirtualWorld(t *testing.T) {
	isBackdoor, err := hypervisorPortCheck()
	internal.AssertNoError(t, err)

	t.Log("Backdoor available: ", isBackdoor)
	t.Log("CPU HV: ", IsVirtualCPU())

	isVM, err := IsVirtualWorld()
	internal.AssertNoError(t, err)
	t.Log("Running in a VM: ", isVM)
}
