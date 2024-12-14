<script lang="ts">
    import {type Component, type Snippet} from "svelte";

    interface Props {
        Header?: Component,
        children: Snippet
    }

    let { Header, children }: Props = $props()
</script>

<Header />
<main class="flex flex-col flex-grow flex-wrap items-center justify-center p-3">
    {@render children()}
</main>