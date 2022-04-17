package migrations

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/laijunbin/go-migrate/config"
	"github.com/laijunbin/go-migrate/pkg/interfaces"
	"github.com/laijunbin/go-migrate/pkg/lib/mysql"
)

func init() {
	config.Migrations = append(config.Migrations, CreateWordsTable())
}

type WordsTable struct{}

func CreateWordsTable() interfaces.Migration {
	return &WordsTable{}
}

func getInitWords() []map[string]interface{} {
	file, _ := os.Open("data/words.txt")
	defer file.Close()

	data, _ := ioutil.ReadAll(file)
	var words []map[string]interface{}

	for _, line := range strings.Split(string(data), "\n") {
		s := strings.Split(line, ":")
		words = append(words, map[string]interface{}{
			"english": s[0],
			"chinese": s[1],
		})
	}

	return words
}

func (t *WordsTable) Up() error {
	words := getInitWords()
	return mysql.Schema.Create("words", func(table interfaces.Blueprint) {
		table.Id("id", 10)
		table.String("english", 100)
		table.String("chinese", 100)
		table.Timestamps()
	}).Seed(words...)
}

func (t *WordsTable) Down() error {
	return mysql.Schema.DropIfExists("words")
}
