import React, {useState} from "react";
import TreeNode from "./TreeNode.jsx";
import "../style/PopUpTree.css";

function PopUpTree({ visible, onClose, item, treeData }) {
  const [currentPage, setCurrentPage] = useState(0)
  if (!visible || !item || !treeData) return null;

  return (
    <div className="popup-overlay">
      <div className="popup-content">
        <button className="close-btn" onClick={() => {onClose(); setCurrentPage(0)}}>X</button>
        <h2>Recipe Tree for: {item.nama}</h2>

        {/* Parent item at the top */}
        <div className="parent-node">
          <div className="tree-combine">
            <div className="item">
              <img src={item.gambar} alt={item.nama} />
              <span>{item.nama}</span>
            </div>
          </div>
        </div>

        {/* Show current tree */}
        <div className="tree-node">
          <TreeNode node={treeData[currentPage]} selectedItem={item} />
        </div>

        {/* Pagination buttons */}
        <div className="tree-navigation">
          <button
            onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 0))}
            disabled={currentPage === 0}
          >
            Prev
          </button>
          <p>{currentPage + 1} / {treeData.length}</p>
          <button
            onClick={() => setCurrentPage((prev) => Math.min(prev + 1, treeData.length - 1))}
            disabled={currentPage === treeData.length - 1}
          >
            Next
          </button>
        </div>
      </div>
    </div>
  );
}

export default PopUpTree;
