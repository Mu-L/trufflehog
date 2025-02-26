// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: custom_detectors.proto

package custom_detectorspb

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

// Validate checks the field values on CustomDetectors with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CustomDetectors) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CustomDetectors with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CustomDetectorsMultiError, or nil if none found.
func (m *CustomDetectors) ValidateAll() error {
	return m.validate(true)
}

func (m *CustomDetectors) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDetectors() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CustomDetectorsValidationError{
						field:  fmt.Sprintf("Detectors[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CustomDetectorsValidationError{
						field:  fmt.Sprintf("Detectors[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CustomDetectorsValidationError{
					field:  fmt.Sprintf("Detectors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CustomDetectorsMultiError(errors)
	}

	return nil
}

// CustomDetectorsMultiError is an error wrapping multiple validation errors
// returned by CustomDetectors.ValidateAll() if the designated constraints
// aren't met.
type CustomDetectorsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CustomDetectorsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CustomDetectorsMultiError) AllErrors() []error { return m }

// CustomDetectorsValidationError is the validation error returned by
// CustomDetectors.Validate if the designated constraints aren't met.
type CustomDetectorsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomDetectorsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomDetectorsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomDetectorsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomDetectorsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomDetectorsValidationError) ErrorName() string { return "CustomDetectorsValidationError" }

// Error satisfies the builtin error interface
func (e CustomDetectorsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomDetectors.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomDetectorsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomDetectorsValidationError{}

// Validate checks the field values on CustomRegex with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CustomRegex) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CustomRegex with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CustomRegexMultiError, or
// nil if none found.
func (m *CustomRegex) ValidateAll() error {
	return m.validate(true)
}

func (m *CustomRegex) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Regex

	for idx, item := range m.GetVerify() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CustomRegexValidationError{
						field:  fmt.Sprintf("Verify[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CustomRegexValidationError{
						field:  fmt.Sprintf("Verify[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CustomRegexValidationError{
					field:  fmt.Sprintf("Verify[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Description

	// no validation rules for Entropy

	if len(errors) > 0 {
		return CustomRegexMultiError(errors)
	}

	return nil
}

// CustomRegexMultiError is an error wrapping multiple validation errors
// returned by CustomRegex.ValidateAll() if the designated constraints aren't met.
type CustomRegexMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CustomRegexMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CustomRegexMultiError) AllErrors() []error { return m }

// CustomRegexValidationError is the validation error returned by
// CustomRegex.Validate if the designated constraints aren't met.
type CustomRegexValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomRegexValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomRegexValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomRegexValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomRegexValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomRegexValidationError) ErrorName() string { return "CustomRegexValidationError" }

// Error satisfies the builtin error interface
func (e CustomRegexValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomRegex.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomRegexValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomRegexValidationError{}

// Validate checks the field values on VerifierConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *VerifierConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifierConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in VerifierConfigMultiError,
// or nil if none found.
func (m *VerifierConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifierConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, err := url.Parse(m.GetEndpoint()); err != nil {
		err = VerifierConfigValidationError{
			field:  "Endpoint",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Unsafe

	if len(errors) > 0 {
		return VerifierConfigMultiError(errors)
	}

	return nil
}

// VerifierConfigMultiError is an error wrapping multiple validation errors
// returned by VerifierConfig.ValidateAll() if the designated constraints
// aren't met.
type VerifierConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifierConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifierConfigMultiError) AllErrors() []error { return m }

// VerifierConfigValidationError is the validation error returned by
// VerifierConfig.Validate if the designated constraints aren't met.
type VerifierConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifierConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifierConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifierConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifierConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifierConfigValidationError) ErrorName() string { return "VerifierConfigValidationError" }

// Error satisfies the builtin error interface
func (e VerifierConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifierConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifierConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifierConfigValidationError{}
