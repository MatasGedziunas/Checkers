import { useState } from "react";
import "./Home.css";
import { useNavigate } from "react-router-dom";

function Home() {
  const navigate = useNavigate();
  const [name, setName] = useState("");

  const handleSubmit = () => {
    console.log(name);
    navigate("/waitingRoom");
  };

  return (
    <div className="home-container">
      <label htmlFor="playerName">Type your name:</label>
      <input
        type="text"
        id="playerName"
        onChange={(e) => setName(e.target.value.trim())}
      ></input>
      <button onClick={handleSubmit}>Create game</button>
    </div>
  );
}

export default Home;
