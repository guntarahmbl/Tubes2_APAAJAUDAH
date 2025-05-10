// TreeNode.jsx
import React from "react";
import PropTypes from "prop-types";
import "../style/TreeNode.css";

function TreeNode({ node }) {
  if (!node) return null;

  return (
    <div className="tree-node">
      {/* Combine two items */}
      <div className="tree-combine">
        <div className="item">
          <img src={node.item1.gambar} alt={node.item1.nama} />
          <span>{node.item1.nama}</span>
        </div>
        <span className="plus-sign">+</span>
        <div className="item">
          <img src={node.item2.gambar}/>
          <span>{node.item2.nama}</span>
        </div>
      </div>

      {(node.children1?.length > 0 || node.children2?.length > 0) && (
        <div className="tree-children-wrapper">
            {/* Children1 - Left Side */}
            {node.children1?.length > 0 && (
            <div className="tree-children-left tree-children">
                {node.children1.map((child, i) => (
                <TreeNode key={`c1-${child.item1.nama}-${i}`} node={child} />
                ))}
            </div>
            )}

            {/* Children2 - Right Side */}
            {node.children2?.length > 0 && (
            <div className="tree-children-right tree-children">
                {node.children2.map((child, i) => (
                <TreeNode key={`c2-${child.item2.nama}-${i}`} node={child} />
                ))}
            </div>
            )}
        </div>
        )}
    </div>
  );
}

TreeNode.propTypes = {
  node: PropTypes.shape({
    item1: PropTypes.object.isRequired,
    children1: PropTypes.array,
    item2: PropTypes.object.isRequired,
    children2: PropTypes.array,
  }).isRequired,
};

export default TreeNode;
