//####################################################################
//
// File: b2/validators.go
//
// Copyright 2021 Backblaze Inc. All Rights Reserved.
//
// License https://www.backblaze.com/using_b2_code.html
//
//####################################################################

package b2

import (
	"encoding/base64"
	"fmt"
)

func validateBase64Key(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if ok {
		decoded, err := base64.StdEncoding.DecodeString(v)
		if err == nil {
			// AES256 (which is the only supported algorithm for now) key should be 256 bits (32 bytes)
			if len(decoded) != 32 {
				errors = append(errors, fmt.Errorf("AES256 key should be 32 bytes, got %d bytes instead",
					len(decoded)))
			}
		} else {
			errors = append(errors, err)
		}
	} else {
		errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
	}

	return warnings, errors
}
