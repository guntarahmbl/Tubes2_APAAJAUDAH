import React, { useState, useEffect } from "react";
import '../style/MainPage.css';
import Card from "./Card.jsx"
import Pagination from "./Pagination.jsx"
import PopUpTree from "./PopUpTree.jsx";

const itemsPerPage = 27;

function MainPage({data, selectedElement ,setSelectedElement, treeList, setTreeList, algorithm, jumlahResep}) {
  const [currentPage, setCurrentPage] = useState(1);
  const [popUpVisible, setPopUpVisible] = useState(false);
  const [selectedItem, setSelectedItem] = useState(null);
  const [time, setTime] = useState(null);
  const [nodeCount, setNodeCount] = useState(null);

  const totalPages = Math.ceil(data.length / itemsPerPage);
  const start = (currentPage - 1) * itemsPerPage;
  const currentItems = data.slice(start, start + itemsPerPage);

  useEffect(() => {
    console.log("useEffect triggered:", selectedElement, algorithm, jumlahResep);

    if (!selectedElement) return;
    const apiUrl = import.meta.env.VITE_API_URL;

    fetch(`${apiUrl}/recipes?target=${selectedElement}&algorithm=${algorithm}&maxRecipe=${jumlahResep}`)
    .then((res) => res.json())
      .then((data) => {
        setTreeList(data.recipes);
        setTime(data.time);
        setNodeCount(data.count);
      })
      .catch((err) => {
        console.error("Gagal mengambil data:", err);
      });
  }, [selectedElement]);
  
  const handleViewClick = (item) => {
    setSelectedItem(item);
    setSelectedElement(item.name);
    setPopUpVisible(true);
    // setTreeList(dummyTreeList);

  };

  const handleClosePopup = () => {
    setPopUpVisible(false);  // cuma ini yang perlu dilakukan dari MainPage
    setSelectedElement(null);
  };

  return (
    <div>
      <div className="card-container">
        {currentItems.map((item, index) => (
          <Card key={item.name} nama={item.name} index={index} gambar={item.image} onView={() => handleViewClick(item)}/>
        ))}
      </div>

      <PopUpTree visible={popUpVisible} treeData={treeList} item={selectedItem} onClose={handleClosePopup} time={time} nodeCount={nodeCount}/>

      <Pagination totalPages={totalPages} currentPage={currentPage} setCurrentPage={setCurrentPage}/>

      <div style={{ backgroundColor: "aliceblue", padding: "5px" }}>
        <a>&copy; {new Date().getFullYear()} Little Alchemy Solver</a>
      </div>
    </div>
  );
}

export default MainPage;