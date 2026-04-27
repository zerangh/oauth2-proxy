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
		// Reduced from 168h (7 days) to 24h for shorter-lived sessions; adjust as needed
		CookieExpire:   24 * time.Hour,
		// Refresh cookie 15 minutes before expiry to avoid abrupt session drops
		CookieRefresh:  23*time.Hour + 45*time.Minute,
		CookieSecure:   true,
		CookieHTTPOnly: true,
		// Default to Lax to balance CSRF protection with usability for top-level navigations
		CookieSameSite: "lax",
		PingPath:       "/ping",
		ReadyPath:      "/ready",
		RequestLogging: true,
	}
}
