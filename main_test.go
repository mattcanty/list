package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	messages "github.com/cucumber/messages-go/v10"
)

var opt = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func thereAreList(expectedListCount int) error {
	if ListCount() != expectedListCount {
		return fmt.Errorf("expected %d lists, but there is %d", expectedListCount, ListCount())
	}
	return nil
}

func aListCalledIsAdded(listName string) error {
	CreateNewList(listName)
	return nil
}

func isAddedToTheList(item, listName string) error {
	AddListItem(listName, item)
	return nil
}

func theListContains(listName string, expectedItems *messages.PickleStepArgument_PickleTable) error {
	list := GetList(listName)

	head := expectedItems.Rows

	for _, row := range head {
		expectedValue := row.Cells[0].Value

		if !contains(list, expectedValue) {
			return fmt.Errorf("could not find '%s' in '%s' list", expectedValue, listName)
		}
	}
	return nil
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func theListHasItems(listName string, count int) error {
	actual := len(GetList(listName))
	if count != actual {
		return fmt.Errorf("expected %d items in %s, but there is %d", count, listName, actual)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^"([^"]*)" is added to the "([^"]*)" list$`, isAddedToTheList)
	s.Step(`^the "([^"]*)" list contains$`, theListContains)
	s.Step(`^there are (\d+) lists$`, thereAreList)
	s.Step(`^a list called "([^"]*)" is added$`, aListCalledIsAdded)
	s.Step(`^there should be (\d+) list$`, thereAreList)
	s.Step(`^the "([^"]*)" list has (\d+) items$`, theListHasItems)

	s.BeforeScenario(func(*messages.Pickle) {
		DeleteAllLists()
	})
}
