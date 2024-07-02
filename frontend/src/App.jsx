import { Route, Routes } from "react-router-dom";
import "./App.css";
import Home from "./Home/Home";
import Game from "./Game/Game";
import BoardEditor from "./BoardEditor/BoardEditor";

function App() {
  return (
    <div className="container">
      <Routes>
        <Route exact path="/" Component={Home}></Route>
        <Route path="/game" Component={Game}></Route>
        <Route path="/boardEditor" Component={BoardEditor}></Route>
      </Routes>
    </div>
  );
}

export default App;
