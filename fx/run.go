package fx

import (
	gifx "github.com/americanas-go/ignite/go.uber.org/fx.v1"
	"go.uber.org/fx"
)

// Run executes the fx app.
func Run(options fx.Option) error {
	app := gifx.NewApp(options)
	if app.Err() != nil {
		return app.Err()
	}
	app.Run()
	return nil
}
