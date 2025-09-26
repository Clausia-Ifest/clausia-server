package bootstrap

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/alecthomas/kingpin/v2"
	"github.com/rs/zerolog/log"
)

func (app *App) handleFlags() error {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate)

	m := kingpin.Flag("migrate", "Run database migration. (up|down|fresh)").
		Enum("up", "down", "fresh")
	s := kingpin.Flag("seed", "Run database seeder to add sample data.").Bool()
	f := kingpin.Flag("force", "Force run for migrate and/or seed").Short('f').Bool()

	kingpin.Parse()

	migrate := *m
	seed := *s
	needsConfirm := (slices.Contains([]string{"prod", "production"}, strings.ToLower(app.config.AppEnv))) && !*f

	if (slices.Contains([]string{"down", "fresh"}, migrate) || seed) && needsConfirm {
		fmt.Print("Are you sure you want to apply to production? [y/N] ")
		var response string
		_, _ = fmt.Scanln(&response)
		response = strings.ToLower(strings.TrimSpace(response))

		if response != "y" {
			var prefix string

			if migrate != "" {
				prefix += "[MIGRATE]"
			}

			if seed {
				prefix += "[SEEDER]"
			}

			log.Info().
				Msgf("%s Cancelled. No changes applied", prefix)

			os.Exit(0)
		}
	}

	if migrate != "" {
		if err := app.db.Migrate(migrate); err != nil {
			return err
		}

		if !seed || migrate == "down" {
			os.Exit(0)
		}
	}

	if seed {
		app.db.Seed()

		os.Exit(0)
	}

	return nil
}
