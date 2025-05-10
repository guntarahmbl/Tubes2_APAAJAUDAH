import React, { useState } from "react";
import '../style/MainPage.css';
import Card from "./Card.jsx"
import Pagination from "./Pagination.jsx"
import PopUpTree from "./PopUpTree.jsx";

const itemsPerPage = 27;

function MainPage({data, setSelectedElement, treeList, setTreeList}) {
  const [currentPage, setCurrentPage] = useState(1);
  const [popUpVisible, setPopUpVisible] = useState(false);
  const [selectedItem, setSelectedItem] = useState(null);

  const totalPages = Math.ceil(data.length / itemsPerPage);
  const start = (currentPage - 1) * itemsPerPage;
  const currentItems = data.slice(start, start + itemsPerPage);

  const handleViewClick = (item) => {
    setSelectedItem(item);
    setSelectedElement(item.nama);
    setPopUpVisible(true);
    setTreeList(dummyTreeList);
  };

  const handleClosePopup = () => {
    setPopUpVisible(false);  // cuma ini yang perlu dilakukan dari MainPage
  };

  return (
    <div>
      <div className="card-container">
        {currentItems.map((item, index) => (
          <Card key={item.nama} nama={item.nama} index={index} gambar={item.gambar} onView={() => handleViewClick(item)}/>
        ))}
      </div>

      <PopUpTree visible={popUpVisible} treeData={treeList} item={selectedItem} onClose={handleClosePopup} />

      <Pagination totalPages={totalPages} currentPage={currentPage} setCurrentPage={setCurrentPage}/>

      <div style={{ backgroundColor: "aliceblue", padding: "5px" }}>
        <a>&copy; {new Date().getFullYear()} Little Alchemy Solver</a>
      </div>
    </div>
  );
}

export default MainPage;