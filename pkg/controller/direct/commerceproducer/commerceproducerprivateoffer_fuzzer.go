// Copyright 2026 Google LLC
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

// +tool:fuzz-gen
// proto.message: google.cloud.commerceproducer.v1beta.PrivateOffer
// api.group: commerceproducer.cnrm.cloud.google.com

package commerceproducer

import (
	pb "cloud.google.com/go/commerceproducer/apiv1beta/commerceproducerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/commerceproducer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(commerceProducerPrivateOfferFuzzer())
}

func commerceProducerPrivateOfferFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*pb.PrivateOffer, krm.CommerceProducerPrivateOfferSpec, krm.CommerceProducerPrivateOfferObservedState](
		&pb.PrivateOffer{},
		CommerceProducerPrivateOfferSpec_FromProto, CommerceProducerPrivateOfferSpec_ToProto,
		CommerceProducerPrivateOfferObservedState_FromProto, CommerceProducerPrivateOfferObservedState_ToProto,
	)

	f.IdentityField(".name")

	// Spec fields
	f.SpecField(".title")
	f.SpecField(".internal_note")
	f.SpecField(".customer_note")
	f.SpecField(".offer_deal_type")
	f.SpecField(".partner_contact")
	f.SpecField(".customer")
	f.SpecField(".accept_deadline_time")
	f.SpecField(".term")
	f.SpecField(".single_product_offer")

	// Status fields
	f.StatusField(".state")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".publish_time")
	f.StatusField(".accept_time")
	f.StatusField(".cancel_time")
	f.StatusField(".cancellation_reason")
	f.StatusField(".cancellation_note")
	f.StatusField(".start_time")
	f.StatusField(".end_time")
	f.StatusField(".previous_state")
	f.StatusField(".offer_term")
	f.StatusField(".reseller_contact")
	f.StatusField(".signer")
	f.StatusField(".purchase_order_id")

	// Unimplemented fields
	f.Unimplemented_NotYetTriaged(".publish_requirement_google_review")
	f.Unimplemented_NotYetTriaged(".customer.address")
	f.Unimplemented_NotYetTriaged(".term.unlimited_renewal")
	f.Unimplemented_NotYetTriaged(".term.effective_term_end_time")
	f.Unimplemented_NotYetTriaged(".single_product_offer.service_level")
	f.Unimplemented_NotYetTriaged(".single_product_offer.reseller_private_offer_plan_id")
	f.Unimplemented_NotYetTriaged(".single_product_offer.effective_installment_timeline")
	f.Unimplemented_NotYetTriaged(".single_product_offer.contract_value")
	f.Unimplemented_NotYetTriaged(".single_product_offer.revenue_share")

	return f
}
