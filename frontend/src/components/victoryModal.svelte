<script>
  import { userState, restartGame, backToMenu } from "../lib/state.svelte";
  let dialog;
</script>

<dialog
  open={!!userState.game.winner}
  bind:this={dialog}
  oncancel={(e) => e.preventDefault()}
>
  <h2>Player {userState.game.winner} wins!</h2>
  <button onclick={restartGame}>Restart</button>
  <button onclick={backToMenu}>Menu</button>
</dialog>
