<script>
  import { userState, placeStone } from "../lib/state.svelte";
  import VictoryModal from "./VictoryModal.svelte";
</script>

<div class="board" style={`--n:${userState.game.boardSize};`}>
  {#each userState.game.board as row, i}
    {#each row as cell, j}
      <button class="cell" onclick={() => placeStone(i, j)}>
        {#if userState.game.board[i][j] == "BLACK"}
          <span class="stone black"></span>
        {:else if userState.game.board[i][j] == "WHITE"}
          <span class="stone white"></span>
        {/if}
      </button>
    {/each}
  {/each}
</div>
<div class="turn">
  Player {userState.game.currentPlayer} to move
</div>
<VictoryModal />

<style>
  .board {
    display: grid;
    grid-template-columns: repeat(var(--n), 1fr);
    grid-template-rows: repeat(var(--n), 1fr);
    gap: 2px;
    aspect-ratio: 1 / 1;
    background: #eaddcf;
    padding: 6px;
    border-radius: 12px;
    box-shadow: 0 6px 24px rgba(0, 0, 0, 0.25);
  }

  .cell {
    position: relative;
    border: none;
    background: #8c7851;
    border-radius: 4px;
    cursor: pointer;
    outline: none;
    aspect-ratio: 1 / 1;
    transition:
      transform 0.05s ease,
      box-shadow 0.12s ease,
      background 0.12s ease;
  }
  .cell:hover {
    background: #716040;
  }
  .cell:active {
    transform: translateY(1px);
  }
  .cell:focus-visible {
    box-shadow: 0 0 0 3px rgba(110, 168, 254, 0.7);
  }

  .stone {
    position: absolute;
    inset: 12%; /* margin from cell edge */
    border-radius: 50%;
    display: block;
    box-shadow:
      inset 0 2px 6px rgba(0, 0, 0, 0.35),
      0 2px 8px rgba(0, 0, 0, 0.25);
  }
  .stone.black {
    background: #101418;
  }
  .stone.white {
    background: #eaeef5;
    box-shadow:
      inset 0 2px 6px rgba(0, 0, 0, 0.25),
      0 0 0 1px rgba(0, 0, 0, 0.15),
      0 2px 8px rgba(0, 0, 0, 0.25);
  }

  /* Turn text */
  .turn {
    margin-top: 2rem;
    font-weight: 600;
    letter-spacing: 0.25px;
    text-align: center;
  }
</style>
