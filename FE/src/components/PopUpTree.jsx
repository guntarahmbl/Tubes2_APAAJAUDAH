import React, { useState } from "react";
import TreeNode from "./TreeNode";
import { Tree } from "react-organizational-chart";
import "../style/PopUpTree.css";

function PopUpTree({ visible, onClose, item, treeData }) {
  const [currentPage, setCurrentPage] = useState(0);

  if (!visible || !item || !treeData || treeData.length === 0) return null;

  return (
    <div className="popup-overlay">
      <div className="popup-content">
        <button className="close-btn" onClick={() => {onClose(); setCurrentPage(0)}}>X</button>
        <h2>Recipe Tree for: {item.nama}</h2>

        {/* Parent item at the top */}
        <div className="parent-node">
          <div className="item">
            <img src={item.gambar} alt={item.nama} />
            <span>{item.nama}</span>
          </div>
        </div>

        {/* Show current tree */}
        <div className="">
          <Tree
            lineWidth={"2px"}
            lineColor={"#ccc"}
            lineBorderRadius={"10px"}
            label={<></>} // Kosongkan label root, karena sudah tampil di parent
          >
            <TreeNode node={treeData[currentPage]} />
          </Tree>
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
            onClick={() =>
              setCurrentPage((prev) => Math.min(prev + 1, treeData.length - 1))
            }
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
