<script>
    import FlightsByPeriod from './charts/FlightsByPeriod.svelte';
    let selectedTab = 0;
</script>

<div class="grid grid-cols-1 lg:grid-cols-1 xl:grid-cols-2 mt-10 gap-6">

<div class="bg-base-100 rounded-box shadow-md pt-6 p-8">
  <div>
      <!-- <div class="text-4xl font-thin text-accent opacity-80 font-mono">{airport.airport_code}</div> -->

      <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-extralight tracking-wider ">Flights Seen</h2>
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
          <FlightsByPeriod period="day"/>
      {:else if selectedTab === 1}
          <FlightsByPeriod period="month"/>
      {:else if selectedTab === 2}
          <FlightsByPeriod period="year"/>
      {/if}
  </div>


</div>
</div>
