<script lang="ts" generics="T">
    import type {Snippet} from "svelte";

    interface Props {
        class: string | null | undefined
        type: "button" | "submit" | "reset" | null | undefined
        children: Snippet
    }

    let {children, class: cls, ...rest}: Props = $props()
    let fields = {...rest}
</script>

<label>
    <button class="
    active:bg-orange-50
    active:shadow-inner
    active:text-orange-700
    bg-orange-100
    font-semibold
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
        {@render children()}
    </button>
</label>