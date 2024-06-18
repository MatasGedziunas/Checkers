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
  const [lastClickedSquare, setLastClickedSquare] = useState([]);

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
            cords: temp.cords,
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
    console.log("updating posssible moves");
    updatePossibleMoves();
  }, [board]);

  function handleCheckerClick(e, row, col) {
    const cellClicked = e.target.closest(".cell"); // Ensure you get the .cell div even if the target is an inner element
    if (cellClicked.querySelector("img.possible-move")) {
      removeAllPossibleMoveCheckers();
      const coordinates = lastClickedSquare.getAttribute("data-key").split("-");
      makeMove(parseInt(coordinates[0]), parseInt(coordinates[1]), row, col);
    } else {
      removeAllPossibleMoveCheckers();
      addPossibleMoveCheckers(row, col);
    }

    setLastClickedSquare(cellClicked);
  }

  function removeAllPossibleMoveCheckers() {
    const possibleMoveImages = document.querySelectorAll(".possible-move");
    possibleMoveImages.forEach((img) => img.remove());
  }

  function makeMove(rowFrom, colFrom, rowTo, colTo) {
    // database stuff validations etc.
    const directionRow = Math.sign(rowTo - rowFrom);
    const directionCol = Math.sign(colTo - colFrom);
    const checkerColorToRemove = board[rowFrom][colFrom] == "w" ? "b" : "w";
    board[rowTo][colTo] = board[rowFrom][colFrom];
    board[rowFrom][colFrom] = ".";
    let curRow = rowFrom + directionRow;
    let curCol = colFrom + directionCol;
    while (
      functionality.isInBounds(board, curRow, curCol) &&
      curRow != rowTo &&
      curCol != colTo
    ) {
      if (board[curRow][curCol][0] == checkerColorToRemove) {
        board[curRow][curCol] = ".";
      }
      curRow += directionRow;
      curCol += directionCol;
    }
    if (directionRow == rowTo - rowFrom) {
      setTurn(turn == "w" ? "b" : "w");
    }
  }

  function addPossibleMoveCheckers(row, col) {
    const checker = board[row][col];
    let checkerPossibleMoves = possibleMoves[row][col];
    console.log(checkerPossibleMoves);
    if (!checkerPossibleMoves.cords) {
      return;
    }
    for (let move of checkerPossibleMoves.cords) {
      const key = `${move[0]}-${move[1]}`;
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
                    data-key={key}
                    onClick={(e) => handleCheckerClick(e, rowIndex, cellIndex)}
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
