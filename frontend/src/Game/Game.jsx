import React, { useState, useEffect, useRef } from "react";
import GameFunctionality from "./GameFunctionality";
import bMSvg from "../assets/bM.svg"; // Import the SVG files
import bKSvg from "../assets/bK.svg";
import wMSvg from "../assets/wM.svg";
import wKSvg from "../assets/wK.svg";
import "./Game.css";

function Game() {
  const functionality = new GameFunctionality();
  const [board, setBoard] = useState(functionality.getStartingBoard());
  const [possibleMoves, setPossibleMoves] = useState([]);
  const [mustCapture, setMustCapture] = useState(false);

  const refs = useRef({});

  function updatePossibleMoves() {
    const newPossibleMoves = [];
    let newMustCapture = false;

    for (let i = 0; i < board.length; i++) {
      newPossibleMoves.push([]);
      for (let j = 0; j < board[i].length; j++) {
        if (board[i][j] == ".") {
          newPossibleMoves[i].push([]);
        } else {
          const temp = functionality.showPossibleMoves(board, i, j);
          if (temp.isCaptures) {
            newMustCapture = true;
          }
          if (newMustCapture && temp.isCaptures) {
            newPossibleMoves[i].push(temp.moves);
          } else if (!newMustCapture) {
            newPossibleMoves[i].push(temp.moves);
          } else {
            newPossibleMoves[i].push([]);
          }
        }
      }
    }

    setPossibleMoves(newPossibleMoves);
    setMustCapture(newMustCapture);
  }

  useEffect(() => {
    updatePossibleMoves();
  }, [board]);

  function handleCheckerClick(row, col) {
    const possibleMoveImages = document.querySelectorAll(".possible-move");
    possibleMoveImages.forEach((img) => img.remove());
    let checkerPossibleMoves = possibleMoves[row][col];
    console.log(checkerPossibleMoves);
    for (let move of checkerPossibleMoves) {
      addPossibleMoveChecker(move[0], move[1], board[row][col]);
    }
  }

  function addPossibleMoveChecker(row, col, checker) {
    const key = `${row}-${col}`;
    if (refs.current[key]) {
      const img = document.createElement("img");
      img.src = getCheckerImage(checker);
      img.alt = "Possible move";
      img.className = "possible-move";
      refs.current[key].appendChild(img);
    }
  }

  function getCheckerImage(cell) {
    switch (cell) {
      case "b":
        return bMSvg;
      case "bk":
        return bKSvg;
      case "w":
        return wMSvg;
      case "wk":
        return wKSvg;
      default:
        return null;
    }
  }

  return (
    <div className="game-container">
      <div className="game-screen">
        <div className="board">
          {board.map((row, rowIndex) => (
            <React.Fragment key={rowIndex}>
              {row.map((cell, cellIndex) => {
                const key = `${rowIndex}-${cellIndex}`;
                return (
                  <div
                    key={key}
                    ref={(el) => (refs.current[key] = el)}
                    className={`cell ${
                      (rowIndex + cellIndex) % 2 === 0 ? "white" : "black"
                    }`}
                    onClick={() => handleCheckerClick(rowIndex, cellIndex)}
                  >
                    {cell === "b" ? (
                      <img src={bMSvg} alt="Black piece" />
                    ) : cell === "bk" ? (
                      <img src={bKSvg} alt="Black king" />
                    ) : cell === "w" ? (
                      <img src={wMSvg} alt="White piece" />
                    ) : cell === "wk" ? (
                      <img src={wKSvg} alt="White king" />
                    ) : (
                      ""
                    )}
                  </div>
                );
              })}
            </React.Fragment>
          ))}
        </div>
      </div>
    </div>
  );
}

export default Game;
