import { useNavigate } from "react-router-dom";
import "../WaitingRoom/WaitingRoom.css";

function WaitingRoom() {
  const navigate = useNavigate();

  function joinGame() {
    navigate("/game");
  }

  return (
    <div className="lobby-container">
      <div className="game-lobby">
        <div className="header">
          <h3>Your games</h3>
        </div>
        <div className="white-player">
          <div>
            <p>User1</p>
          </div>
        </div>
        <div className="black-player">
          <div>
            <p>User2</p>
          </div>
        </div>
        <div className="board-string">
          <div>
            <p>,...............................</p>
          </div>
        </div>
        <div className="join-game">
          <div>
            <button onClick={() => joinGame()}>Join game</button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default WaitingRoom;
