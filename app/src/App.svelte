<script lang="ts">
    import {createRouter} from "./lib/router.svelte";
    import {createUser} from "./lib/auth/user.svelte";

    const router = createRouter()
    let Page = $derived(router.page)

    createUser()
        .fetchProfile()
        .catch(() => {}) // User has no session cookie, that's fine
        .finally(() => router.registerRoutes()) // Router needs to know whether user is logged in
</script>

<Page/>