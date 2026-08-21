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

package shared

import (
	"strings"
	"unicode"
)

var Acronyms = []string{
	"API", "BGP", "BYOID", "CA", "CDN", "CIDR", "CPU", "DNS", "EUC", "FS", "FQDN",
	"GCE", "GB", "GCS", "GKE", "HTML", "HTTP", "HTTPS", "IAM", "IAP", "ID", "IP",
	"IPV4", "IPV6", "KMS", "MiB", "NAT", "OAuth2", "OIDC", "OS", "PD", "PSC",
	"SQL", "SSH", "SSL", "SSO", "TCP", "TLS", "TTL", "UDP", "URI", "URL", "VTPM",
	"VM", "VPC", "VIP", "VPN", "X509", "VPC", "LRO", "KRM",
}

// SplitWords splits an underscore_separated or camelCaseWord string into an array of individual words.
func SplitWords(s string) []string {
	s = strings.ReplaceAll(s, "_", " ")
	var res strings.Builder
	for i, r := range s {
		if i > 0 {
			prev := rune(s[i-1])
			if (unicode.IsLower(prev) && unicode.IsUpper(r)) || (unicode.IsDigit(prev) && unicode.IsUpper(r)) || (unicode.IsLower(prev) && unicode.IsDigit(r)) {
				res.WriteRune(' ')
			}
		}
		res.WriteRune(r)
	}
	return strings.Fields(res.String())
}

func GoFieldName(jsonName string) string {
	words := SplitWords(jsonName)
	for i, w := range words {
		// Check acronyms
		isAcronym := false
		for _, acr := range Acronyms {
			if strings.EqualFold(w, acr) {
				words[i] = acr
				isAcronym = true
				break
			}
		}
		if !isAcronym {
			if len(w) > 0 {
				runes := []rune(w)
				runes[0] = unicode.ToUpper(runes[0])
				words[i] = string(runes)
			}
		}
	}
	return strings.Join(words, "")
}
