// Copyright 2016 Marapongo, Inc. All rights reserved.

package cmdutil

import (
	"fmt"
	"os"
	"strings"

	"github.com/marapongo/mu/pkg/compiler/backends"
	"github.com/marapongo/mu/pkg/compiler/backends/clouds"
	"github.com/marapongo/mu/pkg/compiler/backends/schedulers"
	"github.com/marapongo/mu/pkg/options"
)

func SetCloudArchOptions(arch string, opts *options.Options) {
	// If an architecture was specified, parse the pieces and set the options.  This isn't required because stacks
	// and workspaces can have defaults.  This simply overrides or provides one where none exists.
	if arch != "" {
		// The format is "cloud[:scheduler]"; parse out the pieces.
		var cloud string
		var scheduler string
		if delim := strings.IndexRune(arch, ':'); delim != -1 {
			cloud = arch[:delim]
			scheduler = arch[delim+1:]
		} else {
			cloud = arch
		}

		cloudArch, ok := clouds.Values[cloud]
		if !ok {
			fmt.Fprintf(os.Stderr, "Unrecognized cloud arch '%v'\n", cloud)
			os.Exit(-1)
		}

		var schedulerArch schedulers.Arch
		if scheduler != "" {
			schedulerArch, ok = schedulers.Values[scheduler]
			if !ok {
				fmt.Fprintf(os.Stderr, "Unrecognized cloud scheduler arch '%v'\n", scheduler)
				os.Exit(-1)
			}
		}

		opts.Arch = backends.Arch{
			Cloud:     cloudArch,
			Scheduler: schedulerArch,
		}
	}
}
