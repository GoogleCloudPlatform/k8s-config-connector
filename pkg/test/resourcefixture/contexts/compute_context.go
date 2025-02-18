// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package contexts

func init() {
	resourceContextMap["regionalcomputeaddress"] = ResourceContext{
		ResourceKind: "ComputeAddress",
		SkipUpdate:   true,
		// TestCreateNoChangeUpdateDelete/basic-regionalcomputeaddress: dynamic_controller_integration_test.go:239:
		//    reconcile returned unexpected error: Update call failed: error applying desired state: Error creating
		//   Address: googleapi: Error 400: Invalid value for field 'resource.networkTier': 'PREMIUM'. An address
		//  with type INTERNAL cannot have a network tier., invalid
		SkipDriftDetection: true,
	}

	resourceContextMap["globalcomputeaddress"] = ResourceContext{
		ResourceKind: "ComputeAddress",
		SkipUpdate:   true,
	}

	resourceContextMap["globalcomputebackendservicesecuritysettings"] = ResourceContext{
		ResourceKind: "ComputeBackendService",
		// Underlying API changes dependency resource's project id to number after successful creation.
		// For now TF servicemapping does not have a way to resolve dependency DCL resources' project number.
		// Skip checking no change after creation(testNoChangeAfterCreate) to bypass this temporarily.
		// See https://buganizer.corp.google.com/issues/374166656#comment11 for details.
		SkipNoChange: true,
	}

	resourceContextMap["computeexternalvpngateway"] = ResourceContext{
		ResourceKind: "ComputeExternalVPNGateway",
		SkipUpdate:   true,
	}
	resourceContextMap["computemanagedsslcertificate"] = ResourceContext{
		ResourceKind: "ComputeManagedSSLCertificate",
		// This resource doesn't support update.
		SkipUpdate: true,
	}
	resourceContextMap["cloudfunctioncomputeregionnetworkendpointgroup"] = ResourceContext{
		ResourceKind: "ComputeRegionNetworkEndpointGroup",
		// The GCP resource for ComputeRegionNetworkEndpointGroup doesn't
		// support update.
		SkipUpdate: true,
	}

	resourceContextMap["cloudruncomputeregionnetworkendpointgroup"] = ResourceContext{
		ResourceKind: "ComputeRegionNetworkEndpointGroup",
		// The GCP resource for ComputeRegionNetworkEndpointGroup doesn't
		// support update.
		SkipUpdate: true,
	}

	resourceContextMap["privateserviceconnectioncomputeregionnetworkendpointgroup"] = ResourceContext{
		ResourceKind: "ComputeRegionNetworkEndpointGroup",
		// The GCP resource for ComputeRegionNetworkEndpointGroup doesn't
		// support update.
		SkipUpdate: true,
	}

	resourceContextMap["computenetworkendpointgroup"] = ResourceContext{
		ResourceKind: "ComputeNetworkEndpointGroup",
		SkipUpdate:   true,
	}

	resourceContextMap["computenetworkpeering"] = ResourceContext{
		ResourceKind: "ComputeNetworkPeering",
		SkipUpdate:   true,
	}

	resourceContextMap["computenodetemplate"] = ResourceContext{
		ResourceKind: "ComputeNodeTemplate",
		SkipUpdate:   true,
	}

	resourceContextMap["computenodegroup"] = ResourceContext{
		ResourceKind: "ComputeNodeGroup",
		SkipUpdate:   true, // The only field which supports update is nodeTemplateRef, which currently cannot be used for testing updates because of b/147506185
	}

	resourceContextMap["computesharedvpchostproject"] = ResourceContext{
		ResourceKind: "ComputeSharedVPCHostProject",
		SkipUpdate:   true,
	}

	resourceContextMap["computesharedvpcserviceproject"] = ResourceContext{
		ResourceKind: "ComputeSharedVPCServiceProject",
		SkipUpdate:   true,
	}
	resourceContextMap["computesslcertificate"] = ResourceContext{
		ResourceKind: "ComputeSSLCertificate",
		SkipUpdate:   true, // No input fields in this resource support update.
	}

	resourceContextMap["globalcomputesslcertificate"] = ResourceContext{
		ResourceKind: "ComputeSSLCertificate",
		SkipUpdate:   true, // No input fields in this resource support update.
	}

	resourceContextMap["regionalcomputesslcertificate"] = ResourceContext{
		ResourceKind: "ComputeSSLCertificate",
		SkipUpdate:   true, // No input fields in this resource support update.
	}

	resourceContextMap["computetargetgrpcproxy"] = ResourceContext{
		ResourceKind: "ComputeTargetGRPCProxy",
		SkipUpdate:   true, // No input fields in this resource support update.
	}

	resourceContextMap["computereservation"] = ResourceContext{
		ResourceKind: "ComputeReservation",
		SkipUpdate:   true, // No input fields in this resource support update.
	}

	resourceContextMap["computeresourcepolicy"] = ResourceContext{
		ResourceKind: "ComputeResourcePolicy",
		SkipUpdate:   true,
	}

	resourceContextMap["computerouterinterface"] = ResourceContext{
		ResourceKind: "ComputeRouterInterface",
		SkipUpdate:   true,
	}

	resourceContextMap["computerouternat"] = ResourceContext{
		ResourceKind: "ComputeRouterNAT",
		SkipUpdate:   true,
	}

	resourceContextMap["computerouterpeer"] = ResourceContext{
		ResourceKind: "ComputeRouterPeer",
		SkipUpdate:   true,
	}

	resourceContextMap["computeinstance"] = ResourceContext{
		ResourceKind: "ComputeInstance",
		// TestCreateNoChangeUpdateDelete/basic-computeinstance: dynamic_controller_integration_test.go:239: reconcile
		//    returned unexpected error: Update call failed: error applying desired state: Error creating instance:
		//    googleapi: Error 400: Invalid value for field 'resource.disks[0]': '{  "mode": "READ_WRITE",  "source":
		//    "projects/cnrm-test-tgooo56g38yqbn3k/zones/us-west1-a/disks/comp...'. Cannot specify both 'source' and
		//    'initializeParams'., invalid
		SkipDriftDetection: true,
	}

	resourceContextMap["computeinstancewithencrypteddisk"] = ResourceContext{
		ResourceKind: "ComputeInstance",
		SkipNoChange: true, // Encryption key can't be retrieved so the live state will always be different from the intended state.
		// TestCreateNoChangeUpdateDelete/basic-computeinstancewithencrypteddisk: dynamic_controller_integration_test.go:239:
		//    reconcile returned unexpected error: Update call failed: error applying desired state: Error creating instance:
		//    googleapi: Error 400: Invalid value for field 'resource.disks[0]': '{  "mode": "READ_WRITE",  "source":
		//    "projects/cnrm-test-tgooo56g38yqbn3k/zones/us-west1-a/disks/comp...'. Cannot specify both 'source' and
		//    'initializeParams'., invalid
		SkipDriftDetection: true,
	}

	resourceContextMap["computeinstancefromtemplate"] = ResourceContext{
		ResourceKind: "ComputeInstance",
		// TestCreateNoChangeUpdateDelete/basic-computeinstancefromtemplate: dynamic_controller_integration_test.go:239:
		//   reconcile returned unexpected error: Update call failed: error applying desired state: Error creating instance:
		//   googleapi: Error 400: Invalid value for field 'resource.disks[0]': '{  "mode": "READ_WRITE",  "source":
		//   "projects/cnrm-test-tgooo56g38yqbn3k/zones/us-west1-a/disks/comp...'. Cannot specify both 'source' and
		//   'initializeParams'., invalid
		SkipDriftDetection: true,
	}

	resourceContextMap["computeinstancetemplate"] = ResourceContext{
		ResourceKind: "ComputeInstanceTemplate",
		SkipUpdate:   true, // Update is not supported for Compute Instance Template in the GCP API and TF provider.
	}

	resourceContextMap["computetargetinstance"] = ResourceContext{
		ResourceKind: "ComputeTargetInstance",
		SkipUpdate:   true, // No input fields in this resource support update.
	}

	resourceContextMap["computetargetsslproxy"] = ResourceContext{
		ResourceKind: "ComputeTargetSSLProxy",
		SkipUpdate:   true,
	}

	resourceContextMap["computevpngateway"] = ResourceContext{
		ResourceKind: "ComputeVPNGateway",
		SkipUpdate:   true,
	}

	resourceContextMap["computevpntunnel"] = ResourceContext{
		ResourceKind: "ComputeVPNTunnel",
		// TestCreateNoChangeUpdateDelete/basic-computevpntunnel: dynamic_controller_integration_test.go:149: value
		//    mismatch for label with key 'label-one': got 'value-two', want 'value-one'
		SkipDriftDetection: true,
	}

	resourceContextMap["computetargetvpngateway"] = ResourceContext{
		ResourceKind: "ComputeTargetVPNGateway",
		SkipUpdate:   true,
	}

	resourceContextMap["regionalcomputedisk"] = ResourceContext{
		ResourceKind: "ComputeDisk",
		SkipUpdate:   true,
	}

	resourceContextMap["zonalcomputedisk"] = ResourceContext{
		ResourceKind: "ComputeDisk",
		SkipUpdate:   true,
	}

	resourceContextMap["computediskfromsourcedisk"] = ResourceContext{
		ResourceKind: "ComputeDisk",
		SkipUpdate:   true,
	}

	resourceContextMap["computesnapshot"] = ResourceContext{
		ResourceKind: "ComputeSnapshot",
		SkipUpdate:   true,
	}

	resourceContextMap["computeroute"] = ResourceContext{
		ResourceKind: "ComputeRoute",
		SkipUpdate:   true,
	}

	resourceContextMap["externalwithpartialuri"] = resourceContextMap["computevpngateway"]

	resourceContextMap["sensitivevaluesimple"] = resourceContextMap["computesslcertificate"]

	resourceContextMap["sensitivevaluefromsecret"] = resourceContextMap["computesslcertificate"]

	resourceContextMap["foldercomputefirewallpolicyassociation"] = ResourceContext{
		ResourceKind: "ComputeFirewallPolicyAssociation",
		SkipUpdate:   true, // No update method in DCL at the moment
	}

	resourceContextMap["organizationcomputefirewallpolicyassociation"] = ResourceContext{
		ResourceKind: "ComputeFirewallPolicyAssociation",
		SkipUpdate:   true, // No update method in DCL at the moment
	}

	resourceContextMap["computenetworkfirewallpolicyassociation"] = ResourceContext{
		ResourceKind: "ComputeNetworkFirewallPolicyAssociation",
		SkipUpdate:   true, // No input fields in this resource support update.
	}
}
