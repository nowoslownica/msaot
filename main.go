package main

import (
	"fmt"
	"github.com/eakarpov/msaot/graphematical"
	"github.com/eakarpov/msaot/lexicon/dictionary"
	"github.com/eakarpov/msaot/morphological"
	"github.com/eakarpov/msaot/syntax"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "msaot",
		Usage: "Interslavic natural language processing tool",
		Commands: []*cli.Command{
			{
				Name:    "lemma",
				Aliases: []string{"l"},
				Usage:   "Enter a word to get list of lemmas",
				Action:  func(c *cli.Context) error {
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
				Action:  func(c *cli.Context) error {
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
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
