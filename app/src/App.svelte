<script lang="ts">
    import {createRouter} from "./lib/router.svelte";

    const router = createRouter()
    router.registerRoutes()
    let Page = $derived(router.page)
</script>

<Page/>