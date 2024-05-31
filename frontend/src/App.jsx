import { Route, Routes } from "react-router-dom";
import "./App.css";
import Home from "./Home/Home";
import Game from "./Game/Game";

function App() {
  return (
    <div className="container">
      <Routes>
        <Route exact path="/" Component={Home}></Route>
        <Route path="/game" Component={Game}></Route>
      </Routes>
    </div>
  );
}

export default App;
