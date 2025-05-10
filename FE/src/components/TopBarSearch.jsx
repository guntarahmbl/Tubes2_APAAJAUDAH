import React, { useState } from 'react';
import '../style/TopBar.css';

function TopBarSearch( {setSearchQuery, onSearch} ) {
  const [inputText, setInputText] = useState("");

  const handleSearch = () => {
    setSearchQuery(inputText);
  };

  return (
    <div className="top-bar-search">
      <div className="title">Little Alchemy 2 Solver</div>
      <div className="search-controls">
        <input 
          type="text" 
          placeholder="Search" 
          className="input-search"
          value={inputText}
          onChange={(e) => setInputText(e.target.value)}
        />
        <button className="button" onClick={handleSearch}>Search</button>
      </div>
    </div>
  );
}

export default TopBarSearch;
