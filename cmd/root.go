/*
Copyright © 2025 blacktop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	verbose bool
	logger  *log.Logger
	// Version stores the service's version
	Version string
)

// CompilationDatabase represents the entire compilation database
type CompilationDatabase []CompileCommand

// CompileCommand represents a single compilation command for a source file
type CompileCommand struct {
	Directory string   `json:"directory"` // The working directory for the compilation
	File      string   `json:"file"`      // The source file path
	Output    string   `json:"output"`    // The output file (object file)
	Arguments []string `json:"arguments"` // The compiler arguments
}

func init() {
	// Override the default error level style.
	styles := log.DefaultStyles()
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("204")).
		Foreground(lipgloss.Color("0"))
	// Add a custom style for key `err`
	styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Values["err"] = lipgloss.NewStyle().Bold(true)
	logger = log.New(os.Stderr)
	logger.SetStyles(styles)

	// Define CLI flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose debug logging")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccpaths <compile_commands.json> <local_path> <db_path>",
	Short: "A tool to update paths in compile_commands.json",
	Long: `ccpaths is a command-line tool that helps you update paths in a compile_commands.json file.

It is particularly useful when you have a large number of compilation commands and need to adjust the paths for a specific target or output directory.`,
	Example: `ccpaths compile_commands.json /path/to/target /path/in/db

This command will read the compile_commands.json file, update the paths based on the target_path, and save the updated JSON in place.`,
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		if verbose {
			log.SetLevel(log.DebugLevel)
		}

		dbfile := filepath.Clean(args[0])

		// Read and parse the compile_commands.json file
		data, err := os.ReadFile(dbfile)
		if err != nil {
			return fmt.Errorf("failed to read compile commands file: %w", err)
		}

		var ccdb CompilationDatabase
		if err := json.Unmarshal(data, &ccdb); err != nil {
			return fmt.Errorf("failed to parse compile commands JSON: %w", err)
		}

		logger.Debug("Successfully parsed compile commands", "count", len(ccdb))

		// Process the second argument (likely the target path)
		targetPath := filepath.Clean(args[1])
		logger.Debug("Target path", "path", targetPath)

		// Now we have the compile commands loaded and ready to update paths
		for i, cc := range ccdb {
			ccdb[i].Directory = strings.Replace(cc.Directory, targetPath, "", -1)
			ccdb[i].File = strings.Replace(cc.File, targetPath, "", -1)
			for j, arg := range cc.Arguments {
				ccdb[i].Arguments[j] = strings.Replace(arg, targetPath, "", -1)
			}
		}

		updatedData, err := json.Marshal(ccdb)
		if err != nil {
			return fmt.Errorf("failed to marshal updated compile commands: %w", err)
		}
		if err := os.WriteFile(dbfile, updatedData, 0o644); err != nil {
			return fmt.Errorf("failed to write updated compile commands to file: %w", err)
		}
		log.Info("Successfully updated compile commands", "file", dbfile)

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Failed to execute command", "error", err)
	}
}
