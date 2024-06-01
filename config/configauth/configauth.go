// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configauth implements the configuration settings to
// ensure authentication on incoming requests, and allows
// exporters to add authentication on outgoing requests.
package configauth // import "github.com/oodle-ai/opentelemetry-collector/config/configauth"

import (
	"errors"
	"fmt"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/extension/auth"
)

var (
	errAuthenticatorNotFound = errors.New("authenticator not found")
	errNotClient             = errors.New("requested authenticator is not a client authenticator")
	errNotServer             = errors.New("requested authenticator is not a server authenticator")
)

// Authentication defines the auth settings for the receiver.
type Authentication struct {
	// AuthenticatorID specifies the name of the extension to use in order to authenticate the incoming data point.
	AuthenticatorID component.ID `mapstructure:"authenticator"`
}

// NewDefaultAuthentication returns a default authentication configuration.
func NewDefaultAuthentication() *Authentication {
	return &Authentication{}
}

// GetServerAuthenticator attempts to select the appropriate auth.Server from the list of extensions,
// based on the requested extension name. If an authenticator is not found, an error is returned.
func (a Authentication) GetServerAuthenticator(extensions map[component.ID]component.Component) (auth.Server, error) {
	if ext, found := extensions[a.AuthenticatorID]; found {
		if server, ok := ext.(auth.Server); ok {
			return server, nil
		}
		return nil, errNotServer
	}

	return nil, fmt.Errorf("failed to resolve authenticator %q: %w", a.AuthenticatorID, errAuthenticatorNotFound)
}

// GetClientAuthenticator attempts to select the appropriate auth.Client from the list of extensions,
// based on the component id of the extension. If an authenticator is not found, an error is returned.
// This should be only used by HTTP clients.
func (a Authentication) GetClientAuthenticator(extensions map[component.ID]component.Component) (auth.Client, error) {
	if ext, found := extensions[a.AuthenticatorID]; found {
		if client, ok := ext.(auth.Client); ok {
			return client, nil
		}
		return nil, errNotClient
	}
	return nil, fmt.Errorf("failed to resolve authenticator %q: %w", a.AuthenticatorID, errAuthenticatorNotFound)
}
