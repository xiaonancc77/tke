/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package util

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"tkestack.io/tke/api/auth"
	v1 "tkestack.io/tke/api/auth/v1"
	"tkestack.io/tke/pkg/apiserver/authentication"
)

// FilterLocalIdentity is used to filter localIdentity that do not belong to the tenant.
func FilterLocalIdentity(ctx context.Context, localIdentity *auth.LocalIdentity) error {
	_, tenantID := authentication.GetUsernameAndTenantID(ctx)
	if tenantID == "" {
		return nil
	}
	if localIdentity.Spec.TenantID != tenantID {
		return errors.NewNotFound(v1.Resource("localIdentity"), localIdentity.ObjectMeta.Name)
	}
	return nil
}

// FilterAPIKey is used to filter apiKey that do not belong to the tenant.
func FilterAPIKey(ctx context.Context, apiKey *auth.APIKey) error {
	username, tenantID := authentication.GetUsernameAndTenantID(ctx)
	if tenantID == "" {
		return nil
	}
	if apiKey.Spec.TenantID != tenantID || apiKey.Spec.Username != username {
		return errors.NewNotFound(v1.Resource("apiKey"), apiKey.ObjectMeta.Name)
	}

	return nil
}
