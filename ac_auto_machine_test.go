package ac_auto_machine

import "testing"

func TestAcAutoMachine_Search(t *testing.T) {
	ac := NewAcAutoMachine()
	ac.AddWord("hers")
	ac.AddWord("his")
	ac.AddWord("she")
	ac.AddWord("he")
	ac.AddWord("h")
	ac.AddWord("bbbb")
	ac.Build()
	t.Logf("%+v", ac.Search("ahishershe"))
}
