<script>
  let serverStatus = $state("connecting…");

  $effect(() => {
    const es = new EventSource("/status");

    es.onmessage = (e) => {
      serverStatus = e.data || "unknown";
    };

    es.onerror = () => {
      serverStatus = "connection lost (retrying…)";
    };

    return () => es.close();
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
