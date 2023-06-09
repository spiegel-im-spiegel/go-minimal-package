package facade

import (
	"context"
	"os"
	"os/signal"
	"runtime"

	"github.com/goark/errs"
	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/go-minimal-package/ecode"
)

var (
	//Name is applicatin name
	Name = "go-minimal-package"
	//Version is version for applicatin
	Version = "dev-version"
)

var (
	debugFlag bool //debug flag
)

// newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use: Name,
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugPrint(ui, errs.Wrap(ecode.ErrNoCommand))
		},
	}
	// global options
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "", false, "for debug")

	rootCmd.SilenceUsage = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetArgs(args)
	rootCmd.SetIn(ui.Reader())       //Stdin
	rootCmd.SetOut(ui.ErrorWriter()) //Stdout -> Stderr
	rootCmd.SetErr(ui.ErrorWriter()) //Stderr
	rootCmd.AddCommand(
		newVersionCmd(ui),
	)
	return rootCmd
}

// Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			_ = ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, _, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				_ = ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ": line", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	// create interrupt SIGNAL
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).ExecuteContext(ctx); err != nil {
		exit = exitcode.Abnormal
	}
	return
}

/* Copyright 2023 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
