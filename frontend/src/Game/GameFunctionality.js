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
    board = [
      [".", "b", ".", "b", ".", "b", ".", "b", ".", "."],
      ["b", ".", "b", ".", "b", ".", "b", ".", "b", "."],
      [".", ".", ".", "b", ".", "b", ".", ".", ".", "b"],
      [".", ".", "b", ".", ".", ".", "b", ".", ".", "."],
      [".", ".", ".", "w", ".", "w", ".", ".", ".", "."],
      [".", ".", ".", ".", ".", ".", ".", ".", ".", "."],
      ["w", ".", "w", ".", "w", ".", "w", ".", "w", "."],
      [".", "w", ".", "w", ".", "w", ".", "w", ".", "w"],
      [".", ".", ".", ".", ".", ".", ".", ".", ".", "."],
      ["w", ".", "w", ".", "w", ".", "w", ".", "w", "."],
    ];
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

  getPossibleMoves(board, row, col) {
    console.log(`possible moves for row ${row} col ${col}`);
    if (board[row][col] == ".") {
      return {
        isCaptures: false,
        moves: [],
      };
    }
    let directionRow;
    if (board[row][col][0] == "b") {
      directionRow = 1;
    } else {
      directionRow = -1;
    }

    let captures = [];
    let directions = [-1, 1];

    for (let rowDirection of directions) {
      for (let colDirection of directions) {
        let capturesCount = this.getCaptures(
          board,
          row,
          col,
          row + rowDirection,
          col + colDirection,
          0
        );
        if (capturesCount > 0) {
          captures.push({
            capturesCount: capturesCount,
            move: [row + 2 * rowDirection, col + 2 * colDirection],
          });
        }
      }
    }
    if (captures.length > 0) {
      let captureMoves = this.getMaxCaptureCountMoves(captures);
      return {
        isCaptures: true,
        moves: captureMoves,
      };
    } else {
      return {
        isCaptures: false,
        moves: this.getNonCaptureMoves(board, row, col, directionRow),
      };
    }
  }

  getCaptures(
    board,
    row,
    col,
    rowToBeCaptured,
    colToBeCaptured,
    capturesCount
  ) {
    if (this.canBeCaptured(board, row, col, rowToBeCaptured, colToBeCaptured)) {
      let newBoard = this.cloneBoard(board);
      newBoard[rowToBeCaptured][colToBeCaptured] = ".";
      let rowAfterCapture = rowToBeCaptured + rowToBeCaptured - row;
      let colAfterCapture = colToBeCaptured + colToBeCaptured - col;
      let directions = [-1, 1];
      for (let rowDirection of directions) {
        for (let colDirection of directions) {
          capturesCount = this.getCaptures(
            newBoard,
            rowAfterCapture,
            colAfterCapture,
            rowAfterCapture + rowDirection,
            colAfterCapture + colDirection,
            capturesCount + 1
          );
        }
      }
    }
    return capturesCount;
  }

  getMaxCaptureCountMoves(captures) {
    let maxCaptureCount = 0;
    captures.forEach((capture) => {
      maxCaptureCount = Math.max(maxCaptureCount, capture.capturesCount);
    });
    captures = captures.filter(
      (capture) => capture.capturesCount == maxCaptureCount
    );
    console.log(captures);
    let moves = [];
    for (let capture of captures) {
      moves.push(capture.move);
    }
    return moves;
  }

  getNonCaptureMoves(board, row, col, directionRow) {
    let possibleMoves = [];
    let directions = [-1, 1];
    for (let directionCol of directions) {
      let rowAfterMove = row + directionRow;
      let colAfterMove = col + directionCol;
      if (
        this.isInBounds(board, rowAfterMove, colAfterMove) &&
        board[rowAfterMove][colAfterMove] == "."
      ) {
        possibleMoves.push([rowAfterMove, colAfterMove]);
      }
    }
    return possibleMoves;
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
