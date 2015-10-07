package main

// +test slice:"Size,Contains,Exists,Forall,Count,CountBy,Distinct,DistinctBy,Foreach,First,MaxBy,MinBy,Filter,Partition,SortBy,SortByDesc,IsSortedBy,IsSortedByDesc,Aggregate[Other],Mean[Other],GroupBy[Other],Max[Other],Min[Other],MapTo[Other],Shuffle,Sum[Other]"
type Thing struct {
	Name   string
	Number Other
}
