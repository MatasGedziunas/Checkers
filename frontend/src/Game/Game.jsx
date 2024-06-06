import React, { useState, useEffect, useRef } from "react";
import GameFunctionality from "./GameFunctionality";
import bMSvg from "../assets/bM.svg"; // Import the SVG files
import bKSvg from "../assets/bK.svg";
import wMSvg from "../assets/wM.svg";
import wKSvg from "../assets/wK.svg";
import "./Game.css";

function Game() {
  const functionality = new GameFunctionality();
  // const [board, setBoard] = useState(functionality.getStartingBoard());
  // const [board, setBoard] = useState(functionality.getCapturesBoard());
  const [board, setBoard] = useState(functionality.getDoubleCapturesBoard());
  const [possibleMoves, setPossibleMoves] = useState([]);
  const [turn, setTurn] = useState("w");
  const refs = useRef({});

  function updatePossibleMoves() {
    let newPossibleMoves = [];
    let maxCaptureCount = 0;

    for (let i = 0; i < board.length; i++) {
      newPossibleMoves.push([]);
      for (let j = 0; j < board[i].length; j++) {
        if (board[i][j] == "." || board[i][j][0] != turn) {
          newPossibleMoves[i].push([]);
        } else {
          const temp = functionality.getPossibleMoves(board, i, j);
          maxCaptureCount = Math.max(maxCaptureCount, temp.captureCount);
          newPossibleMoves[i].push({
            cords: temp.moves,
            captureCount: temp.captureCount,
          });
        }
      }
    }
    if (maxCaptureCount > 0) {
      for (let i = 0; i < newPossibleMoves.length; i++) {
        let row = newPossibleMoves[i];
        row.forEach((possibleMove) => {
          if (possibleMove.captureCount != maxCaptureCount) {
            possibleMove.cords = [];
          }
        });
      }
    }
    setPossibleMoves(newPossibleMoves);
  }

  useEffect(() => {
    updatePossibleMoves();
  }, [board]);

  function handleCheckerClick(row, col) {
    const possibleMoveImages = document.querySelectorAll(".possible-move");
    possibleMoveImages.forEach((img) => img.remove());
    let checkerPossibleMoves = possibleMoves[row][col];
    if (!checkerPossibleMoves.cords) {
      return;
    }
    console.log(checkerPossibleMoves);
    for (let move of checkerPossibleMoves.cords) {
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
                    <p>
                      {rowIndex} {cellIndex}
                    </p>
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
