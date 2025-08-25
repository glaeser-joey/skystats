<script>
    import InterestingAircraft from "./InterestingAircraft.svelte";
    import { IconCurrencyDollar } from '@tabler/icons-svelte';
</script>

<InterestingAircraft
    endpoint='api/stats/interesting/civilian'
    title="Civilian Aircraft"
    icon={IconCurrencyDollar}
    aircraftType="civ"
/>