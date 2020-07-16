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

package template

import (
	"sync"

	v1 "tkestack.io/tke/api/platform/v1"
)

type templateHealth struct {
	mu sync.Mutex
	m  map[string]*v1.Template
}

func (s *templateHealth) Exist(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.m[key]
	return ok
}

func (s *templateHealth) Set(key string, template *v1.Template) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = template
}

func (s *templateHealth) Del(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
}