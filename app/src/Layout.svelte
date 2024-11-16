<script lang="ts">
    import type {Component, Snippet} from "svelte";

    interface Props {
        Header?: Component,
        children: Snippet
    }

    let { Header, children }: Props = $props()
</script>

<Header />
<main class="p-3">
    {@render children()}
</main>