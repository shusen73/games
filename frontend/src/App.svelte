<script>
  import Board from "./components/Board.svelte";
  import Menu from "./components/Menu.svelte";
  import VictoryModal from "./components/VictoryModal.svelte";

  import { userState } from "./lib/state.svelte";
</script>

<main class="">
  <div class="gameName">
    <h1>Connect 5</h1>
    <div class="gameDescription"></div>
  </div>

  {#if userState.view === "menu"}
    <Menu />
  {:else if userState.view === "local"}
    <Board />
  {:else if userState.view === "online"}
    <Board />
  {/if}

  <footer></footer>
</main>

<style>
  main {
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
    text-align: center;
  }
</style>
