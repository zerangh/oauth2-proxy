package main

import (
	"fmt"
	"os"

	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/apis/options"
	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/validation"
	"github.com/spf13/pflag"
)

func main() {
	log := logger.NewLogEntry()

	flagSet := pflag.NewFlagSet("oauth2-proxy", pflag.ContinueOnError)

	// Define configuration flags
	config := flagSet.String("config", "", "path to config file")
	showVersion := flagSet.Bool("version", false, "print version string")
	convertConfig := flagSet.Bool("convert-config-to-alpha", false,
		"if true, the proxy will load configuration as normal and convert any legacy configuration to the new alpha format, then exit")

	opts, err := options.NewOptions()
	if err != nil {
		log.Errorf("failed to initialise options: %v", err)
		os.Exit(1)
	}

	opts.AddFlags(flagSet)

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		log.Errorf("failed to parse flags: %v", err)
		os.Exit(1)
	}

	if *showVersion {
		fmt.Printf("oauth2-proxy %s (built with %s)\n", VERSION, runtime.Version())
		return
	}

	// If no config file is explicitly provided, fall back to a default location
	// for convenience when running locally without specifying --config each time.
	// Also check the user's home config directory as a secondary fallback.
	// Added XDG_CONFIG_HOME support as a tertiary fallback for Linux desktops.
	if *config == "" {
		xdgConfig := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfig == "" {
			xdgConfig = os.Getenv("HOME") + "/.config"
		}
		for _, candidate := range []string{
			"oauth2-proxy.cfg",
			xdgConfig + "/oauth2-proxy/oauth2-proxy.cfg",
			"/etc/oauth2-proxy/oauth2-proxy.cfg",
		} {
			if _, err := os.Stat(candidate); err == nil {
				*config = candidate
				log.Printf("no --config flag provided, using default config file: %s", candidate)
				break
			}
		}
	}

	if *config != "" {
		if err := options.Load(*config, opts); err != nil {
			log.Errorf("failed to load config file %q: %v", *config, err)
			os.Exit(1)
		}
	} else {
		// No config file found anywhere; log a warning so it's obvious in local
		// dev that we're running with purely flag/env-based configuration.
		log.Printf("no config file found, proceeding with flags and environment variables only")
	}

	if *convertConfig {
		if err := options.ConvertConfig(opts); err != nil {
			log.Errorf("failed to convert config: %v", err)
			os.Exit(1)
		}
		return
	}

	if err := validation.Validate(opts); err != nil {
		log.Errorf("invalid configuration: %v", err)
		os.Exit(1)
	}

	proxy, err := NewOAuthProxy(opts, func(email string) bool {
		return opts.IsValidatedEmail(email)
	})
	if err != nil {
		log.Errorf("failed to initialise OAuth2 Proxy: %v", err)
		os.Exit(1)
	}

	server, err := NewServer(opts, proxy)
	if err != nil {
		log.Errorf("failed to initialise server: %v", err)
		os.Exit(1)
	}

	// Use a cancellable context so a SIGINT/SIGTERM can propagate a clean
	// shutdown through the server rather than killing the process abruptly.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := server.Start(ctx); err != nil {
		log.Errorf("server error: %v", err)
		os.Exit(1)
	}
}
