package options

// Provider holds all configuration for a single authentication provider.
type Provider struct {
	// ID is a unique identifier for this provider instance.
	// Defaults to the provider type if not specified.
	ID string `flag:"provider-id" cfg:"provider_id"`

	// Type is the OAuth2 provider type (e.g. google, github, oidc, azure).
	Type ProviderType `flag:"provider" cfg:"provider"`

	// Name is the human-readable name displayed on the login button.
	Name string `flag:"provider-display-name" cfg:"provider_display_name"`

	// ClientID is the OAuth2 client ID registered with the provider.
	ClientID string `flag:"client-id" cfg:"client_id"`

	// ClientSecret is the OAuth2 client secret registered with the provider.
	ClientSecret string `flag:"client-secret" cfg:"client_secret"`

	// ClientSecretFile is a path to a file containing the client secret.
	// Takes precedence over ClientSecret if set.
	ClientSecretFile string `flag:"client-secret-file" cfg:"client_secret_file"`

	// CAFiles is a list of paths to CA certificate files used to validate
	// the provider's TLS certificate.
	CAFiles []string `flag:"provider-ca-file" cfg:"provider_ca_files"`

	// LoginURL is the authentication endpoint URL.
	LoginURL string `flag:"login-url" cfg:"login_url"`

	// RedeemURL is the token redemption endpoint URL.
	RedeemURL string `flag:"redeem-url" cfg:"redeem_url"`

	// ProfileURL is the user profile endpoint URL.
	ProfileURL string `flag:"profile-url" cfg:"profile_url"`

	// ProtectedResource is the resource that requires protection (used by
	// some providers such as Azure AD).
	ProtectedResource string `flag:"resource" cfg:"resource"`

	// ValidateURL is the endpoint used to validate access tokens.
	ValidateURL string `flag:"validate-url" cfg:"validate_url"`

	// Scope is the OAuth2 scope to request from the provider.
	Scope string `flag:"scope" cfg:"scope"`

	// Prompt specifies the prompt parameter sent to the provider.
	Prompt string `flag:"prompt" cfg:"prompt"`

	// ApprovalPrompt is the legacy prompt parameter (deprecated, use Prompt).
	ApprovalPrompt string `flag:"approval-prompt" cfg:"approval_prompt"`

	// AcrValues specifies the Authentication Context Class Reference values.
	AcrValues string `flag:"acr-values" cfg:"acr_values"`

	// AllowedGroups is a list of restrict logins to members of these groups.
	AllowedGroups []string `flag:"allowed-group" cfg:"allowed_groups"`

	// OIDC holds configuration specific to OpenID Connect providers.
	OIDCConfig OIDCOptions `cfg:",squash"`
}

// ProviderType represents the type of OAuth2 provider.
type ProviderType string

const (
	ProviderTypeGoogle       ProviderType = "google"
	ProviderTypeGitHub       ProviderType = "github"
	ProviderTypeGitLab       ProviderType = "gitlab"
	ProviderTypeOIDC         ProviderType = "oidc"
	ProviderTypeAzure        ProviderType = "azure"
	ProviderTypeAzureAD      ProviderType = "azuread"
	ProviderTypeBitbucket    ProviderType = "bitbucket"
	ProviderTypeDigitalOcean ProviderType = "digitalocean"
	ProviderTypeFacebook     ProviderType = "facebook"
	ProviderTypeKeycloak     ProviderType = "keycloak"
	ProviderTypeLinkedIn     ProviderType = "linkedin"
)

// OIDCOptions holds configuration specific to OpenID Connect providers.
type OIDCOptions struct {
	// IssuerURL is the OpenID Connect issuer URL.
	IssuerURL string `flag:"oidc-issuer-url" cfg:"oidc_issuer_url"`

	// InsecureAllowUnverifiedEmail allows users with unverified email addresses
	// to authenticate. Use with caution.
	InsecureAllowUnverifiedEmail bool `flag:"insecure-oidc-allow-unverified-email" cfg:"insecure_oidc_allow_unverified_email"`

	// InsecureSkipIssuerVerification skips verification of the issuer claim.
	// Use with caution in development environments only.
	InsecureSkipIssuerVerification bool `flag:"insecure-oidc-skip-issuer-verification" cfg:"insecure_oidc_skip_issuer_verification"`

	// SkipDiscovery disables OIDC discovery and requires manual endpoint
	// configuration via LoginURL, RedeemURL, and JwksURL.
	SkipDiscovery bool `flag:"skip-oidc-discovery" cfg:"skip_oidc_discovery"`

	// JwksURL is the URL of the JSON Web Key Set used to verify ID tokens
	// when SkipDiscovery is enabled.
	JwksURL string `flag:"oidc-jwks-url" cfg:"oidc_jwks_url"`

	// EmailClaim is the claim used to extract the user's email address.
	// Defaults to "email".
	EmailClaim string `flag:"oidc-email-claim" cfg:"oidc_email_claim"`

	// GroupsClaim is the claim used to extract the user's groups.
	// Defaults to "groups".
	GroupsClaim string `flag:"oidc-groups-claim" cfg:"oidc_groups_claim"`

	// AudienceClaims is a list of claims checked for the client ID audience.
	AudienceClaims []string `flag:"oidc-audience-claim" cfg:"oidc_audience_claims"`

	// ExtraAudiences is a list of additional audiences that are allowed
	// beyond the client ID.
	ExtraAudiences []string `flag:"oidc-extra-audience" cfg:"oidc_extra_audiences"`
}

// NewProvider returns a Provider with sensible defaults applied.
func NewProvider() Provider {
	return Provider{
		Type:           ProviderTypeGoogle,
		ApprovalPrompt: "force",
		OIDCConfig: OIDCOptions{
			EmailClaim:     "email",
			GroupsClaim:    "groups",
			AudienceClaims: []string{"aud"},
		},
	}
}
