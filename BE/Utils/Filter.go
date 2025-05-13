package utils


// fungsi untuk membersihkan NullChild pada tree
func FilterNullChildrens(node *TreeNode) {
	if node == nil {
		return
	}

	// Filter Children1
	var filteredChildren1 []*TreeNode
	for _, child := range node.Children1 {
		if child.Item1["Name"] != "Null" && child.Item2["Name"] != "Null" {
			filteredChildren1 = append(filteredChildren1, child)
			FilterNullChildrens(child) 
		}
	}
	node.Children1 = filteredChildren1

	// Filter Children2
	var filteredChildren2 []*TreeNode
	for _, child := range node.Children2 {
		if child.Item1["Name"] != "Null" && child.Item2["Name"] != "Null" {
			filteredChildren2 = append(filteredChildren2, child)
			FilterNullChildrens(child) 
		}
	}
	node.Children2 = filteredChildren2
}


// fungsi untuk membersihkan parent pada tree
func FilterAllParents(nodes []*TreeNode) []*TreeNode {
	var allChildren1 []*TreeNode
	for _, node := range nodes {
		if len(node.Children1) > 0 {
			allChildren1 = append(allChildren1, node.Children1[0])
		}
	}
	return allChildren1
}