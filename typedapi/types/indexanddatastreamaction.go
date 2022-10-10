// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/e0ea3dc890d394d682096cc862b3bd879d9422e9


package types

// IndexAndDataStreamAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e0ea3dc890d394d682096cc862b3bd879d9422e9/specification/indices/modify_data_stream/types.ts#L28-L31
type IndexAndDataStreamAction struct {
	DataStream DataStreamName `json:"data_stream"`
	Index      IndexName      `json:"index"`
}

// IndexAndDataStreamActionBuilder holds IndexAndDataStreamAction struct and provides a builder API.
type IndexAndDataStreamActionBuilder struct {
	v *IndexAndDataStreamAction
}

// NewIndexAndDataStreamAction provides a builder for the IndexAndDataStreamAction struct.
func NewIndexAndDataStreamActionBuilder() *IndexAndDataStreamActionBuilder {
	r := IndexAndDataStreamActionBuilder{
		&IndexAndDataStreamAction{},
	}

	return &r
}

// Build finalize the chain and returns the IndexAndDataStreamAction struct
func (rb *IndexAndDataStreamActionBuilder) Build() IndexAndDataStreamAction {
	return *rb.v
}

func (rb *IndexAndDataStreamActionBuilder) DataStream(datastream DataStreamName) *IndexAndDataStreamActionBuilder {
	rb.v.DataStream = datastream
	return rb
}

func (rb *IndexAndDataStreamActionBuilder) Index(index IndexName) *IndexAndDataStreamActionBuilder {
	rb.v.Index = index
	return rb
}