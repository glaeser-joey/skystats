<script>
// @ts-nocheck
    import { onMount, onDestroy } from 'svelte';
    import { themeChanged } from '../../lib/themeStore.js';
    import {
        Chart,
        BarController,
        BarElement,
        CategoryScale,
        LinearScale,
        TimeScale,
        Title,
        Tooltip,
        Legend
    } from 'chart.js';
    import 'chartjs-adapter-date-fns';


    /* Data setup */

    const endpoint = '/api/stats/charts/flights/year';

    async function fetchData() {
        try {
            loading = true;
            const response = await fetch(endpoint);
            if (!response.ok) {
                throw new Error(`HTTP ${response.status}`);
            }
            const result = await response.json();
            
            if (result.series && result.series.length > 0) {
                const series = result.series[0];
                chartData = {
                    labels: series.points.map(p => new Date(p.x)),
                    dataPoints: series.points.map(p => p.y),
                    label: series.label || 'Aircraft Over Time',
                    unit: result.x?.unit || 'month'
                };
            }
            
            error = null;
        } catch (err) {
            error = err.message || 'Failed to load data';
        } finally {
            loading = false;
        }
    }



    /* Chart Setup */

    Chart.register(
        BarController,
        BarElement,
        CategoryScale,
        LinearScale,
        TimeScale,
        Title,
        Tooltip,
        Legend
    );

    let chartCanvas = null;
    let chart = null;
    let loading = true;
    let error = null;
    let chartData = null;

    // set fallback colours
    let CHART_COLOURS = {
        primary: '#570df8', 
        primaryAlpha: 'rgba(87, 13, 248, 0.2)'
    };

    const MONTH_NAMES = [
        'January', 'February', 'March', 'April', 'May', 'June',
        'July', 'August', 'September', 'October', 'November', 'December'
    ];
    
    
    // this is silly, but seems to be the only way to dynamically get the value of daisyui theme colours
    function getDaisyUIColor(className) {
        const temp = document.createElement('div');
        temp.className = className;
        document.body.appendChild(temp);
        const color = getComputedStyle(temp).backgroundColor || getComputedStyle(temp).color;
        document.body.removeChild(temp);
        return color;
    }

    // get colours from daisyui theme
    function updateChartColors() {
        CHART_COLOURS = {
            primary: getDaisyUIColor('bg-primary'),
            primaryAlpha: getDaisyUIColor('bg-primary') + '33',
            baseContent: getDaisyUIColor('bg-base-content')
        };
        
        // update chart colours and redraw
        if (chart && chartData) {
            chart.data.datasets[0].borderColor = CHART_COLOURS.primary;
            chart.data.datasets[0].backgroundColor = CHART_COLOURS.primaryAlpha;
            chart.options.scales.x.ticks.color = CHART_COLOURS.baseContent;
            chart.options.scales.y.ticks.color = CHART_COLOURS.baseContent;
            chart.options.scales.x.border.color = CHART_COLOURS.baseContent;
            chart.options.scales.y.border.color = CHART_COLOURS.baseContent;
            chart.update();
        }
    }
      
    function createChart() {
        if (!chartCanvas || !chartData) return;
        
        // dataset / chart
        const dataset = {
            label: chartData.label,
            data: chartData.dataPoints,
            borderColor: CHART_COLOURS.primary,
            backgroundColor: CHART_COLOURS.primaryAlpha,
            borderWidth: 1,
            barPercentage: 0.90,
            categoryPercentage: 0.95
        };
        
        // tooltip
        const tooltipCallbacks = {
            title: (tooltipItems) => {
                const index = tooltipItems[0].dataIndex;
                const date = chartData.labels[index];
                return `${MONTH_NAMES[date.getMonth()]} ${date.getFullYear()}`;
            },
            label: (context) => {
                const value = context.parsed.y;
                return value.toLocaleString() + ' aircraft';
            }
        };
        
        // plugin configuration
        const plugins = {
            title: {
                display: false,
                text: 'Flights'
            },
            legend: {
                display: false
            },
            tooltip: {
                intersect: false,
                mode: 'index',
                callbacks: tooltipCallbacks
            }
        };
        
        // scale configuration
        const scales = {
            x: {
                type: 'category',
                grid: {
                    display: false  // Remove vertical gridlines
                },
                ticks: {
                    color: CHART_COLOURS.baseContent,
                    callback: function(value, index) {
                        const date = chartData.labels[index];
                        if (date) {
                            return `${MONTH_NAMES[date.getMonth()].slice(0, 3)} ${date.getFullYear()}`;
                        }
                        return value;
                    }
                },
                border: {
                    color: CHART_COLOURS.baseContent
                }
            },
            y: {
                beginAtZero: true,
                grid: {
                    display: false  // Remove horizontal gridlines
                },
                ticks: {
                    color: CHART_COLOURS.baseContent,
                    callback: (value) => value.toLocaleString()
                },
                border: {
                    color: CHART_COLOURS.baseContent
                }
            }
        };
        
        // chart options
        const options = {
            responsive: true,
            maintainAspectRatio: false,
            plugins,
            scales
        };
        
        // main chart config
        const config = {
            type: 'bar',
            data: {
                labels: chartData.labels,
                datasets: [dataset]
            },
            options
        };

        if (chart) {
            chart.destroy();
        }
        
        chart = new Chart(chartCanvas, config);
    }
    
    
    $: if (!loading && chartData && chartCanvas) {
        setTimeout(() => createChart(), 100);
    }
    
    let unsubscribe;

    onMount(async () => {
        // get colours for first time
        updateChartColors();
        
        // sub to theme change events
        unsubscribe = themeChanged.subscribe(() => {
            setTimeout(() => {
                updateChartColors();
            }, 50);
        });
        
        // get data
        await fetchData();

        // draw chart
        if (chartData && chartCanvas) {
            setTimeout(() => createChart(), 100);
        }
    });

    onDestroy(() => {
        if (chart) {
            chart.destroy();
            chart = null;
        }
        if (unsubscribe) {
            unsubscribe();
        }
    });
</script>

<div class="bg-base-100 rounded-box shadow-md p-8">
    {#if loading}
        <div class="loading">Loading chart data...</div>
    {:else if error}
        <div class="flex alert alert-error">
            <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
            <span>Something went wrong: {error}</span>
        </div>
    {:else if !chartData}
        <div class="alert alert-info">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-current shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
            <span>No data available</span>
        </div>
    {:else}
        <div class="h-100">
            <canvas bind:this={chartCanvas}></canvas>
        </div>
    {/if}
</div>