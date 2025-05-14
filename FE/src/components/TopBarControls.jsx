import React from 'react';
import "../style/TopBar.css"

function TopBarControls({ algorithm, setAlgorithm, jumlahType, setJumlahType, jumlahResep, setJumlahResep }) {
  return (
      <div className="controls">
        <div className="options">
          <label><input type="radio" name="Algorithm" value="bfs"
              checked={algorithm==="bfs"}
              onChange={() => setAlgorithm("bfs")} 
              /> 
              BFS</label>
          <label><input type="radio" name="Algorithm" value="dfs"
              checked={algorithm==="dfs"}
              onChange={() => setAlgorithm("dfs")}
              /> 
              DFS</label>
          <label><input type="radio" name="Jumlah" value="single" 
              checked={jumlahType === "single"} 
              onChange={() => {setJumlahType("single"); 
                              setJumlahResep(1);}
                        } 
            />
              Single Recipe </label>
          <label><input type="radio" name="Jumlah" value="multiple"
              checked={jumlahType === "multiple"}
              onChange={() => setJumlahType("multiple")}
            />
              Multiple Recipe</label>
        </div>

      <div className="input-group">
        {jumlahType === "multiple" && (
          <>
            <label htmlFor="start">Jumlah Resep: </label>
            <input id="start" type="text" className="input" 
                      value={jumlahResep}
                      onChange={(e) => setJumlahResep(Math.max(1, Math.min(Number(e.target.value), 100)))}/>
          </>
        )}
        {/* <button className="button"
                onClick={() => {
                  console.log("Algoritma:", algorithm);
                  console.log("Tipe Pencarian: ", jumlahType);
                  const recipeCount = parseInt(jumlahResep);
                  if (isNaN(recipeCount)) {
                      console.log("Invalid input for Jumlah Resep");
                    } else {
                      console.log("Jumlah Resep:", recipeCount);
                    } }}>
          Run
        </button> */}
      </div>
    </div>
  );
}

export default TopBarControls;
