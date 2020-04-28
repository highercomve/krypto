package main

import (
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "krypto"
	app.HelpName = "krypto"
	app.Usage = "AES cbc cryptography tool"
	app.Description = "AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256."
	app.Version = Version

	app.Commands = []cli.Command{
		decrypt(),
		encrypt(),
	}

	app.Run(os.Args)
}

func encrypt() cli.Command {
	return cli.Command{
		Name:        "encrypt",
		Aliases:     []string{"e"},
		ArgsUsage:   "",
		Usage:       "encrypt file",
		Description: "",
		Action: func(c *cli.Context) error {
			inFile := c.String("in")
			outFile := c.String("out")
			key := c.String("key")
			iv := c.String("iv")

			var output []byte
			input, err := ioutil.ReadFile(inFile)
			if err != nil {
				return cli.NewExitError(err, 1)
			}

			crypter, err := NewCrypter([]byte(key), []byte(iv))
			if err != nil {
				return cli.NewExitError(err, 1)
			}

			output, err = crypter.Encrypt(input)
			if err != nil {
				return cli.NewExitError(err, 1)
			}

			ioutil.WriteFile(outFile, output, 0644)

			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "in, i",
				Usage:    "Input file",
				Required: true,
			},
			cli.StringFlag{
				Name:     "out, o",
				Usage:    "Output file",
				Required: true,
			},
			cli.StringFlag{
				Name:     "key, k",
				Usage:    "Key string",
				Required: true,
			},
			cli.StringFlag{
				Name:     "iv",
				Usage:    "Initial vector string",
				Required: true,
			},
		},
	}
}

func decrypt() cli.Command {
	return cli.Command{
		Name:        "decrypt",
		Aliases:     []string{"d"},
		ArgsUsage:   "",
		Usage:       "decrypt file",
		Description: "",
		Action: func(c *cli.Context) error {
			inFile := c.String("in")
			outFile := c.String("out")
			key := c.String("key")
			iv := c.String("iv")

			var output []byte
			input, err := ioutil.ReadFile(inFile)
			if err != nil {
				return cli.NewExitError(err, 1)
			}

			crypter, err := NewCrypter([]byte(key), []byte(iv))
			if err != nil {
				return cli.NewExitError(err, 1)
			}

			output, err = crypter.Decrypt(input)
			if err != nil {
				return cli.NewExitError(err, 1)
			}

			ioutil.WriteFile(outFile, output, 0644)

			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "in, i",
				Usage:    "Input file",
				Required: true,
			},
			cli.StringFlag{
				Name:     "out, o",
				Usage:    "Output file",
				Required: true,
			},
			cli.StringFlag{
				Name:     "key, k",
				Usage:    "Key string",
				Required: true,
			},
			cli.StringFlag{
				Name:     "iv",
				Usage:    "Initial vector string",
				Required: true,
			},
		},
	}
}
