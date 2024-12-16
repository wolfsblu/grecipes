<script lang="ts">
    import {type Component, type Snippet} from "svelte";

    interface Props {
        Header?: Component,
        children: Snippet
    }

    let { Header, children }: Props = $props()
</script>

<Header />
<main class="flex-grow">
    {@render children()}
</main>