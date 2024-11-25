<script lang="ts">
    import {createRouter} from "./lib/router.svelte";
    import {createUser} from "./lib/auth/user.svelte";

    const router = createRouter()
    let Page = $derived(router.page)

    createUser()
        .fetchProfile()
        .catch(() => {}) // User has no session cookie, that's fine
        .finally(() => router.registerRoutes()) // Routes need to check for logged-in user so we load them after we fetch the profile
</script>

<Page/>