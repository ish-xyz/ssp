package server

type ViewItem struct {
	Name    string
	Metdata map[string]string
}

type ListViewItems struct {
	Items []*ViewItem
}
