package gocloak

import (
	"encoding/json"
	"strconv"
)

// BaseParams provides basic functionality for all QueryParams structures.
// The fields tags must have `json:"<name>,string,omitempty"` format.
// "string" tag allows to convert the structure to map[string]string.
// "omitempty" allows to skip the fields with default values.
type BaseParams struct{}

// GetQueryParams converts the struct to map[string]string
func (s BaseParams) GetQueryParams() (map[string]string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	var res map[string]string
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// APIError represents an api error
type APIError struct {
	Code    int
	Message string
}

// Error stringifies the APIError
func (apiError APIError) Error() string {
	return apiError.Message
}

// CertResponseKey is returned by the certs endpoint
type CertResponseKey struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}

// CertResponse is returned by the certs endpoint
type CertResponse struct {
	Keys []CertResponseKey `json:"keys"`
}

// IssuerResponse is returned by the issuer endpoint
type IssuerResponse struct {
	Realm           string `json:"realm"`
	PublicKey       string `json:"public_key"`
	TokenService    string `json:"token-service"`
	AccountService  string `json:"account-service"`
	TokensNotBefore int    `json:"tokens-not-before"`
}

// RetrospecTokenResult is returned when a token was checked
type RetrospecTokenResult struct {
	Permissions map[string]string `json:"permissions,omitempty"`
	Exp         int               `json:"exp,omitempty"`
	Nbf         int               `json:"nbf,omitempty"`
	Iat         int               `json:"iat,omitempty"`
	Aud         string            `json:"aud,omitempty"`
	Active      bool              `json:"active,omitempty"`
	AuthTime    int               `json:"auth_time,omitempty"`
	Jti         string            `json:"jti,omitempty"`
	Type        string            `json:"typ,omitempty"`
}

// User represents the Keycloak User Structure
type User struct {
	ID                         string              `json:"id,omitempty"`
	CreatedTimestamp           int64               `json:"createdTimestamp,omitempty"`
	Username                   string              `json:"username,omitempty"`
	Enabled                    bool                `json:"enabled,omitempty"`
	Totp                       bool                `json:"totp,omitempty"`
	EmailVerified              bool                `json:"emailVerified,omitempty"`
	FirstName                  string              `json:"firstName,omitempty"`
	LastName                   string              `json:"lastName,omitempty"`
	Email                      string              `json:"email,omitempty"`
	FederationLink             string              `json:"federationLink,omitempty"`
	Attributes                 map[string][]string `json:"attributes,omitempty"`
	DisableableCredentialTypes []interface{}       `json:"disableableCredentialTypes,omitempty"`
	RequiredActions            []interface{}       `json:"requiredActions,omitempty"`
	Access                     map[string]bool     `json:"access,omitempty"`
}

// SetPasswordRequest sets a new password
type SetPasswordRequest struct {
	Type      string `json:"type"`
	Temporary bool   `json:"temporary"`
	Password  string `json:"value"`
}

// Component is a component
type Component struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	ProviderID      string          `json:"providerId"`
	ProviderType    string          `json:"providerType"`
	ParentID        string          `json:"parentId"`
	ComponentConfig ComponentConfig `json:"config"`
	SubType         string          `json:"subType,omitempty"`
}

// ComponentConfig is a componentconfig
type ComponentConfig struct {
	Priority  []string `json:"priority"`
	Algorithm []string `json:"algorithm"`
}

// KeyStoreConfig holds the keyStoreConfig
type KeyStoreConfig struct {
	ActiveKeys ActiveKeys `json:"active"`
	Key        []Key      `json:"keys"`
}

// ActiveKeys holds the active keys
type ActiveKeys struct {
	HS256 string `json:"HS256"`
	RS256 string `json:"RS256"`
	AES   string `json:"AES"`
}

// Key is a key
type Key struct {
	ProviderID       string `json:"providerId"`
	ProviderPriority int    `json:"providerPriority"`
	Kid              string `json:"kid"`
	Status           string `json:"status"`
	Type             string `json:"type"`
	Algorithm        string `json:"algorithm"`
	PublicKey        string `json:"publicKey,omitempty"`
	Certificate      string `json:"certificate,omitempty"`
}

// Attributes holds Attributes
type Attributes struct {
	LDAPENTRYDN []string `json:"LDAP_ENTRY_DN"`
	LDAPID      []string `json:"LDAP_ID"`
}

