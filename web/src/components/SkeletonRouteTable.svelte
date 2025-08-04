<ul class="list bg-base-100 rounded-box shadow-md">
    <li class="p-4 pb-2 text-xs opacity-60 tracking-wide">
        <div class="skeleton h-4 w-24"></div>
    </li>
    {#each Array(5) as _, index}
    <li class="list-row">
        <div>
            <div class="skeleton w-10 h-10 rounded"></div>
        </div>
        <div class="list-col-grow">
            <div class="skeleton h-4 w-32 mb-1"></div>
            <div class="skeleton h-3 w-20"></div>
        </div>
        <div class="text-right">
            <div class="skeleton h-4 w-16 mb-1 ml-auto"></div>
            <div class="skeleton h-3 w-12 ml-auto"></div>
        </div>
    </li>
    {/each}
</ul>