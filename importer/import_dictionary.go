package importer

import (
	"encoding/json"
	"fmt"
	"github.com/eakarpov/msaot/db/dblib"
	"github.com/eakarpov/msaot/db/sqlite/models"
	"github.com/eakarpov/msaot/lexicon/lemmas"
	"github.com/eakarpov/msaot/lexicon/pos"
	"github.com/eakarpov/msaot/lexicon/synthesizer"
	"github.com/volatiletech/null/v8"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func GetData() *sheets.ValueRange {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "1Yl8Kk3M1yKCVa7cAsRxTQaJG5Vdnp6yGy6ceFGQb4z8"
	readRange := "words!A2:X"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	return resp
}

var gPositions []*models.GrammarPosition

func LoadGPositions(database dblib.DB) error {
	grammarPositions, err := database.GPositions().GetAll(context.Background())
	if err != nil {
		return err
	}
	gPositions = grammarPositions
	return nil
}

func GetGPositionByFlexyConfig(fConfig *lemmas.Lemma) *models.GrammarPosition {
	caseVal := null.Int64{
		Int64: 0,
		Valid: false,
	}
	numVal := null.Int64{
		Int64: 0,
		Valid: false,
	}
	genVal := null.Int64{
		Int64: 0,
		Valid: false,
	}
	personVal := null.Int64{
		Int64: 0,
		Valid: false,
	}
	tenseVal := null.Int64{
		Int64: 0,
		Valid: false,
	}
	if fConfig.NCase != nil {
		numVal = null.Int64{
			Int64: int64(fConfig.NCase.Gender),
			Valid: true,
		}
		caseVal = null.Int64{
			Int64: int64(fConfig.NCase.Case),
			Valid: true,
		}
		if fConfig.NCase.Person != nil {
			personVal = null.Int64{
				Int64: int64(*fConfig.NCase.Person),
				Valid: true,
			}
		}
	}
	if fConfig.VCase != nil {
		numVal = null.Int64{
			Int64: int64(fConfig.VCase.Number),
			Valid: true,
		}
		tenseVal = null.Int64{
			Int64: int64(fConfig.VCase.Tense),
			Valid: true,
		}
		personVal = null.Int64{
			Int64: int64(fConfig.VCase.Person),
			Valid: true,
		}
	}
	for _, gPos := range gPositions {
		if gPos.GNumber == numVal &&
			gPos.GCase == caseVal &&
			gPos.GGender == genVal &&
			gPos.GPerson == personVal &&
			gPos.GTense == tenseVal {
			return gPos
		}
	}
	return nil
}

func ImportRecord(database dblib.DB, record []string) {
	if len(record) > 3 {
		lemma := lemmas.GetLemmaConfig(record[3])
		lemma.Normal = record[1]
		l, _ := database.Lemmas().AddOne(context.Background(), lemma)
		var wForms []*lemmas.Lemma
		switch lemma.Pos {
		case pos.NOUN:
			animate := false
			if lemma.Animate != nil {
				animate = *lemma.Animate
			}
			wForms = synthesizer.GetWordForms(lemma.Normal, lemma.Pos, lemma.NCase.Gender, animate)
		}
		if wForms != nil {
			for _, wForm := range wForms {
				gPos := GetGPositionByFlexyConfig(wForm)
				if gPos != nil {
					_, _ = database.Flexies().AddOne(context.Background(), wForm, l.ID, gPos.ID)
				}
			}
		}
		fmt.Println(lemma.Normal)
	}
}
