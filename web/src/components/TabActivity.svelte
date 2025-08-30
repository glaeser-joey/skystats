<script>
  import ActivityFlightsByPeriod from "./ActivityFlightsByPeriod.svelte";
  import MetricAircraftSeen from "./MetricAircraftSeen.svelte";
  import MetricFlightsSeen from "./MetricFlightsSeen.svelte";
</script>

<div class="grid grid-cols-1 lg:grid-cols-1 xl:grid-cols-2 mt-10 gap-6">
  <div class="col">
    <MetricFlightsSeen />
  </div>
  <div class="col">
    <MetricAircraftSeen />
  </div>
  <div class="col">
    <ActivityFlightsByPeriod />
  </div>
</div>
