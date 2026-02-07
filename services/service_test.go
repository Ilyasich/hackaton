package services

import (
	"testing"

	mods "github.com/Ilyasich/hackaton/models"
	"github.com/Ilyasich/hackaton/repositories/chats"
	"github.com/Ilyasich/hackaton/repositories/memory"
)

type mockRest struct {
	balances map[mods.AccountID]float64
}

func (m mockRest) GetBalance(id mods.AccountID) (float64, bool) {
	b, ok := m.balances[id]
	return b, ok
}

func TestSetBansAndAddUser(t *testing.T) {
	var repo memory.Repository
	chatrep := chats.ChatConxRep{}
	// Rest with 0 balance for addr0 and positive for addr1
	rest := mockRest{balances: map[mods.AccountID]float64{
		"addr0": 0,
		"addr1": 10.5,
	}}

	svc := New(&repo, rest, &chatrep)

	// Add user with addr1 should succeed
	ok := svc.AddUser(12345, "addr1")
	if !ok {
		t.Fatalf("expected add user to succeed for addr1")
	}
	if svc.IsBanned(12345) {
		t.Fatalf("user with positive balance should not be banned")
	}

	// Add another user with addr0 (zero balance) should fail on AddUser because GetBalance returns ok but zero -> still allowed to add but banned
	ok = svc.AddUser(22222, "addr0")
	if !ok {
		t.Fatalf("expected add user to succeed for addr0 (repo allows), but it may be banned")
	}
	if !svc.IsBanned(22222) {
		t.Fatalf("user with zero balance should be marked banned")
	}

	// Now change balance and run SetBans to ensure re-evaluation
	rest.balances["addr0"] = 5.0
	svc.SetBans()
	if svc.IsBanned(22222) {
		t.Fatalf("after balance update, user should be unbanned")
	}
}
