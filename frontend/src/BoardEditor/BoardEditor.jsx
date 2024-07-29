import React, { useEffect, useState } from "react";
import GameFunctionality from "../Game/GameFunctionality";
import bMSvg from "../assets/bM.svg"; // Import the SVG files
import bKSvg from "../assets/bK.svg";
import wMSvg from "../assets/wM.svg";
import wKSvg from "../assets/wK.svg";
import "../Game/Game.css";
import "./BoardEditor.css";
import { useNavigate } from "react-router-dom";
import { useLocation } from "react-router-dom";

function BoardEditor() {
  const navigate = useNavigate();
  const location = useLocation();
  const functionality = new GameFunctionality();
  const { encodedBoard } = location.state || {};
  let boardString = encodedBoard
    ? encodedBoard
    : functionality.getEmptyBoardString();
  const [board, setBoard] = useState(functionality.decodeBoard(boardString));
  const [turn, setTurn] = useState("w");
  const [lastClicked, setLastClicked] = useState(null);
  const [isValidBoard, setIsValidBoard] = useState(
    functionality.isValidBoard(board)
  );

  function handleClick(e, row, col, checkerToAdd = null) {
    removeFocus();
    if (lastClicked && checkerToAdd === null) {
      const [lastRow, lastCol, lastChecker] = lastClicked;
      // Move the checker in the board state
      const newBoard = [...board];
      newBoard[row][col] = lastChecker;
      if (lastRow && lastCol) {
        newBoard[lastRow][lastCol] = ".";
      }
      setBoard(newBoard);
      setLastClicked(null);
    } else {
      if (row && col) {
        setLastClicked([row, col, board[row][col]]);
      } else {
        setLastClicked([row, col, checkerToAdd]);
      }
      e.currentTarget.classList.add("focus");
    }
  }

  function removeFocus() {
    const elements = document.querySelectorAll(".focus");
    elements.forEach((element) => {
      element.classList.remove("focus");
    });
  }

  function getImgSrc(cell) {
    switch (cell) {
      case "b":
        return bMSvg;
      case "bq":
        return bKSvg;
      case "w":
        return wMSvg;
      case "wq":
        return wKSvg;
      default:
        return null;
    }
  }

  const handleInputChange = (e) => {
    const text = e.target.value;
    setBoard(functionality.decodeBoard(text));
  };

  useEffect(() => {
    setIsValidBoard(functionality.isValidBoard(board));
  }, [board]);

  return (
    <div className="game-container">
      <div className="game-screen">
        <div className="board">
          {board.map((row, rowIndex) => (
            <React.Fragment key={rowIndex}>
              {row.map((cell, cellIndex) => {
                const key = `${rowIndex}-${cellIndex}`;
                const color =
                  (rowIndex + cellIndex) % 2 === 0 ? "white" : "black";
                return (
                  <div
                    key={key}
                    className={`cell ${color} ${
                      lastClicked &&
                      lastClicked[0] === rowIndex &&
                      lastClicked[1] === cellIndex
                        ? "focus"
                        : ""
                    }`}
                    onClick={
                      color === "black"
                        ? (e) => handleClick(e, rowIndex, cellIndex)
                        : undefined
                    }
                    data-key={key}
                  >
                    {getImgSrc(cell) && <img src={getImgSrc(cell)} alt="" />}
                    <p>
                      {rowIndex} {cellIndex}
                    </p>
                  </div>
                );
              })}
            </React.Fragment>
          ))}
        </div>
        <div className="game-config">
          <div className="board-string">
            <label>Board string representation:</label>
            <input
              type="text"
              value={functionality.encodeBoard(board)}
              onChange={handleInputChange}
            />
            <p style={{ color: isValidBoard ? "rgb(158, 249, 158)" : "red" }}>
              {isValidBoard ? "Board valid" : "Board not valid"}
            </p>
          </div>
          <div className="pieces-container">
            <div onClick={(e) => handleClick(e, null, null, "b")}>
              <img src={bMSvg} alt="Black piece" />
            </div>
            <div onClick={(e) => handleClick(e, null, null, "bq")}>
              <img src={bKSvg} alt="Black king" />
            </div>
            <div onClick={(e) => handleClick(e, null, null, "w")}>
              <img src={wMSvg} alt="White piece" />
            </div>
            <div onClick={(e) => handleClick(e, null, null, "wq")}>
              <img src={wKSvg} alt="White king" />
            </div>
          </div>
          <div className="select-turn-container">
            <label>Select turn:</label>
            <div className="select-turn">
              <button
                className={`${turn == "w" ? "white" : ""}`}
                onClick={() => setTurn("w")}
              >
                White
              </button>
              <button
                className={`${turn == "b" ? "black" : ""}`}
                onClick={() => setTurn("b")}
              >
                Black
              </button>
            </div>
          </div>
          <div className="play">
            <button
              onClick={() =>
                navigate("/game", {
                  state: {
                    encodedBoard: functionality.encodeBoard(board),
                    curTurn: turn,
                  },
                })
              }
            >
              Play from this position
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default BoardEditor;
