// TreeNode.jsx
import React from "react";
import PropTypes from "prop-types";
import { Tree, TreeNode as OrgNode} from "react-organizational-chart";
import "../style/TreeNode.css";

function CustomNode({ item1, item2 }) {
  return (
    <div className="tree-combine">
      <div className="tree-combine-content">
        <div className="item">
          <img src={item1.Image} alt={item1.Name} />
          <span>{item1.Name}</span>
        </div>
        <span className="plus-sign">+</span>
        <div className="item">
          <img src={item2.Image} alt={item2.Name} />
          <span>{item2.Name}</span>
        </div>
      </div>
    </div>
  );
}


function TreeNode({ node }) {
  if (!node) return null;

  return (
    <OrgNode label={<CustomNode item1={node.Item1} item2={node.Item2} />}>
      {/* Children1 - Left Side */}
      {node.Children1 &&
        node.Children1.map((child, idx) => (
          <TreeNode key={`c1-${idx}`} node={child} />
        ))}
      {/* Children2 - Right Side */}
      {node.Children2 &&
        node.Children2.map((child, idx) => (
          <TreeNode key={`c2-${idx}`} node={child} />
        ))}
    </OrgNode>
  );
}

TreeNode.propTypes = {
  node: PropTypes.shape({
    Item1: PropTypes.object.isRequired,
    Children1: PropTypes.array,
    Item2: PropTypes.object.isRequired,
    Children2: PropTypes.array,
  }).isRequired,
};

export default TreeNode;
