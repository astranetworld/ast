// Copyright 2022 The astranet Authors
// This file is part of the astranet library.
//
// The astranet library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The astranet library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the astranet library. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"

	"github.com/astranetworld/ast/params"
	// Force-load the tracer engines to trigger registration
	_ "github.com/astranetworld/ast/internal/tracers/js"
	_ "github.com/astranetworld/ast/internal/tracers/native"
)

func main() {
	fmt.Printf("\u001B[0;1;35;95m┏━\u001B[0;1;31;91m┓┏\u001B[0;1;33;93m━┓\u001B[0;1;32;92m╺┳\u001B[0;1;36;96m╸┏\u001B[0;1;34;94m━┓\u001B[0;1;35;95m┏━\u001B[0;1;31;91m┓┏\u001B[0;1;33;93m┓╻\u001B[0;1;32;92m┏━\u001B[0;1;36;96m╸╺\u001B[0;1;34;94m┳╸\u001B[0m\n")
	fmt.Printf("\u001B[0;1;31;91m┣━\u001B[0;1;33;93m┫┗\u001B[0;1;32;92m━┓\u001B[0m \u001B[0;1;36;96m┃\u001B[0m \u001B[0;1;34;94m┣\u001B[0;1;35;95m┳┛\u001B[0;1;31;91m┣━\u001B[0;1;33;93m┫┃\u001B[0;1;32;92m┗┫\u001B[0;1;36;96m┣╸\u001B[0m  \u001B[0;1;35;95m┃\u001B[0m \n")
	fmt.Printf("\u001B[0;1;33;93m╹\u001B[0m \u001B[0;1;32;92m╹┗\u001B[0;1;36;96m━┛\u001B[0m \u001B[0;1;34;94m╹\u001B[0m \u001B[0;1;35;95m╹\u001B[0;1;31;91m┗╸\u001B[0;1;33;93m╹\u001B[0m \u001B[0;1;32;92m╹╹\u001B[0m \u001B[0;1;36;96m╹\u001B[0;1;34;94m┗━\u001B[0;1;35;95m╸\u001B[0m \u001B[0;1;31;91m╹\u001B[0m \n")
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
		Usage:    "astranet system",
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
