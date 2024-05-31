import { useState } from "react";
import "./Home.css";
import { useNavigate } from "react-router-dom";

function Home() {
  const navigate = useNavigate();
  const [name, setName] = useState("");

  const handleSubmit = () => {
    console.log(name);
    navigate("/game");
  };

  return (
    <div className="home-container">
      <label htmlFor="playerName">Type your name:</label>
      <input
        type="text"
        id="playerName"
        onChange={(e) => setName(e.target.value.trim())}
      ></input>
      <button onClick={handleSubmit}>Submit</button>
    </div>
  );
}

export default Home;
