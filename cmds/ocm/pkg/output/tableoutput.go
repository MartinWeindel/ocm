// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package output

import (
	"strings"

	"github.com/gardener/ocm/cmds/ocm/pkg/data"
	. "github.com/gardener/ocm/cmds/ocm/pkg/processing"
	"github.com/gardener/ocm/pkg/errors"
)

type TableProcessingOutput struct {
	ElementOutput
	header []string
	opts   *Options
}

var _ Output = &TableProcessingOutput{}

func NewProcessingTableOutput(opts *Options, chain ProcessChain, header ...string) *TableProcessingOutput {
	return (&TableProcessingOutput{}).new(opts, chain, header)
}

func (this *TableProcessingOutput) new(opts *Options, chain ProcessChain, header []string) *TableProcessingOutput {
	this.header = header
	this.ElementOutput.new(opts.Context, chain)
	this.opts = opts
	return this
}

func (this *TableProcessingOutput) Out() error {
	lines := [][]string{this.header}

	sort := this.opts.Sort
	slice := data.IndexedSliceAccess(data.Slice(this.Elems))
	if sort != nil {
		cols := make([]string, len(this.header))
		idxs := map[string]int{}
		for i, n := range this.header {
			cols[i] = strings.ToLower(n)
			if strings.HasPrefix(cols[i], "-") {
				cols[i] = cols[i][1:]
			}
			idxs[cols[i]] = i
		}
		for _, k := range sort {
			key := SelectBest(strings.ToLower(k), cols...)
			if key == "" {
				return errors.Newf("unknown field '%s'", k)
			}
			slice.Sort(compare_column(idxs[key]))
		}
	}

	FormatTable(this.Context, "", append(lines, data.StringArraySlice(slice)...))
	return nil
}

func compare_column(c int) CompareFunction {
	return func(a interface{}, b interface{}) int {
		aa := a.([]string)
		ab := b.([]string)
		if len(aa) > c && len(ab) > c {
			return strings.Compare(aa[c], ab[c])
		}
		return len(aa) - len(ab)
	}

}
