<script>
    import MotionFastestAircraft from './MotionFastestAircraft.svelte';
    import MotionSlowestAircraft from './MotionSlowestAircraft.svelte';
    import MotionHighestAircraft from './MotionHighestAircraft.svelte';
    import MotionLowestAircraft from './MotionLowestAircraft.svelte';
</script>

<div class="grid grid-cols-1 lg:grid-cols-2 mt-10 gap-6">
    <MotionFastestAircraft />
    <MotionSlowestAircraft />
    <MotionHighestAircraft /> 
    <MotionLowestAircraft />
</div>