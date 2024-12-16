<script lang="ts" generics="T">
    import type {Component, Snippet} from "svelte";

    interface Props {
        class: string | null | undefined
        icon: Component | null | undefined
        type: "button" | "submit" | "reset" | null | undefined
        children: Snippet
    }

    let {children, class: cls, icon: Icon, ...rest}: Props = $props()
    let fields = {...rest}
</script>

<button class="
active:bg-orange-100
active:shadow-inner
active:text-orange-700
bg-orange-100
flex
font-semibold
gap-2
hover:bg-orange-200
hover:text-orange-800
items-center
px-4
py-2
rounded
text-orange-700
{cls}
"
{...fields}>
    {#if Icon}
        <span class="h-4 w-4 relative">
            <Icon />
        </span>
    {/if}
    {@render children()}
</button>