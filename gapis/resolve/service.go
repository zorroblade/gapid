// Copyright (C) 2017 Google Inc.
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

package resolve

import (
	"github.com/google/gapid/gapis/atom"
	"github.com/google/gapid/gapis/service"
	"github.com/google/gapid/gapis/service/box"
)

func internalToService(v interface{}) (interface{}, error) {
	switch v := v.(type) {
	case atom.Atom:
		return atom.ToService(v)
	case *InternalContext:
		return &service.Context{Name: v.Name, Api: v.Api}, nil
	default:
		return v, nil
	}
}

func serviceToInternal(v interface{}) (interface{}, error) {
	switch v := v.(type) {
	case *service.Command:
		return atom.ToAtom(v)
	case *box.Value:
		return v.Get(), nil
	default:
		return v, nil
	}
}