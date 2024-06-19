// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: relations/v0/lookup.proto

package v0

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on LookupSubjectsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LookupSubjectsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LookupSubjectsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LookupSubjectsRequestMultiError, or nil if none found.
func (m *LookupSubjectsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LookupSubjectsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetResource() == nil {
		err := LookupSubjectsRequestValidationError{
			field:  "Resource",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if a := m.GetResource(); a != nil {

	}

	if utf8.RuneCountInString(m.GetRelation()) < 1 {
		err := LookupSubjectsRequestValidationError{
			field:  "Relation",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSubjectType() == nil {
		err := LookupSubjectsRequestValidationError{
			field:  "SubjectType",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if a := m.GetSubjectType(); a != nil {

	}

	if m.SubjectRelation != nil {
		// no validation rules for SubjectRelation
	}

	if m.Pagination != nil {

		if all {
			switch v := interface{}(m.GetPagination()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupSubjectsRequestValidationError{
						field:  "Pagination",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupSubjectsRequestValidationError{
						field:  "Pagination",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupSubjectsRequestValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return LookupSubjectsRequestMultiError(errors)
	}

	return nil
}

// LookupSubjectsRequestMultiError is an error wrapping multiple validation
// errors returned by LookupSubjectsRequest.ValidateAll() if the designated
// constraints aren't met.
type LookupSubjectsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LookupSubjectsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LookupSubjectsRequestMultiError) AllErrors() []error { return m }

// LookupSubjectsRequestValidationError is the validation error returned by
// LookupSubjectsRequest.Validate if the designated constraints aren't met.
type LookupSubjectsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupSubjectsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupSubjectsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupSubjectsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupSubjectsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupSubjectsRequestValidationError) ErrorName() string {
	return "LookupSubjectsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LookupSubjectsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupSubjectsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupSubjectsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupSubjectsRequestValidationError{}

// Validate checks the field values on LookupSubjectsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LookupSubjectsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LookupSubjectsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LookupSubjectsResponseMultiError, or nil if none found.
func (m *LookupSubjectsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LookupSubjectsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSubject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LookupSubjectsResponseValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LookupSubjectsResponseValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSubject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LookupSubjectsResponseValidationError{
				field:  "Subject",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPagination()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LookupSubjectsResponseValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LookupSubjectsResponseValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LookupSubjectsResponseValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LookupSubjectsResponseMultiError(errors)
	}

	return nil
}

// LookupSubjectsResponseMultiError is an error wrapping multiple validation
// errors returned by LookupSubjectsResponse.ValidateAll() if the designated
// constraints aren't met.
type LookupSubjectsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LookupSubjectsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LookupSubjectsResponseMultiError) AllErrors() []error { return m }

// LookupSubjectsResponseValidationError is the validation error returned by
// LookupSubjectsResponse.Validate if the designated constraints aren't met.
type LookupSubjectsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupSubjectsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupSubjectsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupSubjectsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupSubjectsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupSubjectsResponseValidationError) ErrorName() string {
	return "LookupSubjectsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LookupSubjectsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupSubjectsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupSubjectsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupSubjectsResponseValidationError{}
