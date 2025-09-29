<script>
    import FlightsByPeriod from './charts/FlightsByPeriod.svelte';
    let selectedTab = 0;
</script>

<div class="bg-base-100 rounded-box shadow-md pt-6 p-8">
  <div>
      <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-extralight tracking-wider ">Flights Seen</h2>
      <div class="tabs tabs-box">
        <input
          type="radio"
          name="activity_tabs_flights"
          class="tab"
          aria-label="24hrs"
          bind:group={selectedTab}
          value={0}
          checked
        />
        <input
          type="radio"
          name="activity_tabs_flights"
          class="tab"
          aria-label="Month"
          bind:group={selectedTab}
          value={1}
        />
        <input
          type="radio"
          name="activity_tabs_flights"
          class="tab"
          aria-label="Year"
          bind:group={selectedTab}
          value={2}
        />
    </div>
    </div>
      {#if selectedTab === 0}
          <FlightsByPeriod type="flights" period="day"/>
      {:else if selectedTab === 1}
          <FlightsByPeriod type="flights" period="month"/>
      {:else if selectedTab === 2}
          <FlightsByPeriod type="flights" period="year"/>
      {/if}
  </div>


</div>