// Access represents access
type Access struct {
	ManageGroupMembership bool `json:"manageGroupMembership"`
	View                  bool `json:"view"`
	MapRoles              bool `json:"mapRoles"`
	Impersonate           bool `json:"impersonate"`
	Manage                bool `json:"manage"`
}

// UserGroup is a UserGroup
type UserGroup struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

// GetUsersParams represents the optional parameters for getting users
type GetUsersParams struct {
	BaseParams
	BriefRepresentation *bool  `json:"briefRepresentation,string,omitempty"`
	Email               string `json:"email,string,omitempty"`
	First               int    `json:"first,string,omitempty"`
	FirstName           string `json:"firstName,string,omitempty"`
	LastName            string `json:"lastName,string,omitempty"`
	Max                 int    `json:"max,string,omitempty"`
	Search              string `json:"search,string,omitempty"`
	Username            string `json:"username,string,omitempty"`
}

// GetQueryParams converts the struct to map[string]string
func (s GetUsersParams) GetQueryParams() (map[string]string, error) {
	var res map[string]string

	res["briefRepresentation"] = strconv.FormatBool(*s.BriefRepresentation)
	res["email"] = s.Email
	res["first"] = strconv.FormatInt(int64(s.First), 10)
	res["firstName"] = s.FirstName
	res["lastName"] = s.LastName
	res["max"] = strconv.FormatInt(int64(s.Max), 10)
	res["search"] = s.Search
	res["username"] = s.Username

	return res, nil
}

// ExecuteActionsEmail represents parameters for executing action emails
type ExecuteActionsEmail struct {
	BaseParams
	UserID      string   `json:"-"`
	ClientID    string   `json:"clientId,string,omitempty"`
	Lifespan    int      `json:"lifespan,string,omitempty"`
	RedirectURI string   `json:"redirect_uri,string,omitempty"`
	Actions     []string `json:"-"`
}

