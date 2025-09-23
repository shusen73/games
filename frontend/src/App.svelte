<script>
  let serverStatus = $state("unknown");

  async function checkHealth() {
    try {
      const res = await fetch("/healthz", { cache: "no-store" });
      serverStatus = res.ok ? "online" : "offline";
    } catch {
      serverStatus = "offline";
    }
  }

  $effect(() => {
    checkHealth();

    // poll every 10 seconds
    const id = setInterval(checkHealth, 10000);

    // cleanup when component is destroyed
    return () => clearInterval(id);
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
  <div class="message" role="status" aria-live="polite">
    Server status: {serverStatus}
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
