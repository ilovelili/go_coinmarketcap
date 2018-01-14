// Copyright © 2018 Min Ju <route666@live.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ilovelili/coinmarketcap"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Tickers.",
	Long:  `Get Tickers info from coinmarkercap API`,
	RunE:  get,
}

const (
	accessKey       = ""
	secretAccesskey = ""
	usage           = "Usage: ./coinmarketcap get ticker"
)

func gettickers() error {
	client, err := coinmarketcap.NewClient(accessKey, secretAccesskey)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	result, err := client.GetTickers(ctx)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}

func get(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New(usage)
	}

	switch args[0] {

	case "ticker":
		return gettickers()
	default:
		return errors.New(usage)
	}
}

func init() {
	rootCmd.AddCommand(getCmd)
}
