/*
METRO2 API

Moov Metro2 ([Automated Clearing House](https://en.wikipedia.org/wiki/Automated_Clearing_House)) implements an HTTP API for creating, parsing and validating Metro2 files. Metro2 is an open-source consumer credit history report for credit report file creation and validation.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// File struct for File
type File struct {
	Header HeaderRecord `json:"header"`
	Data []DataRecord `json:"data,omitempty"`
	Trailer TrailerRecord `json:"trailer"`
}

// NewFile instantiates a new File object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFile(header HeaderRecord, trailer TrailerRecord) *File {
	this := File{}
	this.Header = header
	this.Trailer = trailer
	return &this
}

// NewFileWithDefaults instantiates a new File object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileWithDefaults() *File {
	this := File{}
	return &this
}

// GetHeader returns the Header field value
func (o *File) GetHeader() HeaderRecord {
	if o == nil {
		var ret HeaderRecord
		return ret
	}

	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value
// and a boolean to check if the value has been set.
func (o *File) GetHeaderOk() (*HeaderRecord, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Header, true
}

// SetHeader sets field value
func (o *File) SetHeader(v HeaderRecord) {
	o.Header = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *File) GetData() []DataRecord {
	if o == nil || o.Data == nil {
		var ret []DataRecord
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetDataOk() ([]DataRecord, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *File) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []DataRecord and assigns it to the Data field.
func (o *File) SetData(v []DataRecord) {
	o.Data = v
}

// GetTrailer returns the Trailer field value
func (o *File) GetTrailer() TrailerRecord {
	if o == nil {
		var ret TrailerRecord
		return ret
	}

	return o.Trailer
}

// GetTrailerOk returns a tuple with the Trailer field value
// and a boolean to check if the value has been set.
func (o *File) GetTrailerOk() (*TrailerRecord, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Trailer, true
}

// SetTrailer sets field value
func (o *File) SetTrailer(v TrailerRecord) {
	o.Trailer = v
}

func (o File) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["header"] = o.Header
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["trailer"] = o.Trailer
	}
	return json.Marshal(toSerialize)
}

type NullableFile struct {
	value *File
	isSet bool
}

func (v NullableFile) Get() *File {
	return v.value
}

func (v *NullableFile) Set(val *File) {
	v.value = val
	v.isSet = true
}

func (v NullableFile) IsSet() bool {
	return v.isSet
}

func (v *NullableFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFile(val *File) *NullableFile {
	return &NullableFile{value: val, isSet: true}
}

func (v NullableFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


