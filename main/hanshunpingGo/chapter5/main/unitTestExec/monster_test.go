package unitTestExec

import "testing"

func TestMonster_Store(t *testing.T) {
	monster := Monster{Name: "牛魔王", Age: 500, Skill: "牛魔拳"}
	err := monster.Store()
	if err != nil {
		t.Fatal("monster store err=", err)
	}
	t.Log("monster store end...", monster)
}

func TestMonster_ReStore(t *testing.T) {
	monster := Monster{}
	err := monster.ReStore()
	if err != nil {
		t.Fatal("monster restore err=", err)
	}
	t.Log("monster restore end...", monster)
}
