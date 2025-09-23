export const PLAYERS = {
    "BLACK": "BLACK",
    "WHITE": "WHITE",
}

export class Game {
    constructor(boarderSize = 15) {
        this.winner = null;
        this.boardSize = boarderSize;
        this.board = Array.from(Array(boarderSize), () => new Array(boarderSize).fill(null));
        this.currentPlayer = PLAYERS.BLACK;      
    }
}
