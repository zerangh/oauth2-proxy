package options

import (
	"fmt"
	"net/url"
	"time"
)

// Options holds all configuration options for the OAuth2 Proxy.
type Options struct {
	// Server configuration
	HTTPAddress  string `mapstructure:"http-address"`
	HTTPSAddress string `mapstructure:"https-address"`

	// TLS configuration
	TLSCertFile string `mapstructure:"tls-cert-file"`
	TLSKeyFile  string `mapstructure:"tls-key-file"`

	// Upstream configuration
	Upstreams []string `mapstructure:"upstreams"`

	// OAuth2 provider configuration
	Provider     string `mapstructure:"provider"`
	ClientID     string `mapstructure:"client-id"`
	ClientSecret string `mapstructure:"client-secret"`

	// Redirect URL after authentication
	RedirectURL string `mapstructure:"redirect-url"`

	// Cookie configuration
	CookieName     string        `mapstructure:"cookie-name"`
	CookieSecret   string        `mapstructure:"cookie-secret"`
	CookieDomain   string        `mapstructure:"cookie-domain"`
	CookiePath     string        `mapstructure:"cookie-path"`
	CookieExpire   time.Duration `mapstructure:"cookie-expire"`
	CookieRefresh  time.Duration `mapstructure:"cookie-refresh"`
	CookieSecure   bool          `mapstructure:"cookie-secure"`
	CookieHTTPOnly bool          `mapstructure:"cookie-httponly"`
	CookieSameSite string        `mapstructure:"cookie-samesite"`

	// Session configuration
	SessionStoreType string `mapstructure:"session-store-type"`

	// Email/domain restrictions
	EmailDomains      []string `mapstructure:"email-domains"`
	AuthenticatedEmailsFile string `mapstructure:"authenticated-emails-file"`

	// Allowed groups
	AllowedGroups []string `mapstructure:"allowed-groups"`

	// Proxy behavior
	PassAccessToken  bool `mapstructure:"pass-access-token"`
	PassBasicAuth    bool `mapstructure:"pass-basic-auth"`
	PassUserHeaders  bool `mapstructure:"pass-user-headers"`
	PassHostHeader   bool `mapstructure:"pass-host-header"`
	SetXAuthRequest  bool `mapstructure:"set-xauthrequest"`
	SetAuthorization bool `mapstructure:"set-authorization-header"`

	// Logging
	RequestLogging bool   `mapstructure:"request-logging"`
	LoggingFilename string `mapstructure:"logging-filename"`

	// Skip authentication for certain paths
	SkipAuthRegex  []string `mapstructure:"skip-auth-regex"`
	SkipAuthRoutes []string `mapstructure:"skip-auth-routes"`

	// Reverse proxy
	ReverseProxy bool `mapstructure:"reverse-proxy"`

	// Ping path for health checks
	PingPath       string `mapstructure:"ping-path"`
	PingUserAgent  string `mapstructure:"ping-user-agent"`
	ReadyPath      string `mapstructure:"ready-path"`
}

// NewOptions returns a new Options struct with default values applied.
func NewOptions() *Options {
	return &Options{
		HTTPAddress:    "127.0.0.1:4180",
		HTTPSAddress:   ":443",
		Provider:       "google",
		CookieName:     "_oauth2_proxy",
		CookiePath:     "/",
		CookieExpire:   168 * time.Hour,
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: "",
		SessionStoreType: "cookie",
		PassBasicAuth:  true,
		PassUserHeaders: true,
		PassHostHeader: true,
		RequestLogging: true,
		PingPath:       "/ping",
		ReadyPath:      "/ready",
		ReverseProxy:   false,
	}
}

// Validate checks that all required options are set and valid.
func (o *Options) Validate() error {
	if o.ClientID == "" {
		return fmt.Errorf("missing required option: client-id")
	}
	if o.ClientSecret == "" {
		return fmt.Errorf("missing required option: client-secret")
	}
	if o.CookieSecret == "" {
		return fmt.Errorf("missing required option: cookie-secret")
	}
	if len(o.Upstreams) == 0 {
		return fmt.Errorf("missing required option: upstreams")
	}

	// Validate redirect URL if provided
	if o.RedirectURL != "" {
		if _, err := url.Parse(o.RedirectURL); err != nil {
			return fmt.Errorf("invalid redirect-url %q: %w", o.RedirectURL, err)
		}
	}

	// Validate cookie same-site value
	switch o.CookieSameSite {
	case "", "none", "lax", "strict":
		// valid values
	default:
		return fmt.Errorf("invalid cookie-samesite value %q: must be one of '', 'none', 'lax', or 'strict'", o.CookieSameSite)
	}

	return nil
}
