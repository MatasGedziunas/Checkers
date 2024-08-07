let Boards = {
  startingBoard: [
    [".", "b", ".", "b", ".", "b", ".", "b", ".", "b"],
    ["b", ".", "b", ".", "b", ".", "b", ".", "b", "."],
    [".", "b", ".", "b", ".", "b", ".", "b", ".", "b"],
    ["b", ".", "b", ".", "b", ".", "b", ".", "b", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", ".", "."],
    [".", "w", ".", "w", ".", "w", ".", "w", ".", "w"],
    ["w", ".", "w", ".", "w", ".", "w", ".", "w", "."],
    [".", "w", ".", "w", ".", "w", ".", "w", ".", "w"],
    ["w", ".", "w", ".", "w", ".", "w", ".", "w", "."],
  ],
  capturesBoard: [
    [".", "b", ".", "b", ".", "b", ".", "b", ".", "."],
    ["b", ".", "b", ".", "b", ".", "b", ".", "b", "."],
    [".", ".", ".", "b", ".", "b", ".", "w", ".", "b"],
    [".", ".", "b", ".", ".", ".", "b", ".", "b", "."],
    [".", ".", ".", "w", ".", "w", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", ".", "."],
    [".", "w", ".", "w", ".", "w", ".", "w", ".", "."],
    ["w", ".", "w", ".", "w", ".", "w", ".", "w", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", ".", "."],
    ["w", ".", "w", ".", "w", ".", "w", ".", "w", "."],
  ],

  doubleCapturesBoard: [
    [".", "b", ".", "b", ".", "b", ".", "b", ".", "."],
    ["b", ".", "b", ".", "b", ".", "b", ".", "b", "."],
    [".", "b", ".", ".", ".", "b", ".", ".", ".", "b"],
    [".", ".", ".", ".", "b", ".", "b", ".", "b", "."],
    [".", ".", ".", "b", ".", "w", ".", ".", ".", "."],
    [".", ".", ".", ".", "w", ".", ".", ".", ".", "."],
    [".", "w", ".", "w", ".", "w", ".", "w", ".", "."],
    ["w", ".", "w", ".", "w", ".", "w", ".", "w", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", ".", "."],
    ["w", ".", "w", ".", "w", ".", "w", ".", "w", "."],
  ],
  startingBoardString:
    ".b.b.b.b.b b.b.b.b.b. .b.b.b.b.b b.b.b.b.b. .......... .......... .w.w.w.w.w w.w.w.w.w. .w.w.w.w.w w.w.w.w.w.",
  capturesBoardString:
    ".b.b.b.b.b b.b.b.b.b. .b.b.b.b.b b.b.b.b.b. .......... ....b..... .b.w...w.. .w...w.... .b........ w.w.w.w.w.",
  emptyBoardString:
    ".......... .......... .......... .......... .......... .......... .......... .......... .......... ..........",
};

export default Boards;
