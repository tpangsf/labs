/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "swap client"
	app.Usage = "interact with node swap contract"
	app.HideVersion = true

	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "deploy the contract in seele",
			Flags: []cli.Flag{
				addressFlag,
				fromFlag,
				feeFlag,
				nonceFlag,
			},
			Action: Deploy,
		},
		{
			Name:  "create",
			Usage: "create a new contract for htlc transaction",
			Flags: []cli.Flag{
				addressFlag,
				fromFlag,
				toFlag,
				amountFlag,
				secretHashFlag,
				feeFlag,
				nonceFlag,
			},
			Action: Create,
		},
		{
			Name:  "withdraw",
			Usage: "get seele from the contract",
			Flags: []cli.Flag{
				addressFlag,
				fromFlag,
				secretFlag,
				feeFlag,
				nonceFlag,
			},
			Action: Withdraw,
		},
		{
			Name:  "refund",
			Usage: "refund seele from the contract",
			Flags: []cli.Flag{
				addressFlag,
				fromFlag,
				feeFlag,
				nonceFlag,
			},
			Action: Refund,
		},
		{
			Name:   "gensecret",
			Usage:  "generate secret and secret hash, secret hash is generated by sha256 ",
			Action: GenSecret,
		},
	}

	// sort commands and flags by name
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
