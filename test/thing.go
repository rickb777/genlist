package main

// +test slice:"Exists,Forall,Count,Distinct,DistinctBy,Foreach,First,MaxBy,MinBy,Single,Filter,SortBy,SortByDesc,IsSortedBy,IsSortedByDesc,Aggregate[Other],Average[Other],GroupBy[Other],Max[Other],Min[Other],MapTo[Other],Shuffle,Sum[Other]"
type Thing struct {
	Name   string
	Number Other
}
