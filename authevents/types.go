// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package authevents

import (
	"encoding/json"
)

type (
	// Message reporting that a client has changed
	//
	// See http://schemas.taskcluster.net/auth/v1/client-message.json#
	ClientMessage struct {

		// `clientId` of the client that was changed
		//
		// Syntax:     ^[A-Za-z0-9@/:.+|_-]+$
		//
		// See http://schemas.taskcluster.net/auth/v1/client-message.json#/properties/clientId
		ClientID string `json:"clientId"`

		// Message version number
		//
		// Possible values:
		//   * 1
		//
		// See http://schemas.taskcluster.net/auth/v1/client-message.json#/properties/version
		Version json.RawMessage `json:"version"`
	}

	// Message reporting that a role has changed
	//
	// See http://schemas.taskcluster.net/auth/v1/role-message.json#
	RoleMessage struct {

		// `roleId` of the role that was changed
		//
		// Syntax:     ^[\x20-\x7e]+$
		//
		// See http://schemas.taskcluster.net/auth/v1/role-message.json#/properties/roleId
		RoleID string `json:"roleId"`

		// Message version number
		//
		// Possible values:
		//   * 1
		//
		// See http://schemas.taskcluster.net/auth/v1/role-message.json#/properties/version
		Version json.RawMessage `json:"version"`
	}
)