// Group is a Group
type Group struct {
	ID        string        `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
	Path      string        `json:"path,omitempty"`
	SubGroups []interface{} `json:"subGroups,omitempty"`
}

// GetGroupsParams represents the optional parameters for getting groups
type GetGroupsParams struct {
	BaseParams
	First  int    `json:"first,string,omitempty"`
	Max    int    `json:"max,string,omitempty"`
	Search string `json:"search,string,omitempty"`
}

// Role is a role
type Role struct {
	ID                 string              `json:"id,omitempty"`
	Name               string              `json:"name,omitempty"`
	ScopeParamRequired bool                `json:"scopeParamRequired,omitempty"`
	Composite          bool                `json:"composite,omitempty"`
	ClientRole         bool                `json:"clientRole,omitempty"`
	ContainerID        string              `json:"containerId,omitempty"`
	Description        string              `json:"description,omitempty"`
	Attributes         map[string][]string `json:"attributes,omitempty"`
}

// ClientMappingsRepresentation is a client role mappings
type ClientMappingsRepresentation struct {
	ID       string `json:"id"`
	Client   string `json:"client"`
	Mappings []Role `json:"mappings"`
}

// MappingsRepresentation is a representation of role mappings
type MappingsRepresentation struct {
	ClientMappings map[string]ClientMappingsRepresentation `json:"clientMappings,omitempty"`
	RealmMappings  []Role                                  `json:"realmMappings,omitempty"`
}

// ClientScope is a ClientScope
type ClientScope struct {
	ID                    string                `json:"id"`
	Name                  string                `json:"name"`
	Description           string                `json:"description"`
	Protocol              string                `json:"protocol"`
	ClientScopeAttributes ClientScopeAttributes `json:"attributes"`
	ProtocolMappers       ProtocolMappers       `json:"protocolMappers,omitempty"`
}

// ClientScopeAttributes are attributes of client scopes
type ClientScopeAttributes struct {
	ConsentScreenText      string `json:"consent.screen.text"`
	DisplayOnConsentScreen string `json:"display.on.consent.screen"`
}

// ProtocolMappers are protocolmappers
type ProtocolMappers struct {
	ID                    string                `json:"id"`
	Name                  string                `json:"name"`
	Protocol              string                `json:"protocol"`
	ProtocolMapper        string                `json:"protocolMapper"`
	ConsentRequired       bool                  `json:"consentRequired"`
	ProtocolMappersConfig ProtocolMappersConfig `json:"config"`
}

// ProtocolMappersConfig is a config of a protocol mapper
type ProtocolMappersConfig struct {
	UserinfoTokenClaim string `json:"userinfo.token.claim"`
	UserAttribute      string `json:"user.attribute"`
	IDTokenClaim       string `json:"id.token.claim"`
	AccessTokenClaim   string `json:"access.token.claim"`
	ClaimName          string `json:"claim.name"`
	JSONTypeLabel      string `json:"jsonType.label"`
}

// Client is a Client
type Client struct {
	ID       string `json:"id"`
	ClientID string `json:"clientId"`
}

// GetClientsParams represents the query parameters
type GetClientsParams struct {
	BaseParams
	ClientID     string `json:"clientId,string,omitempty"`
	ViewableOnly bool   `json:"viewableOnly,string,omitempty"`
}

// UserInfo is returned by the userinfo endpoint
type UserInfo struct {
	Sub               string      `json:"sub"`
	EmailVerified     bool        `json:"email_verified"`
	Address           interface{} `json:"address"`
	PreferredUsername string      `json:"preferred_username"`
	Email             string      `json:"email"`
}

// RealmRepresentation represent a realm
type RealmRepresentation struct {
	AccessCodeLifespan                  int               `json:"accessCodeLifespan"`
	AccessCodeLifespanLogin             int               `json:"accessCodeLifespanLogin"`
	AccessCodeLifespanUserAction        int               `json:"accessCodeLifespanUserAction"`
	AccessTokenLifespan                 int               `json:"accessTokenLifespan"`
	AccessTokenLifespanForImplicitFlow  int               `json:"accessTokenLifespanForImplicitFlow"`
	AccountTheme                        string            `json:"accountTheme"`
	ActionTokenGeneratedByAdminLifespan int               `json:"actionTokenGeneratedByAdminLifespan"`
	ActionTokenGeneratedByUserLifespan  int               `json:"actionTokenGeneratedByUserLifespan"`
	AdminEventsDetailsEnabled           bool              `json:"adminEventsDetailsEnabled"`
	AdminEventsEnabled                  bool              `json:"adminEventsEnabled"`
	AdminTheme                          string            `json:"adminTheme"`
	Attributes                          map[string]string `json:"attributes"`
	AuthenticationFlows                 []interface{}     `json:"authenticationFlows"`
	AuthenticatorConfig                 []interface{}     `json:"authenticatorConfig"`
	BrowserFlow                         string            `json:"browserFlow"`
	BrowserSecurityHeaders              map[string]string `json:"browserSecurityHeaders"`
	BruteForceProtected                 bool              `json:"bruteForceProtected"`
	ClientAuthenticationFlow            string            `json:"clientAuthenticationFlow"`
	ClientScopeMappings                 map[string]string `json:"clientScopeMappings"`
	ClientScopes                        []interface{}     `json:"clientScopes"`
	Clients                             []interface{}     `json:"clients"`
	Components                          interface{}       `json:"components"`
	DefaultDefaultClientScopes          []string          `json:"defaultDefaultClientScopes"`
	DefaultGroups                       []string          `json:"defaultGroups"`
	DefaultLocale                       string            `json:"defaultLocale"`
	DefaultOptionalClientScopes         []string          `json:"defaultOptionalClientScopes"`
	DefaultRoles                        []string          `json:"defaultRoles"`
	DefaultSignatureAlgorithm           string            `json:"defaultSignatureAlgorithm"`
	DirectGrantFlow                     string            `json:"directGrantFlow"`
	DisplayName                         string            `json:"displayName"`
	DisplayNameHTML                     string            `json:"displayNameHtml"`
	DockerAuthenticationFlow            string            `json:"dockerAuthenticationFlow"`
	DuplicateEmailsAllowed              bool              `json:"duplicateEmailsAllowed"`
	EditUsernameAllowed                 bool              `json:"editUsernameAllowed"`
	EmailTheme                          string            `json:"emailTheme"`
	Enabled                             bool              `json:"enabled"`
	EnabledEventTypes                   []string          `json:"enabledEventTypes"`
	EventsEnabled                       bool              `json:"eventsEnabled"`
	EventsExpiration                    int64             `json:"eventsExpiration"`
	EventsListeners                     []string          `json:"eventsListeners"`
	FailureFactor                       int               `json:"failureFactor"`
	FederatedUsers                      []interface{}     `json:"federatedUsers"`
	Groups                              []interface{}     `json:"groups"`
	ID                                  string            `json:"id"`
	IdentityProviderMappers             []interface{}     `json:"identityProviderMappers"`
	IdentityProviders                   []interface{}     `json:"identityProviders"`
	InternationalizationEnabled         bool              `json:"internationalizationEnabled"`
	KeycloakVersion                     string            `json:"keycloakVersion"`
	LoginTheme                          string            `json:"loginTheme"`
	LoginWithEmailAllowed               bool              `json:"loginWithEmailAllowed"`
	MaxDeltaTimeSeconds                 int               `json:"maxDeltaTimeSeconds"`
	MaxFailureWaitSeconds               int               `json:"maxFailureWaitSeconds"`
	MinimumQuickLoginWaitSeconds        int               `json:"minimumQuickLoginWaitSeconds"`
	NotBefore                           int               `json:"notBefore"`
	OfflineSessionIdleTimeout           int               `json:"offlineSessionIdleTimeout"`
	OfflineSessionMaxLifespan           int               `json:"offlineSessionMaxLifespan"`
	OfflineSessionMaxLifespanEnabled    bool              `json:"offlineSessionMaxLifespanEnabled"`
	OtpPolicyAlgorithm                  string            `json:"otpPolicyAlgorithm"`
	OtpPolicyDigits                     int               `json:"otpPolicyDigits"`
	OtpPolicyInitialCounter             int               `json:"otpPolicyInitialCounter"`
	OtpPolicyLookAheadWindow            int               `json:"otpPolicyLookAheadWindow"`
	OtpPolicyPeriod                     int               `json:"otpPolicyPeriod"`
	OtpPolicyType                       string            `json:"otpPolicyType"`
	OtpSupportedApplications            []string          `json:"otpSupportedApplications"`
	PasswordPolicy                      string            `json:"passwordPolicy"`
	PermanentLockout                    bool              `json:"permanentLockout"`
	ProtocolMappers                     []interface{}     `json:"protocolMappers"`
	QuickLoginCheckMilliSeconds         int64             `json:"quickLoginCheckMilliSeconds"`
	Realm                               string            `json:"realm"`
	RefreshTokenMaxReuse                int               `json:"refreshTokenMaxReuse"`
	RegistrationAllowed                 bool              `json:"registrationAllowed"`
	RegistrationEmailAsUsername         bool              `json:"registrationEmailAsUsername"`
	RegistrationFlow                    string            `json:"registrationFlow"`
	RememberMe                          bool              `json:"rememberMe"`
	RequiredActions                     []interface{}     `json:"requiredActions"`
	ResetCredentialsFlow                string            `json:"resetCredentialsFlow"`
	ResetPasswordAllowed                bool              `json:"resetPasswordAllowed"`
	RevokeRefreshToken                  bool              `json:"revokeRefreshToken"`
	Roles                               interface{}       `json:"roles"`
	ScopeMappings                       []interface{}     `json:"scopeMappings"`
	SMTPServer                          map[string]string `json:"smtpServer"`
	SslRequired                         string            `json:"sslRequired"`
	SsoSessionIdleTimeout               int               `json:"ssoSessionIdleTimeout"`
	SsoSessionIdleTimeoutRememberMe     int               `json:"ssoSessionIdleTimeoutRememberMe"`
	SsoSessionMaxLifespan               int               `json:"ssoSessionMaxLifespan"`
	SsoSessionMaxLifespanRememberMe     int               `json:"ssoSessionMaxLifespanRememberMe"`
	SupportedLocales                    []string          `json:"supportedLocales"`
	UserFederationMappers               []interface{}     `json:"userFederationMappers"`
	UserFederationProviders             []interface{}     `json:"userFederationProviders"`
	UserManagedAccessAllowed            bool              `json:"userManagedAccessAllowed"`
	Users                               []interface{}     `json:"users"`
	VerifyEmail                         bool              `json:"verifyEmail"`
	WaitIncrementSeconds                int               `json:"waitIncrementSeconds"`
}

// MultivaluedHashMap represents something
type MultivaluedHashMap struct {
	Empty      bool    `json:"empty"`
	LoadFactor float32 `json:"loadFactor"`
	Threshold  int32   `json:"threshold"`
}

// CredentialRepresentation represents credentials
type CredentialRepresentation struct {
	Algorithm         string             `json:"algorithm"`
	Config            MultivaluedHashMap `json:"config"`
	Counter           int32              `json:"counter"`
	CreatedDate       int64              `json:"createdDate"`
	Device            string             `json:"device"`
	Digits            int32              `json:"digits"`
	HashIterations    int32              `json:"hashIterations"`
	HashedSaltedValue string             `json:"hashedSaltedValue"`
	Period            int32              `json:"period"`
	Salt              string             `json:"salt"`
	Temporary         bool               `json:"temporary"`
	Type              string             `json:"type"`
	Value             string             `json:"value"`
}
