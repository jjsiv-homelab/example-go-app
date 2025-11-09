package commands

import (
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	ready           bool
	flagPort        string
	flagInsecure    bool
	flagTLSCertPath string
	flagTLSKeyPath  string
)

func initApp() {
	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if !ready {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})

	log.Info().Msg("initializing app...")
	time.Sleep(15 * time.Second)
	log.Info().Msg("app is ready")
	ready = true
}

func ServeCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "Start the HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello, World!"))
			})

			http.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			go initApp()

			log.Info().Msgf("server is starting on port %s", flagPort)
			if flagInsecure {
				log.Warn().Msg("running in insecure mode")
				if err := http.ListenAndServe(":"+flagPort, nil); err != nil {
					log.Error().Err(err).Msg("server failed to start")
					os.Exit(1)
				}
			} else {
				if flagTLSCertPath == "" || flagTLSKeyPath == "" {
					log.Fatal().Msg("TLS key and cert path must be specified!")
				}
				if err := http.ListenAndServeTLS(":"+flagPort, flagTLSCertPath, flagTLSKeyPath, nil); err != nil {
					log.Error().Err(err).Msg("server failed to start")
					os.Exit(1)
				}
			}
		},
	}

	command.PersistentFlags().StringVar(&flagPort, "port", "8443", "Port to listen on")
	command.PersistentFlags().BoolVar(&flagInsecure, "insecure", false, "Run in HTTP mode")
	command.PersistentFlags().StringVar(&flagTLSCertPath, "tls-cert-path", "", "Path to the TLS certificate")
	command.PersistentFlags().StringVar(&flagTLSKeyPath, "tls-key-path", "", "Path to the TLS key")

	return command
}
