<script>
    import TypesByPeriod from './charts/TypesByPeriod.svelte';
    export let aircraftorflight = 'flights';
    let selectedTab = 0;
    let title = aircraftorflight === 'flights' ? 'Flights Seen' : 'Unique Aircraft Seen';
</script>

<div class="bg-base-100 rounded-box shadow-md pt-6 p-8">
    <div>
      <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-extralight tracking-wider ">{title}</h2>
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
        <input
          type="radio"
          name="activity_tabs"
          class="tab"
          aria-label="All"
          bind:group={selectedTab}
          value={3}
        />
    </div>
    </div>
      {#if selectedTab === 0}
          <TypesByPeriod period="day" aircraftorflight={aircraftorflight}/>
      {:else if selectedTab === 1}
          <TypesByPeriod period="month" aircraftorflight={aircraftorflight}/>
      {:else if selectedTab === 2}
          <TypesByPeriod period="year" aircraftorflight={aircraftorflight}/>
      {:else if selectedTab === 3}
          <TypesByPeriod period="all" aircraftorflight={aircraftorflight}/>
      {/if}
  </div>
</div>


