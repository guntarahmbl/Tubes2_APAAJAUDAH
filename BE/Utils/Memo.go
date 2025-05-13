package utils



// variable-variable global untuk memoization

var recipesMemo = make(map[string][]*TreeNode)  // menyimpan resep-resep dari elemen
var combinationsMemo = make(map[string][]*TreeNode)  // menyimpan kombinasi resep-resep dari dua elemen
