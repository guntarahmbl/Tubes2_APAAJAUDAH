package main

type Stack []int

// Push menambahkan elemen ke atas stack
func (s *Stack) Push(v int) {
    *s = append(*s, v)
}

// Pop menghapus dan mengembalikan elemen paling atas
func (s *Stack) Pop() (int, bool) {
    if len(*s) == 0 {
        return 0, false // stack kosong
    }
    // ambil elemen terakhir
    idx := len(*s) - 1
    val := (*s)[idx]
    // potong slice
    *s = (*s)[:idx]
    return val, true
}
