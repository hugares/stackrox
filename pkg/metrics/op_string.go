// Code generated by "stringer -type=Op"; DO NOT EDIT.

package metrics

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Add-0]
	_ = x[AddMany-1]
	_ = x[Count-2]
	_ = x[Dedupe-3]
	_ = x[Exists-4]
	_ = x[Get-5]
	_ = x[GetAll-6]
	_ = x[GetMany-7]
	_ = x[GetFlowsForDeployment-8]
	_ = x[GetGrouped-9]
	_ = x[List-10]
	_ = x[Prune-11]
	_ = x[Reset-12]
	_ = x[Rename-13]
	_ = x[Remove-14]
	_ = x[RemoveMany-15]
	_ = x[RemoveFlowsByDeployment-16]
	_ = x[Search-17]
	_ = x[Update-18]
	_ = x[UpdateMany-19]
	_ = x[Upsert-20]
	_ = x[UpsertAll-21]
}

const _Op_name = "AddAddManyCountDedupeExistsGetGetAllGetManyGetFlowsForDeploymentGetGroupedListPruneResetRenameRemoveRemoveManyRemoveFlowsByDeploymentSearchUpdateUpdateManyUpsertUpsertAll"

var _Op_index = [...]uint8{0, 3, 10, 15, 21, 27, 30, 36, 43, 64, 74, 78, 83, 88, 94, 100, 110, 133, 139, 145, 155, 161, 170}

func (i Op) String() string {
	if i < 0 || i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
