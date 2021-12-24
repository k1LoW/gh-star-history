/*
Copyright © 2021 Ken'ichiro Oyama <k1lowxb@gmail.com>

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
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v40/github"
	"github.com/k1LoW/gh-star-history/version"
	"github.com/k1LoW/go-github-client/v40/factory"
	"github.com/spf13/cobra"
	"github.com/zhangyunhao116/skipmap"
)

var (
	owner    string
	repos    []string
	perDay   bool
	perMonth bool
	perYear  bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "gh-star-history",
	Short:        "Show star history of repositories",
	Long:         `Show star history of repositories.`,
	Version:      version.Version,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		c, err := factory.NewGithubClient()
		if err != nil {
			return err
		}
		var f string
		if perDay {
			f = "2006-01-02"
		} else if perMonth {
			f = "2006-01"
		} else if perYear {
			f = "2006"
		} else {
			return errors.New("required to specify --per-day or --per-month or --per-year")
		}

		if len(repos) == 0 {
			page := 1
			for {
				rs, res, err := c.Repositories.List(ctx, owner, &github.RepositoryListOptions{
					Type: "public",
					ListOptions: github.ListOptions{
						Page:    page,
						PerPage: 100,
					},
				})
				if err != nil {
					return err
				}
				for _, r := range rs {
					repos = append(repos, r.GetName())
				}
				if res.NextPage == 0 {
					break
				}
				page += 1
			}
		}

		m := skipmap.NewInt64()

		for _, repo := range repos {
			cmd.PrintErrf("Aggregate stars of %s/%s ...\n", owner, repo)
			page := 1
			for {
				stars, res, err := c.Activity.ListStargazers(ctx, owner, repo, &github.ListOptions{
					Page:    page,
					PerPage: 100,
				})
				if err != nil {
					return err
				}
				for _, s := range stars {
					var k int64
					if f != "" {
						kt, err := time.Parse(f, s.GetStarredAt().UTC().Format(f))
						if err != nil {
							return err
						}
						k = kt.Unix()
					} else {
						k = -1
					}
					v, ok := m.Load(k)
					if ok {
						m.Store(k, v.(int)+1)
					} else {
						m.Store(k, 1)
					}
				}
				if res.NextPage == 0 {
					break
				}
				page += 1
			}
		}

		m.Range(func(k int64, v interface{}) bool {
			_, _ = fmt.Fprintf(os.Stdout, "%s\t%d\n", time.Unix(k, 0).Format(f), v.(int))
			return true
		})

		return nil
	},
}

func Execute() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stdin)

	log.SetOutput(io.Discard)
	if env := os.Getenv("DEBUG"); env != "" {
		log.SetOutput(os.Stderr)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&owner, "owner", "", "", "repository owner or org")
	if err := rootCmd.MarkFlagRequired("owner"); err != nil {
		panic(err)
	}
	rootCmd.Flags().StringSliceVarP(&repos, "repo", "", []string{}, "repository name")
	rootCmd.Flags().BoolVarP(&perDay, "per-day", "", false, "count stars per day")
	rootCmd.Flags().BoolVarP(&perMonth, "per-month", "", false, "count stars per month")
	rootCmd.Flags().BoolVarP(&perYear, "per-year", "", false, "count stars per year")
}
