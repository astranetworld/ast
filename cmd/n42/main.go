// Copyright 2022 The N42 Authors
// This file is part of the N42 library.
//
// The N42 library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The N42 library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the N42 library. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"

	"github.com/n42blockchain/N42/params"
	// Force-load the tracer engines to trigger registration
	_ "github.com/n42blockchain/N42/internal/tracers/js"
	_ "github.com/n42blockchain/N42/internal/tracers/native"
)

func main() {
	fmt.Printf("┏━┓┏━┓╺┳╸┏━┓┏━┓┏┓╻┏━╸╺┳╸\n")
	fmt.Printf("┣━┫┗━┓ ┃ ┣┳┛┣━┫┃┗┫┣╸  ┃\n")
	fmt.Printf("╹ ╹┗━┛ ╹ ╹┗╸╹ ╹╹ ╹┗━╸ ╹\n")
	flags := append(networkFlags, consensusFlag...)
	flags = append(flags, loggerFlag...)
	flags = append(flags, pprofCfg...)
	flags = append(flags, nodeFlg...)
	flags = append(flags, rpcFlags...)
	flags = append(flags, authRPCFlag...)
	flags = append(flags, configFlag...)
	flags = append(flags, settingFlag...)
	flags = append(flags, accountFlag...)
	flags = append(flags, metricsFlags...)
	flags = append(flags, p2pFlags...)
	flags = append(flags, p2pLimitFlags...)

	rootCmd = append(rootCmd, walletCommand, accountCommand, exportCommand, initCommand)
	commands := rootCmd

	app := &cli.App{
		Name:     "ast",
		Usage:    "N42 system",
		Flags:    flags,
		Commands: commands,
		//Version:                version.FormatVersion(),
		Version:                params.VersionWithCommit(params.GitCommit, ""),
		UseShortOptionHandling: true,
		Action:                 appRun,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("failed ast system setup %v", err)
	}
}
