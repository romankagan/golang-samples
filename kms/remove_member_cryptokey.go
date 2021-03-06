// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kms

// [START kms_remove_member_from_cryptokey_policy]
import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/iam"
	cloudkms "cloud.google.com/go/kms/apiv1"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

// removeMemberCryptoKeyPolicy removes a specified member from an IAM role for the key.
func removeMemberCryptoKeyPolicy(w io.Writer, name, member string, role iam.RoleName) error {
	// name: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"
	// member := "user@gmail.com"
	// role := iam.Viewer
	ctx := context.Background()
	client, err := cloudkms.NewKeyManagementClient(ctx)
	if err != nil {
		return fmt.Errorf("cloudkms.NewKeyManagementClient: %v", err)
	}
	// Get the desired CryptoKey.
	keyObj, err := client.GetCryptoKey(ctx, &kmspb.GetCryptoKeyRequest{Name: name})
	if err != nil {
		return fmt.Errorf("GetCryptoKey: %v", err)
	}
	// Get IAM Policy.
	handle := client.CryptoKeyIAM(keyObj)
	policy, err := handle.Policy(ctx)
	if err != nil {
		return fmt.Errorf("Policy: %v", err)
	}
	// Remove Member.
	policy.Remove(member, role)
	if err = handle.SetPolicy(ctx, policy); err != nil {
		return fmt.Errorf("SetPolicy: %v", err)
	}
	fmt.Fprintf(w, "Removed member %s from cryptokey policy.", member)
	return nil
}

// [END kms_remove_member_from_cryptokey_policy]
