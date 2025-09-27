<script>
    import TypesByPeriod from './charts/TypesByPeriod.svelte';
    let selectedTab = 0;
</script>

<div class="bg-base-100 rounded-box shadow-md pt-6 p-8">
    <div>
      <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-extralight tracking-wider ">Aircraft Types Seen</h2>
      <div class="tabs tabs-box">
        <input
          type="radio"
          name="activity_tabs"
          class="tab"
          aria-label="24hrs"
          bind:group={selectedTab}
          value={0}
          checked
        />
        <input
          type="radio"
          name="activity_tabs"
          class="tab"
          aria-label="Month"
          bind:group={selectedTab}
          value={1}
        />
        <input
          type="radio"
          name="activity_tabs"
          class="tab"
          aria-label="Year"
          bind:group={selectedTab}
          value={2}
        />
    </div>
    </div>
      {#if selectedTab === 0}
          <TypesByPeriod period="day"/>
      {:else if selectedTab === 1}
          <TypesByPeriod period="month"/>
      {:else if selectedTab === 2}
          <TypesByPeriod period="year"/>
      {/if}
  </div>
</div>


