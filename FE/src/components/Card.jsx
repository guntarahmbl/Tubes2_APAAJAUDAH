import React from "react";
import PropTypes from 'prop-types';
import "../style/Card.css";

function Card(props) {
  return (
    <div className="card">
      <img className="card-img" src={props.gambar} alt={props.gambar} />
      <div className="card-title">{props.nama}</div>
      <button className="button" onClick={props.onView}>View</button>
    </div>
  );
}

Card.propTypes = {
  index: PropTypes.number,
  nama: PropTypes.string,
  gambar: PropTypes.string,
  onView: PropTypes.func,
};

export default Card;
