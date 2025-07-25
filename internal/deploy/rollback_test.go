package deploy

import "testing"

func TestRollback(t *testing.T) {
	t.Log("test rollback")
	err := RollbackInstall(APP, "20250725135833")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("test rollback success")
}

func TestRollbackDeep(t *testing.T) {
	err := RollbackInstall(APP2, "20250725144723")
	if err != nil {
		t.Error(err)
	}
}
