import React, { useState } from "react";
import TreeNode from "./TreeNode";
import { Tree } from "react-organizational-chart";
import "../style/PopUpTree.css";

import zoomInIcon from "../assets/zoom-in.png";
import zoomOutIcon from "../assets/zoom-out.png";

function PopUpTree({ visible, onClose, item, treeData, time, nodeCount }) {
  const [currentPage, setCurrentPage] = useState(0);
  const [zoomLevel, setZoomLevel] = useState(1);

  if (!visible || !item || !treeData || treeData.length === 0) return null;

  return (
    <div className="popup-overlay">
      <div className="popup-content">

        <div className="popup-header">
          <button className="close-btn" onClick={() => {onClose(); setCurrentPage(0); setZoomLevel(1)}}>X</button>
          <div className="zoom-controls">
            <img
              src={zoomOutIcon}
              alt="Zoom Out"
              className="zoom-icon"
              onClick={() => setZoomLevel(z => Math.max(z - 0.1, 0.1))}
            />
            <span>{Math.round(zoomLevel * 100)}%</span>
            <img
              src={zoomInIcon}
              alt="Zoom In"
              className="zoom-icon"
              onClick={() => setZoomLevel(z => Math.min(z + 0.1, 2))}
            />
          </div>
          <h2>Recipe Tree for: {item.nama}</h2>
          <h2>Times: {(time*1000000).toFixed(2)} µs</h2>
          <h2>Node traversed: {nodeCount}</h2>

          {/* Zoom controls */}
        </div>

        {/* Parent item at the top */}
        <div className="parent-node" style={{ transform: `scale(${zoomLevel/0.7})`, transformOrigin: "top center" }}>
          <div className="item">
            <img src={item.gambar} alt={item.nama} />
            <span>{item.nama}</span>
          </div>
        </div>

        {/* Show current tree */}
        <div className="tree-container" style={{ transform: `scale(${zoomLevel})`, transformOrigin: "top center" }}>
          <Tree
            lineWidth={"2px"}
            lineColor={"#ccc"}
            lineBorderRadius={"10px"}
            label={<></>} // Kosongkan label root, karena sudah tampil di parent
          >
            <TreeNode node={treeData[currentPage]} />
          </Tree>
        </div>

        {/* Pagination */}
        <div className="tree-navigation sticky-footer">
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
