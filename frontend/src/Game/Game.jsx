import React, { useState, useEffect, useRef } from "react";
import GameFunctionality from "./GameFunctionality";
import bMSvg from "../assets/bM.svg"; // Import the SVG files
import bKSvg from "../assets/bK.svg";
import wMSvg from "../assets/wM.svg";
import wKSvg from "../assets/wK.svg";
import "./Game.css";
import { useNavigate } from "react-router-dom";
import { useLocation } from "react-router-dom";

function Game() {
  const navigate = useNavigate();
  const location = useLocation();
  const functionality = new GameFunctionality();
  // const [board, setBoard] = useState(functionality.getStartingBoard());
  const { encodedBoard, curTurn } = location.state || {};
  let boardString = encodedBoard
    ? encodedBoard
    : functionality.getCapturesBoardString();
  const [board, setBoard] = useState(functionality.decodeBoard(boardString));
  // const [board, setBoard] = useState(functionality.getDoubleCapturesBoard());
  const [possibleMoves, setPossibleMoves] = useState([]);
  const [turn, setTurn] = useState(curTurn ? curTurn : "w");
  const [lastClickedSquare, setLastClickedSquare] = useState([]);
  const [checkerLastCaptured, setCheckerLastCaptured] = useState(null);
  const [maxCaptureCountForTurn, setMaxCaptureCountForTurn] = useState();

  const refs = useRef({});

  function updatePossibleMoves() {
    const encodedBoard = functionality.encodeBoard(board);
    const url =
      "http://localhost:3000/possibleMoves?boardString=" + encodedBoard;
    let possibleMoves = fetch(url, { method: "GET" })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok " + response.statusText);
        }
        return response.json();
      })
      .then((data) => {
        console.log(data);
        return data;
      })
      .catch((error) => console.log("Failed fetching possible moves: ", error));

    if (checkerLastCaptured != null) {
      const row = checkerLastCaptured[0];
      const col = checkerLastCaptured[1];
      const newPossibleMoves = possibleMoves.map((row) =>
        new Array(row.length).fill([])
      );
      newPossibleMoves[row][col] = functionality.getPossibleMoves(
        board,
        row,
        col
      );
      setPossibleMoves(newPossibleMoves);
    } else {
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
      console.log(`Max capture count: ${maxCaptureCount}`);
      setMaxCaptureCountForTurn(maxCaptureCount);
      setPossibleMoves(newPossibleMoves);
    }
  }

  useEffect(() => {
    console.log("updating posssible moves");
    updatePossibleMoves();
  }, [board]);

  useEffect(() => {
    if (checkerLastCaptured != null) {
      const row = checkerLastCaptured[0];
      const col = checkerLastCaptured[1];
      removeAllPossibleMoveCheckers();
      addPossibleMoveCheckers(row, col);
    }
  }, [possibleMoves]);

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
    // Create a new copy of the board to ensure state change detection
    let newBoard = board.map((row) => row.slice());

    const directionRow = Math.sign(rowTo - rowFrom);
    const directionCol = Math.sign(colTo - colFrom);
    const checkerColorToRemove = newBoard[rowFrom][colFrom] === "w" ? "b" : "w";

    newBoard[rowTo][colTo] = newBoard[rowFrom][colFrom];
    newBoard[rowFrom][colFrom] = ".";

    let curRow = rowFrom + directionRow;
    let curCol = colFrom + directionCol;

    while (
      functionality.isInBounds(newBoard, curRow, curCol) &&
      curRow !== rowTo &&
      curCol !== colTo
    ) {
      if (newBoard[curRow][curCol][0] === checkerColorToRemove) {
        newBoard[curRow][curCol] = ".";
      }
      curRow += directionRow;
      curCol += directionCol;
    }

    if (maxCaptureCountForTurn <= 1) {
      setTurn(turn === "w" ? "b" : "w");
      setCheckerLastCaptured(null);
      newBoard = functionality.updateQueens(newBoard);
    } else {
      setCheckerLastCaptured([rowTo, colTo]);
      setMaxCaptureCountForTurn(maxCaptureCountForTurn - 1);
    }
    // Update the board state
    setBoard(newBoard);
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
                    ) : cell === "bq" ? (
                      <img src={bKSvg} alt="Black king" />
                    ) : cell === "w" ? (
                      <img src={wMSvg} alt="White piece" />
                    ) : cell === "wq" ? (
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
        <div className="game-config">
          <div className="board-editor">
            <button onClick={() => navigate("/boardEditor")}>
              Board editor
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Game;
