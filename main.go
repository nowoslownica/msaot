package main

import (
	"fmt"
	"github.com/eakarpov/msaot/graphematical"
	_case "github.com/eakarpov/msaot/lexicon/case"
	"github.com/eakarpov/msaot/lexicon/dictionary"
	"github.com/eakarpov/msaot/lexicon/gender"
	"github.com/eakarpov/msaot/lexicon/number"
	"github.com/eakarpov/msaot/lexicon/pos"
	"github.com/eakarpov/msaot/lexicon/synthesizer"
	"github.com/eakarpov/msaot/morphological"
	"github.com/eakarpov/msaot/syntax"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var genderVal int
	var posVal string

	app := &cli.App{
		Name:  "msaot",
		Usage: "Interslavic natural language processing tool",
		Commands: []*cli.Command{
			{
				Name:    "lemma",
				Aliases: []string{"l"},
				Usage:   "Enter a word to get list of lemmas",
				Action: func(c *cli.Context) error {
					word := c.Args().First()
					if word != "" {
						_, close, err := dictionary.InitMain()
						if err != nil {
							return err
						}
						defer close()
						lemmas, err := morphological.NormalizeAuto(word)
						if err != nil {
							return err
						}
						for _, lemma := range lemmas {
							fmt.Printf(`%s (%s)`, lemma.Normal, lemma.Pos)
							fmt.Println()
						}
						return nil
					}
					return fmt.Errorf("no input")
				},
			},
			{
				Name:    "syntaxtree",
				Aliases: []string{"st"},
				Usage:   "Build a syntax tree of a sentence",
				Action: func(c *cli.Context) error {
					sentence := c.Args().First()
					if sentence != "" {
						_, close, err := dictionary.InitMain()
						if err != nil {
							return err
						}
						defer close()
						tokenizer := graphematical.Tokenizer{}
						tokenizer.Parse(sentence)
						LoS, err := morphological.BuildLemmaChains(tokenizer.Items)
						if err != nil {
							return err
						}
						sentences := morphological.BuildSentenceLemmaChains(LoS)
						for _, sntnce := range sentences {
							tree, err := syntax.BuildSyntaxTree(sntnce.Lemmas)
							if err != nil {
								return err
							}
							fmt.Println(tree)
						}
						return nil
					}
					return fmt.Errorf("no input")
				},
			},
			{
				Name:    "lat2cyr",
				Aliases: []string{"l2c"},
				Usage:   "Enter an Interslavic latin word form and recieve a cyrillic one",
				Action: func(context *cli.Context) error {
					word := context.Args().First()
					result := graphematical.ImportedLat2SimpleCyr(word)
					fmt.Println(result)
					return nil
				},
			},
			{
				Name:    "declinator",
				Aliases: []string{"decl"},
				Usage:   "Enter a word (noun) and its gender and get the declination forms",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "gender",
						Aliases:     []string{"g"},
						Required:    true,
						Usage:       "1 - masculine, 2 - feminine, 3 - neutral",
						Destination: &genderVal,
					},
					&cli.StringFlag{
						Name:        "pos",
						Aliases:     []string{"p"},
						Usage:       "n - noun, adj - adjective",
						Value:       "n",
						Destination: &posVal,
					},
				},
				Action: func(context *cli.Context) error {
					word := context.Args().First()
					lemmas := synthesizer.GetWordForms(word, pos.POS(posVal), gender.Gender(genderVal), false)
					fmt.Println("Normal form:", word)
					for _, l := range lemmas {
						str := fmt.Sprintf(
							`[Case: %s, Number: %s] - %s`,
							_case.ToString(l.NCase.Case),
							number.ToString(l.NCase.Number),
							l.Value,
						)
						fmt.Println(str)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
