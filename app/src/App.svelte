<script lang="ts">
    import {createRouter} from "./lib/router.svelte";
    import Navbar from "./lib/components/Navbar.svelte";
    import Layout from "./Layout.svelte";

    const router = createRouter()
    router.registerRoutes()
    let Page = $derived(router.page)
</script>

<Page/>