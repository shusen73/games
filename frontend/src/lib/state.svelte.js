import { Game, PLAYERS } from "./lib";

export const userState = $state({
  view: "menu", // local, online
  game: {
    boardSize: 15,
    board: Array(15)
      .fill(null)
      .map(() => Array(15).fill(null)),
    currentPlayer: PLAYERS.BLACK,
    winner: null,
  },
});

export const newLocalGame = () => {
  const n = 15;
  userState.view = "local";
  userState.game.boardSize = n;
  userState.game.board = Array.from({ length: n }, () => Array(n).fill(null));
  userState.game.currentPlayer = PLAYERS.BLACK;
  userState.game.winner = null;
};

export const placeStone = (i, j) => {
  if (userState.game.winner != null) {
    return;
  }
  // if the current cell is alreay opcupies, do nothing
  if (userState.game.board[i][j] != null) {
    return;
  }
  userState.game.board[i][j] = userState.game.currentPlayer;
  if (hasWinFrom(userState.game.board, i, j)) {
    console.log("Winner:", userState.game.currentPlayer);
    userState.game.winner = userState.game.currentPlayer;
  }
  if (userState.game.currentPlayer == PLAYERS.BLACK) {
    userState.game.currentPlayer = PLAYERS.WHITE;
  } else {
    userState.game.currentPlayer = PLAYERS.BLACK;
  }
};

/**
 * True if the just-placed stone at (row,col) creates a contiguous run
 * of at least `target` horizontally/vertically/diagonally.
 */
const hasWinFrom = (board, row, col) => {
  const directions = [
    [0, 1], // horizontal →
    [1, 0], // vertical ↓
    [1, 1], // diagonal ↘
    [1, -1], // diagonal ↙
  ];
  for (const [dr, dc] of directions) {
    const length = countConnected(board, row, col, dr, dc);
    if (length >= 5) return true;
  }
  return false;
};

export function restartGame() {
  newLocalGame();
}

export function backToMenu() {
  userState.view = "menu";
}

/**
 * Count the contiguous stones of the same player that include (row,col)
 * along a single direction (dr, dc), scanning both ways until a gap.
 * Returns the run length and the start/end coordinates of the run.
 */
const countConnected = (board, row, col, dr, dc) => {
  const player = board[row][col];

  const rows = board.length;
  const cols = rows;
  const inBounds = (r, c) => r >= 0 && r < rows && c >= 0 && c < cols;

  let length = 1; // count the placed stone

  // expand forward
  let r = row + dr,
    c = col + dc;
  while (inBounds(r, c) && board[r][c] === player) {
    length++;
    r += dr;
    c += dc;
  }

  // expand backward
  r = row - dr;
  c = col - dc;
  while (inBounds(r, c) && board[r][c] === player) {
    length++;
    r -= dr;
    c -= dc;
  }

  return length;
};
