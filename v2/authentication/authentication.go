// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package authentication implements the DocuSign SDK
// category Authentication.
//
// Use the Authentication category to manage various account login tasks including:
//
// * Getting login information for a user.
// * Managing linked social accounts.
// * Getting and revoking OAuth tokens.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/esign-rest-api/v2/reference/Authentication
// Usage example:
//
//   import (
//       "github.com/ofio/esign"
//       "github.com/ofio/esign/v2/authentication"
//       "github.com/ofio/esign/v2/model"
//   )
//   ...
//   authenticationService := authentication.New(esignCredential)
package authentication // import "github.com/ofio/esign/v2/authentication"

import (
	"context"
	"net/url"
	"strings"

	"github.com/ofio/esign"
	"github.com/ofio/esign/v2/model"
)

// Service implements DocuSign Authentication Category API operations
type Service struct {
	credential esign.Credential
}

// New initializes a authentication service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// Login gets login information for a specified user.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/authentication/authentication/login
//
// SDK Method Authentication::login
func (s *Service) Login() *LoginOp {
	return &LoginOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "/v2/login_information",
		QueryOpts:  make(url.Values),
	}
}

// LoginOp implements DocuSign API SDK Authentication::login
type LoginOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *LoginOp) Do(ctx context.Context) (*model.LoginInformation, error) {
	var res *model.LoginInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// APIPassword reserved for DocuSign.
func (op *LoginOp) APIPassword(val string) *LoginOp {
	if op != nil {
		op.QueryOpts.Set("api_password", val)
	}
	return op
}

// EmbedAccountIDGUID set the call query parameter embed_account_id_guid
func (op *LoginOp) EmbedAccountIDGUID(val string) *LoginOp {
	if op != nil {
		op.QueryOpts.Set("embed_account_id_guid", val)
	}
	return op
}

// IncludeAccountIDGUID when set to **true**, shows the account ID GUID in the response.
func (op *LoginOp) IncludeAccountIDGUID() *LoginOp {
	if op != nil {
		op.QueryOpts.Set("include_account_id_guid", "true")
	}
	return op
}

// LoginSettings determines whether login settings are returned in the response.
//
// Valid Values:
//
// * all -  All the login settings are returned.
// * none - no login settings are returned.
func (op *LoginOp) LoginSettings(val string) *LoginOp {
	if op != nil {
		op.QueryOpts.Set("login_settings", val)
	}
	return op
}

// UpdatePassword updates the password for a specified user.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/authentication/authentication/updatepassword
//
// SDK Method Authentication::updatePassword
func (s *Service) UpdatePassword(loginPart string, userPasswordInformation *model.UserPasswordInformation) *UpdatePasswordOp {
	return &UpdatePasswordOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"", "v2", "login_information", loginPart}, "/"),
		Payload:    userPasswordInformation,
		QueryOpts:  make(url.Values),
	}
}

// UpdatePasswordOp implements DocuSign API SDK Authentication::updatePassword
type UpdatePasswordOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdatePasswordOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// UserSocialAccountLoginsDelete deletes user's social account.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/authentication/usersocialaccountlogins/delete
//
// SDK Method Authentication::deleteSocialLogin
func (s *Service) UserSocialAccountLoginsDelete(userID string, userSocialAccountLogins *model.SocialAccountInformation) *UserSocialAccountLoginsDeleteOp {
	return &UserSocialAccountLoginsDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"users", userID, "social"}, "/"),
		Payload:    userSocialAccountLogins,
		QueryOpts:  make(url.Values),
	}
}

// UserSocialAccountLoginsDeleteOp implements DocuSign API SDK Authentication::deleteSocialLogin
type UserSocialAccountLoginsDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UserSocialAccountLoginsDeleteOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// UserSocialAccountLoginsList gets a list of a user's social accounts.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/authentication/usersocialaccountlogins/list
//
// SDK Method Authentication::listSocialLogins
func (s *Service) UserSocialAccountLoginsList(userID string) *UserSocialAccountLoginsListOp {
	return &UserSocialAccountLoginsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"users", userID, "social"}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// UserSocialAccountLoginsListOp implements DocuSign API SDK Authentication::listSocialLogins
type UserSocialAccountLoginsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UserSocialAccountLoginsListOp) Do(ctx context.Context) (*model.UserSocialIDResult, error) {
	var res *model.UserSocialIDResult
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// UserSocialAccountLoginsUpdate adds social account for a user.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/authentication/usersocialaccountlogins/update
//
// SDK Method Authentication::updateSocialLogin
func (s *Service) UserSocialAccountLoginsUpdate(userID string, userSocialAccountLogins *model.SocialAccountInformation) *UserSocialAccountLoginsUpdateOp {
	return &UserSocialAccountLoginsUpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"users", userID, "social"}, "/"),
		Payload:    userSocialAccountLogins,
		QueryOpts:  make(url.Values),
	}
}

// UserSocialAccountLoginsUpdateOp implements DocuSign API SDK Authentication::updateSocialLogin
type UserSocialAccountLoginsUpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UserSocialAccountLoginsUpdateOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}
