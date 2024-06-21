import Boards from "./Boards";

class GameFunctionality {
  constructor(logging) {
    if (logging == undefined) {
      logging = false;
    }
    this.logging = logging;
  }

  log(string) {
    if (this.logging) {
      console.log(string);
    }
  }

  getStartingBoard() {
    return Boards.startingBoard;
  }

  getCapturesBoard() {
    return Boards.capturesBoard;
  }

  getDoubleCapturesBoard() {
    return Boards.doubleCapturesBoard;
  }

  encodeBoard(board) {
    let encoded = "";
    for (let i = 0; i < board.length; i++) {
      for (let j = 0; j < board[i].length; j++) {
        encoded += board[i][j];
      }
      encoded += " ";
    }
    return encoded;
  }

  decodeBoard(boardString) {
    let board = [[]];
    let i = 0;
    let j = 0;
    while (i < boardString.length - 1) {
      if (boardString[i] == " ") {
        board.push([]);
        j += 1;
      } else {
        board[j].push(boardString[i]);
      }
      i += 1;
    }
    return board;
  }

  updateQueens(board) {
    for (let i = 0; i < board.length; i++) {
      for (let j = 0; j < board[i].length; j++) {
        if (this.isQueening(board, i, j) && board[i][j] != ".") {
          board = this.makeQueen(board, i, j);
        }
      }
    }
    console.log(board);
    return board;
  }

  isQueening(board, row, col) {
    if (this.isQueen(board, row, col) || board[row][col] == ".") {
      return false;
    }
    const color = board[row][col];
    if (color == "w") {
      return row == 0;
    } else {
      return row == 9;
    }
  }

  isQueen(board, row, col) {
    return board[row][col][1] == "q";
  }

  makeQueen(board, row, col) {
    console.log(row, col);
    board[row][col] == "w"
      ? (board[row][col] = "wq")
      : (board[row][col] = "bq");
    return board;
  }

  isInBounds(board, row, col) {
    return row >= 0 && row < board.length && col >= 0 && col < board[0].length;
  }

  canBeCaptured(board, rowFrom, colFrom, rowTo, colTo) {
    const directionRow = rowTo - rowFrom;
    const directionCol = colTo - colFrom;
    return (
      this.isInBounds(board, rowTo + directionRow, colTo + directionCol) &&
      board[rowTo][colTo] != "." &&
      board[rowFrom][colFrom] != "." &&
      board[rowFrom][colFrom][0] !== board[rowTo][colTo][0] && // pieces not same color
      board[rowTo + directionRow][colTo + directionCol] === "."
    );
  }

  cloneBoard(arr) {
    return arr.map((innerArr) => innerArr.slice());
  }
}

export default GameFunctionality;
