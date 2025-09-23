<script>
    import { newLocalGame, userState } from "../lib/state.svelte";
</script>

<div class="gameArea">
    <button
        class=""
        onclick={() => {
            userState.view = "local";
            newLocalGame();
        }}
    >
        Play Local
    </button>
    <button
        class=""
        onclick={() => {
            userState.view = "online";
        }}
    >
        Play Online
    </button>
</div>

<style></style>
