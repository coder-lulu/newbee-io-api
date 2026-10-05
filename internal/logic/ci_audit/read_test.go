package ciaudit

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/coder-lulu/newbee-common/v2/middleware/keys"
	"github.com/coder-lulu/newbee-common/v2/orm/ent/hooks"
	"github.com/coder-lulu/newbee-io-api/internal/types"
	"github.com/coder-lulu/newbee-io-rpc/ent"
	_ "modernc.org/sqlite"
)

func TestRejectMissingTenantBeforeDatabaseAccess(t *testing.T) {
	for _, tenant := range []string{"", "0", "-1", "invalid"} {
		ctx := keys.NewContextManager().SetTenantID(context.Background(), tenant)
		if _, err := ListChangeHistory(ctx, nil, &types.ListChangeHistoryReq{}); err == nil {
			t.Fatalf("accepted tenant %q", tenant)
		}
		if _, err := ListLifecycleState(ctx, nil, &types.ListLifecycleStateReq{}); err == nil {
			t.Fatalf("accepted tenant %q", tenant)
		}
	}
}

func TestListsKeepTenantIsolationAndFalseFilters(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		t.Fatal(err)
	}
	client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.SQLite, db)))
	t.Cleanup(func() { _ = client.Close() })
	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := hooks.QuickSetup(client); err != nil {
		t.Fatal(err)
	}
	for tenant := uint64(1); tenant <= 2; tenant++ {
		ctx := hooks.SetTenantIDToContext(context.Background(), tenant)
		for n := 0; n < 2; n++ {
			_, err := client.CiChangeHistory.Create().SetOperationID(time.Now().String()).SetOperationType("update").SetNeedsApproval(n == 1).Save(ctx)
			if err != nil {
				t.Fatal(err)
			}
			_, err = client.CiLifecycleState.Create().SetStateID(time.Now().String()).SetStateName("completed").SetStateType("completed").SetEnteredAt(time.Now()).SetTriggerType("manual").SetHasError(n == 1).Save(ctx)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
	ctx := hooks.SetTenantIDToContext(context.Background(), 1)
	f := false
	history, err := ListChangeHistory(ctx, client, &types.ListChangeHistoryReq{Page: 1, PageSize: 1, NeedsApproval: &f})
	if err != nil {
		t.Fatal(err)
	}
	if history.Data.Total != 1 || len(history.Data.Data) != 1 {
		t.Fatalf("history tenant/filter leak: %+v", history)
	}
	state, err := ListLifecycleState(ctx, client, &types.ListLifecycleStateReq{Page: 1, PageSize: 1, HasError: &f})
	if err != nil {
		t.Fatal(err)
	}
	if state.Data.Total != 1 || len(state.Data.Data) != 1 {
		t.Fatalf("state tenant/filter leak: %+v", state)
	}
	if state.Data.Data[0]["exitedAt"] != "" || state.Data.Data[0]["hasError"] != false {
		t.Fatalf("optional time/boolean mapping: %+v", state.Data.Data[0])
	}
	pageTwo, err := ListChangeHistory(ctx, client, &types.ListChangeHistoryReq{Page: 2, PageSize: 1, NeedsApproval: &f})
	if err != nil || pageTwo.Data.Total != 1 || len(pageTwo.Data.Data) != 0 {
		t.Fatalf("pagination: %+v, %v", pageTwo, err)
	}
}
