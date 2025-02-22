// Copyright 2023 Matrix Origin
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

package fileservice

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/matrixorigin/matrixone/pkg/common/moerr"
)

type ObjectStorageArguments struct {
	Name                string
	Endpoint            string
	Region              string
	Bucket              string
	KeyID               string
	KeySecret           string
	SessionToken        string
	SecurityToken       string
	BearerToken         string
	KeyPrefix           string
	RoleARN             string
	RoleSessionName     string
	RAMRole             string
	ExternalID          string
	SharedConfigProfile string

	IsMinio bool
}

func (o *ObjectStorageArguments) SetFromString(arguments []string) error {
	for _, pair := range arguments {
		key, value, ok := strings.Cut(pair, "=")
		if !ok {
			return moerr.NewInvalidInputNoCtx("invalid S3 argument: %s", pair)
		}

		switch key {
		case "endpoint":
			o.Endpoint = value
		case "region":
			o.Region = value
		case "bucket":
			o.Bucket = value
		case "key":
			o.KeyID = value
		case "secret":
			o.KeySecret = value
		case "token", "session-token":
			o.SessionToken = value
		case "security-token":
			o.SecurityToken = value
		case "bearer-token":
			o.BearerToken = value
		case "prefix":
			o.KeyPrefix = value
		case "role-arn":
			o.RoleARN = value
		case "role-session-name":
			o.RoleSessionName = value
		case "ram-role":
			o.RAMRole = value
		case "external-id":
			o.ExternalID = value
		case "name":
			o.Name = value
		case "shared-config-profile":
			o.SharedConfigProfile = value
		case "is-minio":
			o.IsMinio = value != "false" && value != "0"
		default:
			return moerr.NewInvalidInputNoCtx("invalid S3 argument: %s", pair)
		}

	}
	return nil
}

func (o *ObjectStorageArguments) validate() error {

	// validate endpoint
	var endpointURL *url.URL
	if o.Endpoint != "" {
		var err error
		endpointURL, err = url.Parse(o.Endpoint)
		if err != nil {
			return err
		}
		if endpointURL.Scheme == "" {
			endpointURL.Scheme = "https"
		}
		o.Endpoint = endpointURL.String()
	}

	// region
	if o.Region == "" {
		// try to get region from bucket
		// only works for AWS S3
		resp, err := http.Head("https://" + o.Bucket + ".s3.amazonaws.com")
		if err == nil {
			if value := resp.Header.Get("x-amz-bucket-region"); value != "" {
				o.Region = value
			}
		}
	}

	return nil
}
