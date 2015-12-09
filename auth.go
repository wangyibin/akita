package akita

import "database/sql"

// AuthConfig holds configuration parameters used when authenticating a user and

// creating a secure cookie user session.
type AuthConfig struct {
	Secret               []byte
	db                   sql.DB
	LoginRedirect        string
	LoginSuccessRedirect string
}

// Config is the default implementation of Config, and is used by
// DetaultAuthCallback, Secure, and SecureFunc.
var Config = &AuthConfig{
	LoginRedirect:        "/auth/login",
	LoginSuccessRedirect: "/",
}

// Passes back the OAuth Token. This will likely be the oauth2.Token or the
// oauth1.AccessToken... will need to cast to the appropriate value if you
// need specific fields (for now).
type Token interface {
	Token() string
}

// A User is returned by the AuthProvider upon success authentication.
type User interface {
	Id() string       // Unique identifier of the user
	Provider() string // Name of the Authentication Provider (ie google, github)
	Name() string     // Name of the User (ie lastname, firstname)
	Email() string    // Email Address of the User
	Org() string      // Company or Organization the User belongs to
	Picture() string  // URL of the User's Profile picture
	Link() string     // URL of the User's Profile page
}

// An implementation of User, for internal package use only.
type user struct {
	id       string
	provider string
	name     string
	email    string
	org      string
	link     string
	picture  string
}

func (u *user) Id() string       { return u.id }
func (u *user) Provider() string { return u.provider }
func (u *user) Name() string     { return u.name }
func (u *user) Email() string    { return u.email }
func (u *user) Org() string      { return u.org }
func (u *user) Link() string     { return u.link }
func (u *user) Picture() string  { return u.picture }
func (u *user) Avatar() string   { return u.picture }
