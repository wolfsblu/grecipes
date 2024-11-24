<script lang="ts">
    import {type Component} from 'svelte';

    interface Props {
        icon: Component
        text?: string
        href: string
    }

    const {icon: Icon, text, href}: Props = $props()
</script>

<a class="
        active:bg-orange-50
        active:shadow-inner
        active:text-orange-700
        bg-orange-50
        flex
        gap-x-2
        group
        hover:bg-orange-100
        hover:text-orange-800
        items-center
        p-2
        rounded-full
        shadow-md
        text-orange-700
    "
   {href}>
    <span class="group-active:text-orange-700 group-hover:text-orange-800 text-orange-700 relative h-7 w-7">
        <Icon/>
    </span>
    {#if text}
        <span class="hidden sm:inline font-semibold">{text}</span>
    {/if}
</a>