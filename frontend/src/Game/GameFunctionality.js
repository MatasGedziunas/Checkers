class GameFunctionality {
  constructor() {}

  getStartingBoard() {
    const size = 10;
    let board = Array.from({ length: size }, () => Array(size).fill("."));

    // Place black pieces
    for (let i = 0; i < 4; i++) {
      for (let j = 0; j < size; j++) {
        if ((i + j) % 2 === 1) {
          board[i][j] = "b";
        }
      }
    }

    // Place white pieces
    for (let i = size - 1; i >= size - 4; i--) {
      for (let j = 0; j < size; j++) {
        if ((i + j) % 2 === 1) {
          board[i][j] = "w";
        }
      }
    }

    return board;
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
    while (i < boardString.length) {
      if (boardString[i] == " ") {
        board.push([]);
        j += 1;
      } else {
        board[j].push(boardString[i]);
      }
      i += 1;
    }
  }

  showPossibleMoves(board, row, col) {
    let directionY;
    if (board[row][col] == ".") {
      return {
        isCaptures: false,
        moves: [],
      };
    }
    if (board[row][col][0] == "b") {
      directionY = 1;
    } else {
      directionY = -1;
    }
    let possibleMoves = [];
    let captures = [];
    if (
      this.isInBounds(board, row, col, 1, directionY) &&
      board[row + directionY][col + 1] == "." &&
      captures.length == 0
    ) {
      possibleMoves.push([row + directionY, col + 1]);
    } else if (this.canBeCaptured(board, row, col, row + directionY, col + 1)) {
      captures.push([row + directionY, col + 1]);
    }
    if (
      this.isInBounds(board, row, col, -1, directionY) &&
      board[row + directionY][col - 1] == "." &&
      captures.length == 0
    ) {
      possibleMoves.push([row + directionY, col - 1]);
    } else if (this.canBeCaptured(board, row, col, row + directionY, col - 1)) {
      captures.push([row + directionY, col - 1]);
    }
    if (
      this.isInBounds(board, row, col, +1, -directionY) &&
      this.canBeCaptured(board, row - directionY, col + 1)
    ) {
      captures.push([row - directionY, col + 1]);
    }
    if (
      this.isInBounds(board, row, col, -1, -directionY) &&
      this.canBeCaptured(board, row - directionY, col - 1)
    ) {
      captures.push([row - directionY, col - 1]);
    }
    if (captures.length > 0) {
      console.log(captures.length);
      return {
        isCaptures: true,
        moves: captures,
      };
    } else {
      console.log(possibleMoves.length);
      return {
        isCaptures: false,
        moves: possibleMoves,
      };
    }
  }

  isInBounds(board, row, col, directionX, directionY) {
    return (
      row + directionY >= 0 &&
      row + directionY < board.length &&
      col + directionX >= 0 &&
      col + directionX < board[0].length
    );
  }

  canBeCaptured(board, rowFrom, colFrom, rowTo, colTo) {
    const directionY = rowTo - rowFrom;
    const directionX = colTo - colFrom;
    return (
      this.isInBounds(board, rowTo, colTo, directionX, directionY) &&
      board[rowTo][colTo] != "." &&
      board[rowFrom][colFrom] != "." &&
      board[rowFrom][colFrom][0] !== board[rowTo][colTo][0] && // pieces not same color
      board[rowTo + directionY][colTo + directionX] === "."
    );
  }
}

export default GameFunctionality;
