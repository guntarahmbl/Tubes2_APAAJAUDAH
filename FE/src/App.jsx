import React, { useState, useEffect } from 'react';
import TopBarControls from "./components/TopBarControls";
import TopBarSearch from "./components/TopBarSearch";
import MainPage from "./components/MainPage";
import "./App.css";

// dummyTreeListData

const dummyTreeList = [
  // Pohon 1
  {
    item1: { nama: "Water", gambar: "/images/Water.png" },
    item2: { nama: "Fire", gambar: "/images/Fire.png" },
    children1: [
      {
        item1: { nama: "Ice", gambar: "/images/Ice.png" },
        item2: { nama: "Heat", gambar: "/images/Heat.png" },
        children1: null,
        children2: null,
      },
    ],
    children2: null,
  },

  // Pohon 2
  {
    item1: { nama: "Cloud", gambar: "/images/Cloud.png" },
    item2: { nama: "Sun", gambar: "/images/Sun.png" },
    children1: null,
    children2: [
      {
        item1: { nama: "Light", gambar: "images/Light.png" },
        item2: { nama: "Heat", gambar: "images/Heat.png" },
        children1: null,
        children2: null,
      },
    ],
  },

  // Pohon 3
  {
  item1: { nama: "Air", gambar: "/images/Air.png" },
  children1: [
    {
      item1: { nama: "Earth", gambar: "/images/Earth.png" },
      children1: [
        {
          item1: { nama: "Earth", gambar: "/images/Earth.png" },
          children1: null,
          item2: { nama: "Pressure", gambar: "/images/Pressure.png" },
          children2: null,
        },
      ],
      item2: { nama: "Pressure", gambar: "/images/Pressure.png" },
      children2: [
        {
          item1: { nama: "Earth", gambar: "/images/Earth.png" },
          children1: null,
          item2: { nama: "Pressure", gambar: "/images/Pressure.png" },
          children2: null,
        },
      ],
    },
  ],
  item2: { nama: "Fire", gambar: "/images/Fire.png" },
  children2: [
    {
      item1: { nama: "Energy", gambar: "/images/Energy.png" },
      children1: [
        {
          item1: { nama: "Earth", gambar: "/images/Earth.png" },
          children1: null,
          item2: { nama: "Pressure", gambar: "/images/Pressure.png" },
          children2: null,
        },
      ],
      item2: { nama: "Sun", gambar: "/images/Sun.png" },
      children2: [
        {
          item1: { nama: "Earth", gambar: "/images/Earth.png" },
          children1: null,
          item2: { nama: "Pressure", gambar: "/images/Pressure.png" },
          children2: null,
        },
      ],
    },
  ],
  },
];


function App() {
  const [algorithm, setAlgorithm] = useState("bfs");
  const [jumlahType, setJumlahType] = useState("single");
  const [jumlahResep, setJumlahResep] = useState(1);
  const [selectedElement, setSelectedElement] = useState("");
  const [treeList, setTreeList] = useState([]);
  const [data, setData] = useState([]);
  const [searchQuery, setSearchQuery] = useState("");

  // SetDummyTreeList
  // useEffect(() => {
  //   setTreeList(dummyTreeList);
  // }, []);

  useEffect(() => {
    fetch("/allElements.json")
      .then((res) => res.json())
      .then((json) => setData(json))
      .catch((err) => console.error("Failed to load data:", err));
  }, []);

  


  const filteredData = data.filter(item =>
    item.nama.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <div className="app-container">
      <div className="top-bar">
        <TopBarSearch setSearchQuery={setSearchQuery} />
        <TopBarControls 
          algorithm={algorithm} 
          setAlgorithm={setAlgorithm} 
          jumlahType={jumlahType} 
          setJumlahType={setJumlahType} 
          jumlahResep={jumlahResep} 
          setJumlahResep={setJumlahResep} 
        />
        <MainPage data={filteredData} selectedElement={selectedElement} setSelectedElement={setSelectedElement} treeList={treeList} setTreeList={setTreeList} algorithm={algorithm} jumlahResep={jumlahResep}/>
      </div>
    </div>
  );
}

export default App;
