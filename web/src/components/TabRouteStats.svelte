<script>
    import RouteTopAirlines from './RouteTopAirlines.svelte';
    import RouteTopAirportsDomestic from './RouteTopAirportsDomestic.svelte';
    import RouteTopAirportsInternational from './RouteTopAirportsInternational.svelte';
    import RouteTopRoutes from './RouteTopRoutes.svelte';
    import RouteTopCountriesOrigin from './RouteTopCountriesOrigin.svelte';
    import RouteTopCountriesDestination from './RouteTopCountriesDestination.svelte';
    // import TypesTop from './TypesTop.svelte';

</script>
<div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 mt-10 gap-6">
    <RouteTopAirlines />
    <RouteTopAirportsDomestic  />
    <RouteTopAirportsInternational />
    <RouteTopRoutes />
    <RouteTopCountriesOrigin />
    <RouteTopCountriesDestination />
    <!-- <TypesTop /> -->
  </div>