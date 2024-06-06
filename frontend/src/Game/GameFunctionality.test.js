/* eslint-disable no-undef */
import GameFunctionality from "./GameFunctionality";

import Boards from "./boards";

let functionalityWithoutLogging = new GameFunctionality(false);
let functionalityWithLogging = new GameFunctionality(true);

let startingBoard = Boards.startingBoard;
let startingBoardString = Boards.startingBoardString;
let capturesBoard = Boards.capturesBoard;

test("getStartingBoard_shouldReturnStartingBoard", () => {
  expect(functionalityWithoutLogging.getStartingBoard()).toEqual(startingBoard);
});

test("getCapturesBoard_shouldReturnCapturesBoard", () => {
  expect(functionalityWithoutLogging.getCapturesBoard()).toEqual(capturesBoard);
});

test("getCapturesBoard_shouldReturnDoubleCapturesBoard", () => {
  expect(functionalityWithoutLogging.getDoubleCapturesBoard()).toEqual(
    Boards.doubleCapturesBoard
  );
});

test("encodeBoard_startingBoard_shouldReturnStartingBoardString", () => {
  expect(functionalityWithoutLogging.encodeBoard(startingBoard)).toEqual(
    startingBoardString
  );
});

test("decodeBoard_startingBoardString_shouldReturnStartingBoard", () => {
  // console.log(functionality.decodeBoard(startingBoardString));
  expect(functionalityWithoutLogging.decodeBoard(startingBoardString)).toEqual(
    startingBoard
  );
});

test("canBeCaptured_boardRowOutOfBounds_shouldReturnFalse", () => {
  expect(functionalityWithoutLogging.isInBounds(startingBoard, 10, 1)).toEqual(
    false
  );
});

test("canBeCaptured_boardColOutOfBounds_shouldReturnFalse", () => {
  expect(functionalityWithoutLogging.isInBounds(startingBoard, 1, 10)).toEqual(
    false
  );
});

test("canBeCaptured_correctCol_shouldReturnTrue", () => {
  expect(functionalityWithoutLogging.isInBounds(startingBoard, 0, 0)).toEqual(
    true
  );
});

test("canBeCaptured_capturesBoardFrontCapture_shouldReturnTrue", () => {
  expect(
    functionalityWithoutLogging.canBeCaptured(capturesBoard, 2, 7, 1, 8)
  ).toEqual(true);
});

test("canBeCaptured_capturesBoardFrontCapture_shouldReturnFalse", () => {
  expect(
    functionalityWithoutLogging.canBeCaptured(capturesBoard, 5, 7, 6, 6)
  ).toEqual(false);
});

test("canBeCaptured_capturesBoardBackCapture_shouldReturnTrue", () => {
  expect(
    functionalityWithoutLogging.canBeCaptured(capturesBoard, 2, 7, 3, 8)
  ).toEqual(true);
});

test("canBeCaptured_capturesBoardBackCapture_shouldReturnFalse", () => {
  expect(
    functionalityWithoutLogging.canBeCaptured(capturesBoard, 3, 8, 2, 7)
  ).toEqual(false);
});

test("getCaptures_capturesBoard_shouldReturnOneCaptureMove", () => {
  // console.log(functionality.getCaptures(capturesBoard, 3, 2, 4, 3, 0));
  expect(
    functionalityWithoutLogging.getCaptures(capturesBoard, 3, 2, 4, 3, 0)
  ).toEqual(1);
});

test("getCaptures_doubleCapturesBoard_shouldReturnTwoCaptureMoves", () => {
  expect(
    functionalityWithoutLogging.getCaptures(
      Boards.doubleCapturesBoard,
      4,
      5,
      3,
      6,
      0
    )
  ).toEqual(2);
});

test("getPossibleMoves_doubleCapturesBoardMultipleCaptures_shouldReturnOnlyMoveWithMaxCaptureCount", () => {
  expect(
    functionalityWithLogging.getPossibleMoves(Boards.doubleCapturesBoard, 4, 5)
  ).toEqual({ cords: [[2, 7]], captureCount: 2 });
});
