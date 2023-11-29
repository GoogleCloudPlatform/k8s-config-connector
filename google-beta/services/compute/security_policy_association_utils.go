// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute

import (
	"log"
	"strings"

	"github.com/hashicorp/errwrap"
	"google.golang.org/api/googleapi"
)

func transformSecurityPolicyAssociationReadError(err error) error {
	if gErr, ok := errwrap.GetType(err, &googleapi.Error{}).(*googleapi.Error); ok {
		if gErr.Code == 400 && strings.Contains(gErr.Message, "An association with that name does not exist") {
			// This error occurs when attempting a GET after deleting the security policy association. It leads to to
			// inconsistent behavior as HandleNotFoundError(...) expects an error code of 404 when a resource does not
			// exist. To get the desired behavior from HandleNotFoundError, modify the return code to 404 so that
			// HandleNotFoundError(...) will treat this as a NotFound error
			gErr.Code = 404
		}

		log.Printf("[DEBUG] Transformed security policy association error")
		return gErr
	}

	return err
}
