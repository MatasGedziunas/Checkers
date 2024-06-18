import Boards from "./boards";

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

  getPossibleMoves(board, row, col) {
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
      return this.getMaxCaptureCountMoves(captures);
      // returns same structure object as else block (with captureCount)
    } else {
      return {
        captureCount: 0,
        cords: this.getNonCaptureMoves(board, row, col, directionRow),
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
    this.log(
      `In row ${row} ; col ${col}. Trying to capture row ${rowToBeCaptured} ; col ${colToBeCaptured} ; piece ${
        this.isInBounds(board, rowToBeCaptured, colToBeCaptured)
          ? board[rowToBeCaptured][colToBeCaptured]
          : "out of bounds"
      }`
    );
    if (this.canBeCaptured(board, row, col, rowToBeCaptured, colToBeCaptured)) {
      this.log(
        `Capturing row ${rowToBeCaptured} ; col ${colToBeCaptured} ; piece ${
          this.isInBounds(board, rowToBeCaptured, colToBeCaptured)
            ? board[rowToBeCaptured][colToBeCaptured]
            : "out of bounds"
        }`
      );
      let newBoard = this.cloneBoard(board);
      let rowAfterCapture = rowToBeCaptured + rowToBeCaptured - row;
      let colAfterCapture = colToBeCaptured + colToBeCaptured - col;
      newBoard[rowToBeCaptured][colToBeCaptured] = ".";
      newBoard[row][col] = ".";
      newBoard[rowAfterCapture][colAfterCapture] = board[row][col];
      let directions = [-1, 1];
      const startingCapturesCount = capturesCount;
      for (let rowDirection of directions) {
        for (let colDirection of directions) {
          this.log(` captures count in loop: ${capturesCount}`);
          let tempCapturesCount = this.getCaptures(
            newBoard,
            rowAfterCapture,
            colAfterCapture,
            rowAfterCapture + rowDirection,
            colAfterCapture + colDirection,
            startingCapturesCount + 1
          );
          if (tempCapturesCount > capturesCount) {
            capturesCount = tempCapturesCount;
          }
        }
      }
    }
    this.log(` captures count ${capturesCount}`);
    return capturesCount;
  }

  getMaxCaptureCountMoves(captures) {
    let maxCaptureCount = 0;
    captures.forEach((capture) => {
      maxCaptureCount = Math.max(maxCaptureCount, capture.capturesCount);
    });
    this.log(captures);
    this.log(`Max capture count ${maxCaptureCount}`);
    captures = captures.filter(
      (capture) => capture.capturesCount == maxCaptureCount
    );
    let moves = [];
    for (let capture of captures) {
      moves.push(capture.move);
    }
    return { cords: moves, captureCount: maxCaptureCount };
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
