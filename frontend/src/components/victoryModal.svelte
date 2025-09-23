<script>
  import { userState } from "../lib/state.svelte";


    let dialog
</script>

<button on:click={() => dialog.showModal()}>Open It!</button>

{#if userState.game.winner !== null}
<dialog bind:this={dialog} on:close={() => console.log('closed')}>
    This is amazingly simple! (press esc to close)
</dialog>{/if}
