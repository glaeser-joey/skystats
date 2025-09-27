<script>
    import { onMount, onDestroy } from 'svelte';

    let data = [];
    let endpoint = 'api/stats/types/top'
    let loading = true;
    let error = null;
    let interval = null;

    async function fetchData() {

        try {
            const response = await fetch(endpoint);
            if (!response.ok) {
                throw new Error(`${response.status}`);
            }
            const result = await response.json();
            data = result.filter(aircraft => aircraft.aircraft_type);
            error = null;
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
        }
    }

    onMount(() => {
        fetchData();
        interval = setInterval(fetchData, 60000);
    })

    onDestroy(() => {
        if (interval) {
            clearInterval(interval);
        }
    });
</script>

{#if loading}
    <div class="flex justify-center py-8">
        <span class="loading loading-ring loading-lg"></span>
    </div>
{:else if error}
    <div class="flex alert alert-error">
        <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
        <span>Something went wrong: {error}</span>
    </div>
{:else if data.length === 0}
    <div class="alert alert-info">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-current shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
        <span>No Top Types data available</span>
    </div>
{:else}
    <ul class="list bg-base-100 rounded-box shadow-md">
    <li class="p-4 pb-2 text-xs opacity-60 tracking-wide">Top Types</li>

    {#each data as type, index}
    <li class="list-row">
        <div class="flex items-center justify-center">
            <div class="opacity-80"><a href="https://skybrary.aero/aircraft/{type.aircraft_type}" target="_blank" rel="noopener noreferrer" class="badge badge-accent badge-soft badge-outline badge-md sm:badge-lg md:badge-xl lg:badge-2xl rounded-md hover:badge-accent">
            <span class="!text-md sm:!text-lg md:!text-xl lg:!text-2xl font-thin text-accent opacity-80 font-mono">{type.aircraft_type}</span>
        </a></div>
        </div>
        <div class="flex flex-col justify-center list-col-grow">
            <progress class="progress progress-accent w-auto" value="{type.percentage}" max="100"></progress>
        </div>
        <div class="text-right">
            <div class="font-semibold">{type.count.toLocaleString()}</div>
            <div class="text-xs opacity-60">aircraft</div>
        </div>
        </li>
    {/each}
    </ul>
{/if}
