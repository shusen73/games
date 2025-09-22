<script>
  import { onMount } from "svelte";

  let backendOK = null;

  onMount(async () => {
    try {
      const ac = new AbortController();
      const timeout = setTimeout(() => ac.abort(), 1500);
      const res = await fetch("/healthz", { signal: ac.signal });
      clearTimeout(timeout);
      backendOK = res.ok;
    } catch {
      backendOK = false;
    }
  });
</script>

<main class="">
  <div class="gameName">
    <h1>Connect 5</h1>
    <div class="gameDescription">
      <p class="">Five-in-a-row, simple & fast.</p>
    </div>
  </div>
  <div class="gameArea">
    <button class=""> Play Local </button>
    <button class=""> Play Online </button>
  </div>
  <div class="message">
    <div role="status" aria-live="polite" class="meta">
      <span></span>
      <span>
        {backendOK === null
          ? "Backend: unknown"
          : backendOK
            ? "Backend: reachable"
            : "Backend: unreachable"}
      </span>
    </div>
  </div>

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
