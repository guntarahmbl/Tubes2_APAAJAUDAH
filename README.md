<h1 align="center"> Tugas Besar 2 IF2211 Strategi Algoritma </h1>
<h1 align="center">Pemanfaatan Algoritma BFS dan DFS dalam Pencarian Recipe  pada Permainan Little Alchemy 2</h1>

![APAAJAUDAH](apaajaudah.png)

## Deskripsi Algoritma BFS dan DFS pada Program
### Langkah Pembangunan Tree dengan BFS
1. Pertama-tama, dibuat queue  string dengan 1 anggota yaitu elemen yang dicari. Queue digunakan untuk menampung antrian elemen yang akan disimpan
2. Dilakukan iterasi sampai queue menjadi kosong
3. Setiap iterasi dilakukan dequeue dari pasangan elemen paling depan
4. Untuk setiap pasangan (kanan dan kiri) akan dicari pasangan-pasangan pembentuk elemen tersebut dari hasil scraping, seterusnya akan disebut sebagai children element
5. Jika children element adalah base element atau time, lanjutkan ke iterasi berikutnya
6. Jika tier children element lebih tinggi atau sama dengan element saat ini, lanjutkan
7. Tree baru dibentuk, masukkan ke children tree iterasi saat ini. kemudian Tree baru tersebut masukkan ke dalam queue
8. Iterasi selesai jika queue habis

### Langkah Pembangunan Tree dengan DFS
1. Pembangunan tree dengan DFS dilakukan dengan cara rekursif
2. Pertama-tama, fungsi menerima parameter tree yang mengandung elemen yang dicari (karena tree membutuhkan 2 elemen, dipasangkan dengan nullTree)
3. Untuk setiap item1 dan item2, akan dicari elemen-elemen pembentuknya (children element)
4. Jika children element adalah base element atau time, lanjutkan ke iterasi berikutnya
5. Jika tier children element lebih tinggi atau sama dengan element saat ini, lanjutkan
6. Tree baru dibentuk, sambungkan sebagai children
7. Kemudian fungsi rekursi dipanggil dengan parameter tree yang baru dibentuk


## Project Structure
```bash
|-- .gitignore
|-- .vscode
|-- BE
|   |-- Scrape
|   |   |-- allImages.go
|   |   |-- elementRecipes.go
|   |   |-- scrapeElements.go
|   |   |-- scrapeImagesLink.go
|   |   +-- scrapeTiers.go
|   +-- Utils
|       |-- BFS.go
|       |-- Converter.go
|       |-- DFS.go
|       |-- Filter.go
|       |-- Generator.go
|       |-- Memo.go
|       |-- Tree.go
|       +-- io.go
|-- FE
|   |-- .gitignore
|   |-- README.md
|   |-- dist
|   |   |-- allElements.json
|   |   |-- icon.jpg
|   |   |-- images.json
|   |   |-- index.html
|   |   +-- vite.svg
|   |-- eslint.config.js
|   |-- index.html
|   |-- package-lock.json
|   |-- package.json
|   |-- public
|   |   |-- allElements.json
|   |   |-- icon.jpg
|   |   |-- images.json
|   |   +-- vite.svg
|   |-- src
|   |   |-- App.css
|   |   |-- App.jsx
|   |   |-- components
|   |   |   |-- Card.jsx
|   |   |   |-- MainPage.jsx
|   |   |   |-- Pagination.jsx
|   |   |   |-- PopUpTree.jsx
|   |   |   |-- TopBarControls.jsx
|   |   |   |-- TopBarSearch.jsx
|   |   |   +-- TreeNode.jsx
|   |   |-- index.css
|   |   |-- main.jsx
|   |   +-- style
|   |       |-- Card.css
|   |       |-- MainPage.css
|   |       |-- Pagination.css
|   |       |-- PopUpTree.css
|   |       |-- TopBar.css
|   |       +-- TreeNode.css
|   +-- vite.config.js
|-- README.md
|-- data
|   |-- allElementsImage.json
|   |-- allElementsRecipes.json
|   |-- allElementsTiers.json
|   |-- elements_list.txt
|   |-- images.json
|   +-- recipes.json
|-- go.mod
|-- go.sum
|-- main.go
```


## Command Build Program (Untuk lokal)
1. Clone repository ini 
2. Lakukan build:
    ```bash
    npm run build
    ```
3. Jalankan di localhost:
    ```bash
    npm run dev
    ```
4. Web dapat digunakan

## Pemakaian Aplikasi (dari Internet)
Akses web melalui web browser dengan link di bawah:
```bash
https://tubes2-apaajaudah.vercel.app/
```


## Authors
### **Kelompok "APA AJA UDAH"**
|   NIM    |                  Nama                  |
| :------: | :-------------------------------------:|
| 13523056 |              Salman Hanif              |
| 13523114 |             Guntara Hambali            |
| 13523116 |           Fityatul Haq Rosyidi         |

## Link Youtube Presentasi
https://youtu.be/KCirBsWO4Gg